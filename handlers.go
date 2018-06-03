package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	OWNER               = "owner"
	COLLABORATOR        = "collaborator"
	ORGANIZATION_MEMBER = "organizationMember"
)

func handleGetRepos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	af := r.URL.Query().Get("affiliation")
	var repos []repository
	var err error
	if af == COLLABORATOR {
		repos, err = getCollabRepos(client)
	} else if af == ORGANIZATION_MEMBER {
		repos, err = getOrgRepos(client)
	} else {
		repos, err = getOwnerRepos(client)
	}
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
