package main

import (
	"context"
	"sort"

	cache "github.com/patrickmn/go-cache"
	"github.com/shurcooL/githubql"
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

func getOwnerRepos(client *githubql.Client) ([]repository, error) {
	var ownerQ struct {
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

	cached, found := c.Get(OWNER)
	if found {
		return cached.(fromOldestTimeSlice), nil
	}

	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}
	var repos []repository
	for {
		if err := client.Query(context.Background(), &ownerQ, variables); err != nil {
			return nil, err
		}
		repos = append(repos, ownerQ.Viewer.Repositories.Nodes...)
		if !ownerQ.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(ownerQ.Viewer.Repositories.PageInfo.EndCursor)
	}
	timeSortedRepos := make(fromOldestTimeSlice, 0, len(repos))
	timeSortedRepos = append(timeSortedRepos, repos...)
	sort.Sort(timeSortedRepos)
	c.Set(OWNER, timeSortedRepos, cache.DefaultExpiration)
	return timeSortedRepos, nil
}

func getCollabRepos(client *githubql.Client) ([]repository, error) {
	var collabQ struct {
		Viewer struct {
			Repositories struct {
				Nodes    []repository
				PageInfo struct {
					EndCursor   githubql.String
					HasNextPage githubql.Boolean
				}
			} `graphql:"repositories(last:100, after:$repositoriesCursor, affiliations:[COLLABORATOR])"`
		}
	}

	cached, found := c.Get(COLLABORATOR)
	if found {
		return cached.(fromOldestTimeSlice), nil
	}

	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}
	var repos []repository
	for {
		if err := client.Query(context.Background(), &collabQ, variables); err != nil {
			return nil, err
		}
		repos = append(repos, collabQ.Viewer.Repositories.Nodes...)
		if !collabQ.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(collabQ.Viewer.Repositories.PageInfo.EndCursor)
	}
	timeSortedRepos := make(fromOldestTimeSlice, 0, len(repos))
	timeSortedRepos = append(timeSortedRepos, repos...)
	sort.Sort(timeSortedRepos)
	c.Set(COLLABORATOR, timeSortedRepos, cache.DefaultExpiration)
	return timeSortedRepos, nil
}

func getOrgRepos(client *githubql.Client) ([]repository, error) {
	var orgQ struct {
		Viewer struct {
			Repositories struct {
				Nodes    []repository
				PageInfo struct {
					EndCursor   githubql.String
					HasNextPage githubql.Boolean
				}
			} `graphql:"repositories(last:100, after:$repositoriesCursor, affiliations:[ORGANIZATION_MEMBER])"`
		}
	}

	cached, found := c.Get(ORGANIZATION_MEMBER)
	if found {
		return cached.(fromOldestTimeSlice), nil
	}

	variables := map[string]interface{}{
		"repositoriesCursor": (*githubql.String)(nil),
	}
	var repos []repository
	for {
		if err := client.Query(context.Background(), &orgQ, variables); err != nil {
			return nil, err
		}
		repos = append(repos, orgQ.Viewer.Repositories.Nodes...)
		if !orgQ.Viewer.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["repositoriesCursor"] = githubql.NewString(orgQ.Viewer.Repositories.PageInfo.EndCursor)
	}
	timeSortedRepos := make(fromOldestTimeSlice, 0, len(repos))
	timeSortedRepos = append(timeSortedRepos, repos...)
	sort.Sort(timeSortedRepos)
	c.Set(ORGANIZATION_MEMBER, timeSortedRepos, cache.DefaultExpiration)
	return timeSortedRepos, nil
}
