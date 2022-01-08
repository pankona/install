package main

import (
	"context"
	"log"
	"path/filepath"
)

func installPrettier(ctx context.Context) {
	ec := errContainer{}
	ec.execCommand(ctx, currentDir, "yarn", "global", "add", "prettier")
	ec.execCommand(ctx, currentDir, "bash", "-c", "source "+filepath.Join(homeDir(), ".bashrc")+"; asdf reshim")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing prettier succeeded")
}
