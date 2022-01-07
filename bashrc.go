package main

import (
	"context"
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

//go:embed mybashrc
var myBashrc []byte

func installBashrc(ctx context.Context) {
	homeDir := homeDir()

	if err := os.WriteFile(filepath.Join(homeDir, "mybashrc"), myBashrc, 0664); err != nil {
		log.Fatalf("failed to write mybashrc: %v", err)
	}

	cmd := exec.CommandContext(ctx, "grep", "mybashrc", filepath.Join(homeDir, ".bashrc"))
	if err := cmd.Run(); err == nil {
		// grep returned 0. It means .bashrc already includes mybashrc
		return
	}

	f, err := os.OpenFile(filepath.Join(homeDir, ".bashrc"), os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open bashrc: %v", err)
	}
	cmd = exec.CommandContext(ctx, "echo", "\n. $HOME/mybashrc")
	cmd.Stdout = f
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to redirect mybashrc to .bashrc: %v", err)
	}

	log.Printf("installing bashrc succeeded")
}
