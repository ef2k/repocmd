package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/mux"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

type fromOldestTimeSlice []repository

func (r fromOldestTimeSlice) Len() int {
	return len(r)
}
func (r fromOldestTimeSlice) Less(i, j int) bool {
	return r[i].PushedAt.After(r[j].PushedAt.Time)
}
func (r fromOldestTimeSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type totalCount struct {
	TotalCount githubql.Int `json:"totalCount"`
}

type repository struct {
	ID            githubql.ID       `json:"id"`
	Name          githubql.String   `json:"name"`
	NameWithOwner githubql.String   `json:"nameWithOwner"`
	Description   githubql.String   `json:"description"`
	URL           githubql.URI      `json:"url"`
	CreatedAt     githubql.DateTime `json:"createdAt"`
	UpdatedAt     githubql.DateTime `json:"updatedAt"`
	PushedAt      githubql.DateTime `json:"pushedAt"`
	IsArchived    githubql.Boolean  `json:"isArchived"`
	IsFork        githubql.Boolean  `json:"isFork"`
	IsPrivate     githubql.Boolean  `json:"isPrivate"`
	Stargazers    totalCount        `json:"stargazers"`
	Watchers      totalCount        `json:"watchers"`
	Forks         totalCount        `json:"forks"`
	Ref           struct {
		Target struct {
			Commit struct {
				ID              githubql.ID     `json:"id"`
				AbbreviatedOID  githubql.String `json:"abbreviatedOid"`
				MessageHeadline githubql.String `json:"message"`
			} `json:"commit" graphql:"... on Commit"`
		} `json:"target"`
	} `json:"ref" graphql:"ref(qualifiedName:master)"`
}

var q struct {
	Viewer struct {
		Repositories struct {
			Nodes    []repository
			PageInfo struct {
				EndCursor   githubql.String
				HasNextPage githubql.Boolean
			}
		} `graphql:"repositories(last:100, after:$repositoriesCursor, affiliations:[OWNER])"`
	}
}

func getRepos(client *githubql.Client) ([]repository, error) {
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}
	var repos []repository
	for {
		if err := client.Query(context.Background(), &q, variables); err != nil {
			return nil, err
		}
		repos = append(repos, q.Viewer.Repositories.Nodes...)
		if !q.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(q.Viewer.Repositories.PageInfo.EndCursor)
	}
	timeSortedRepos := make(fromOldestTimeSlice, 0, len(repos))
	timeSortedRepos = append(timeSortedRepos, repos...)
	sort.Sort(timeSortedRepos)
	return timeSortedRepos, nil
}

var (
	httpClient *http.Client
	client     *githubql.Client
)

func handleGetRepos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	repos, err := getRepos(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(repos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handlePatchRepo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	dec := json.NewDecoder(r.Body)
	var repo repository
	if err := dec.Decode(&repo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// NOTE GitHub v4 API is missing various mutations. Updating
	// a repository is missing entirely. As a workaround, we'll use
	// the v3 API to PATCH the repo using REST. Luckily, we have
	// an authenticated client to do this.
	URL := fmt.Sprintf("https://api.github.com/repos/%s", repo.NameWithOwner)
	patchedJSON := struct {
		Archived bool `json:"archived"`
	}{
		Archived: bool(repo.IsArchived),
	}
	json, err := json.Marshal(patchedJSON)

	log.Printf("Raw JSON: %s \n", string(json))

	if err != nil {
		log.Print(err)
	}
	req, err := http.NewRequest("PATCH", URL, bytes.NewBuffer(json))
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	log.Printf("response: %v", resp)
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient = oauth2.NewClient(context.Background(), src)
	client = githubql.NewClient(httpClient)

	r := mux.NewRouter()

	r.HandleFunc("/repos", handleGetRepos).Methods("GET")
	r.HandleFunc("/repos", handlePatchRepo).Methods("PATCH")
	http.Handle("/", r)

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
