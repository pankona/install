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
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "locales-all")

	// 後々 apt install で必要になるものはここで適当に入れちゃう
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "unzip")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "libssl-dev")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "libz-dev")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "-y", "libexpat-dev")
	// https: //github.com/cli/cli/discussions/6222
	ec.execCommand(ctx, currentDir, "sudo", "apt-key", "adv", "--keyserver", "keyserver.ubuntu.com", "--recv-keys", "23F3D4EA75716059")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing ag succeeded")
}
