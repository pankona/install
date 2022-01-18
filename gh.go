package main

import (
	"context"
	"log"
)

func installGitHubCLI(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "apt-key", "adv", "--keyserver", "keyserver.ubuntu.com", "--recv-key", "C99B11DEB97541F0")
	ec.execCommand(ctx, currentDir, "sudo", "install", "software-properties-common")
	ec.execCommand(ctx, currentDir, "sudo", "apt-add-repository", "https://cli.github.com/packages", "-y")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "update")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "gh")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing github cli succeeded")
}
