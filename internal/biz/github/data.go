package github

import "github.com/google/go-github/v45/github"

type GithubRepositoryRelease struct {
	Owner             string
	Repo              string
	RepositoryRelease *github.RepositoryRelease
}
