package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/patrickmn/go-cache"
	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

var (
	httpClient *http.Client
	client     *githubql.Client

	c = cache.New(3*time.Minute, 3*time.Minute)
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient = oauth2.NewClient(context.Background(), src)
	httpClient.Timeout = 10 * time.Second
	client = githubql.NewClient(httpClient)

	r := mux.NewRouter()

	r.HandleFunc("/repos", handleGetRepos).Methods("GET")
	r.HandleFunc("/repos", handlePatchRepo).Methods("POST")
	r.HandleFunc("/feed/{owner}/{name}", handleGetFeed).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	h := c.Handler(r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), h))
}
