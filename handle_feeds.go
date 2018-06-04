package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleGetFeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	URL := fmt.Sprintf("https://github.com/%s/%s/commits/master.atom", vars["owner"], vars["name"])
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), resp.StatusCode)
		return
	}
	copyResponse(w, resp)
}
