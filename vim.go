package main

import (
	"context"
	"log"
	"path/filepath"
)

func installVim(ctx context.Context) {
	ec := errContainer{}
	vimDir := filepath.Join(homeDir(), "go", "src", "github.com", "vim", "vim")

	ec.execCommand(ctx, currentDir, "ghq", "get", "vim/vim")
	
	// Configure to install in $HOME/bin instead of system-wide /usr/local
	ec.execCommand(ctx, vimDir, "./configure", "--prefix="+homeDir())
	ec.execCommand(ctx, vimDir, "make")
	ec.execCommand(ctx, vimDir, "make", "install")

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
