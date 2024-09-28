package main

import (
	"context"
	"log"
)

func installGitHubCLI(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "apt", "update", "-y")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "gh", "-y")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing github cli succeeded")
}
