package resolver

import (
	"context"
	"github.com/gofri/go-github-ratelimit/github_ratelimit"
	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func Init(token string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	rateLimiter, err := github_ratelimit.NewRateLimitWaiterClient(tc.Transport)
	if err != nil {
		panic(err)
	}
	client = github.NewClient(rateLimiter)
}
