repocmd
=======

> To learn how Go, Vue, and Electron can play together, see: [RepoCMD - An Adventure into Electron](https://eddieflores.com/tech/repocmd).

This is a standalone Go server containing all business logic used by [repocmd-desktop](https://github.com/ef2k/repocmd-desktop), an Electron app.

## To install:

```
$ go get github.com/ef2k/repocmd
```

## Running the server:

```
$ PORT=3000 GITHUB_TOKEN=<TOKEN> go run main.go
```

## Expected env variables:

```bash
#.env
GITHUB_TOKEN=<YOUR_TOKEN_HERE>
PORT=3000
```

> If using `codegangsta/gin`, it'll load your `.env` variables automatically.


## Vendored Dependencies:

```
+ github.com/shurcooL/githubql  - Used to request Github's GraphQL API
+ github.com/patrickmn/go-cache - Used to cache GitHub API calls
+ github.com/mmcdole/gofeed     - Parse Atom feeds
```

## Endpoints:

```
[GET]    /repos  - A list of repositories

[PATCH]  /repos - Patch a repository (only archiving is supported)

                  Expected JSON Body:

                      {
                        nameWithOwner: "ef2k/repocmd",
                        isArchived: <true or false>
                      }
```
