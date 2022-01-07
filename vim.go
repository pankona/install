package main

import (
	"context"
	"log"
	"path/filepath"
)

func installVim(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "update")
	ec.execCommand(ctx, currentDir, "sudo", "-S", "apt", "install", "vim", "-y")

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing vim succeeded")
}

func installVimrc(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "ghq", "get", "pankona/vimrc")
	ec.execCommand(ctx, currentDir, filepath.Join(homeDir(), "go", "src", "github.com", "pankona", "vimrc", "install.sh"))

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing vimrc succeeded")

}

func installAg(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "apt", "update")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "silversearcher-ag")

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing ag succeeded")
}
