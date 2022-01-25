package main

import (
	"context"
	"log"
)

func installAptTools(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "apt", "-y", "update")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "silversearcher-ag")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "less")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing ag succeeded")
}
