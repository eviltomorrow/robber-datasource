package main

import (
	"github.com/eviltomorrow/robber-core/pkg/system"
	"github.com/eviltomorrow/robber-datasource/internal/command"
)

var (
	GitSha      = ""
	GitTag      = ""
	GitBranch   = ""
	BuildTime   = ""
	MainVersion = "v3.0"
)

func init() {
	system.MainVersion = MainVersion
	system.GitSha = GitSha
	system.GitTag = GitTag
	system.GitBranch = GitBranch
	system.BuildTime = BuildTime
}

func main() {
	command.Execute()
}
