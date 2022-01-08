package main

import (
	"context"
	"log"
)

var gitConfigs = map[string]string{
	"user.email":  "yosuke.akatsuka@gmail.com",
	"user.name":   "pankona",
	"core.editor": "vim",
	"pull.ff":     "only",
	"ghq.root":    "~/go/src",
}

// get git source code for git completion
func installGit(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "ghq", "get", "git/git")

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing git succeeded")
}

func installGitConfig(ctx context.Context) {
	ec := errContainer{}
	for configKey, configValue := range gitConfigs {
		ec.execCommand(ctx, currentDir, "git", "config", "--global", configKey, configValue)
	}
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing gitconfig succeeded")
}
