package main

import (
	"flag"
	"github-commit-exporter/pkg/exporter"
	"github-commit-exporter/pkg/resolver"
)

var repo = flag.String("repo", "", "the Github repository name")
var token = flag.String("token", "", "the Github Token")
var out = flag.String("out", "", "the output file path")

func main() {
	flag.Parse()

	resolver.Init(*token)

	commits := resolver.ResolveCommits(*repo)
	repository := resolver.ResolveRepoInfo(*repo)

	exporter.ExportToXlsx(repository, commits, *out)
}
