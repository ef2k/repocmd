package main

import (
	"context"
	"log"
	"os"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

type repository struct {
	Name       githubql.String
	IsArchived githubql.Boolean
}

var q struct {
	Viewer struct {
		Repositories struct {
			Nodes    []repository
			PageInfo struct {
				EndCursor   githubql.String
				HasNextPage githubql.Boolean
			}
		} `graphql:"repositories(first:100, after:$repositoriesCursor, affiliations:[OWNER])"`
	}
}

func getRepositories(client *githubql.Client) []repository {
	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}
	var repos []repository
	for {
		if err := client.Query(context.Background(), &q, variables); err != nil {
			log.Fatal(err)
		}
		repos = append(repos, q.Viewer.Repositories.Nodes...)
		if !q.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(q.Viewer.Repositories.PageInfo.EndCursor)
	}
	return repos
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubql.NewClient(httpClient)

	repos := getRepositories(client)
	log.Printf("Got all repositories (%d)\n", len(repos))
}
