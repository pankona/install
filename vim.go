package main

import (
	"context"
	"log"
	"path/filepath"
)

func installDein(ctx context.Context) {
	// 以下のコマンドを実行する
	// sh -c "$(curl -fsSL https://raw.githubusercontent.com/Shougo/dein-installer.vim/master/installer.sh)"
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sh", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Shougo/dein-installer.vim/master/installer.sh)")

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing dein succeeded")
}

func installVim(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "ghq", "get", "vim/vim")
	ec.execCommand(ctx, filepath.Join(homeDir(), "go", "src", "github.com", "vim", "vim"), "make")
	ec.execCommand(ctx, filepath.Join(homeDir(), "go", "src", "github.com", "vim", "vim"), "sudo", "make", "install")

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
