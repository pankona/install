package main

import (
	"context"
	"log"
)

func installPrettier(ctx context.Context) {
	ec := errContainer{}
	ec.execCommand(ctx, currentDir, "yarn", "global", "add", "prettier")
	ec.execCommand(ctx, currentDir, "asdf", "reshim")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing prettier succeeded")
}
