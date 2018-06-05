package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mmcdole/gofeed/atom"
	"github.com/shurcooL/githubql"
)

type Feeds struct {
	Releases *atom.Feed `json:"releases"`
	Commits  *atom.Feed `json:"commits"`
}

type RepoResponse struct {
	Summary *RepoSummary `json:"summary"`
	Feeds   *Feeds       `json:"feeds"`
}

type RepoSummary struct {
	ID          githubql.ID      `json:"id"`
	Description githubql.String  `json:"description"`
	URL         githubql.URI     `json:"url"`
	IsArchived  githubql.Boolean `json:"isArchived"`
	IsFork      githubql.Boolean `json:"isFork"`
	IsPrivate   githubql.Boolean `json:"isPrivate"`
	Stargazers  totalCount       `json:"stargazers"`
	Watchers    totalCount       `json:"watchers"`
	Forks       totalCount       `json:"forks"`
	SSHURL      githubql.String  `json:"sshURL"`
}

func getRepoSummary(client *githubql.Client, owner, name string) (*RepoSummary, error) {
	var repoQ struct {
		Repository RepoSummary `graphql:"repository(owner: $owner, name: $name)"`
	}

	variables := map[string]interface{}{
		"owner": (githubql.String)(owner),
		"name":  (githubql.String)(name),
	}

	if err := client.Query(context.Background(), &repoQ, variables); err != nil {
		return nil, err
	}
	return &repoQ.Repository, nil
}

func parseFeed(reader io.Reader) (*atom.Feed, error) {
	fp := atom.Parser{}
	return fp.Parse(reader)
}

func handleGetFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	name := vars["name"]

	// Get releases
	releaseURL := fmt.Sprintf("https://github.com/%s/%s/releases.atom", owner, name)
	releaseReq, err := http.NewRequest("GET", releaseURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	releaseResp, err := httpClient.Do(releaseReq)
	if err != nil {
		http.Error(w, err.Error(), releaseResp.StatusCode)
		return
	}
	releases, err := parseFeed(releaseResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get recent commits
	commitsURL := fmt.Sprintf("https://github.com/%s/%s/commits/master.atom", owner, name)
	commitsReq, err := http.NewRequest("GET", commitsURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	commitsResp, err := httpClient.Do(commitsReq)
	if err != nil {
		http.Error(w, err.Error(), commitsResp.StatusCode)
		return
	}
	commits, err := parseFeed(commitsResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get repository summary
	summary, err := getRepoSummary(client, owner, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repo := RepoResponse{
		Feeds: &Feeds{
			Releases: releases,
			Commits:  commits,
		},
		Summary: summary,
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(repo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
