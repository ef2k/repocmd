package main

import (
	"context"
	"os"

	"github.com/shurcooL/githubql"
	"golang.org/x/oauth2"
)

var q struct {
	Viewer struct {
		Repository struct {
			Limit githubql.Int `graphql:"repository(limit:\"30\")"`
		}
	}
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubql.NewClient(httpClient)
}
