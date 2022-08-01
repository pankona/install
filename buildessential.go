package main

import (
	"context"
	"log"
)

func installBuildEssential(ctx context.Context) {
	ec := errContainer{}
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "unzip")
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "libz-dev")
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "libssl-dev")
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "build-essential")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing build-essential succeeded")
}
