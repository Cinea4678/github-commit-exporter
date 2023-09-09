package resolver

import (
	"context"
	"github.com/google/go-github/v55/github"
	"log"
	"strings"
)

func getAllCommits(owner string, repo string) []*github.RepositoryCommit {
	ctx := context.Background()

	r := make([]*github.RepositoryCommit, 0)

	for i := 1; true; i++ {
		log.Printf("resolve: Getting commits page #%d...", i)
		opt := github.CommitsListOptions{
			ListOptions: github.ListOptions{
				Page:    i,
				PerPage: 100,
			},
		}
		cl, _, err := client.Repositories.ListCommits(ctx, owner, repo, &opt)
		if err != nil {
			panic(err.Error())
		}
		if len(cl) == 0 {
			break
		}
		r = append(r, cl...)
	}

	return r
}

func getRepo(owner string, repo string) *github.Repository {
	ctx := context.Background()

	rp, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		panic(err.Error())
	}

	return rp
}

func ResolveCommits(repository string) []*github.RepositoryCommit {
	log.Println("resolve: Getting all commits...")
	i := strings.IndexRune(repository, '/')
	owner := repository[:i]
	repo := repository[i+1:]

	commits := getAllCommits(owner, repo)
	return commits
}

func ResolveRepoInfo(repository string) *github.Repository {
	log.Println("resolve: Getting repository information...")
	i := strings.IndexRune(repository, '/')
	owner := repository[:i]
	repo := repository[i+1:]

	r := getRepo(owner, repo)
	return r
}
