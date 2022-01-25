package main

import (
	"context"
	"log"
	"path/filepath"
)

func installFzf(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "ghq", "get", "junegunn/fzf")
	ec.execCommand(ctx, currentDir, filepath.Join(homeDir(), "go", "src", "github.com", "junegunn", "fzf", "install"), "--all")

	log.Printf("installing fzf succeeded")
}
