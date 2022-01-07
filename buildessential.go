package main

import (
	"context"
	"log"
)

func installBuildEssential(ctx context.Context) {
	ec := errContainer{}
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "build-essential")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing build-essential succeeded")
}
