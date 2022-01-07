package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
)

const goVersion = "1.17.6"

func installGouse(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "ghq", "get", "pankona/gouse")
	ec.err = nil // ignore if already exists
	if err := os.MkdirAll(filepath.Join(homeDir(), "bin"), os.ModeDir); err != nil {
		log.Fatalf("failed to create directory %s: %v", filepath.Join(homeDir(), "bin"), err)
	}
	if err := os.RemoveAll(filepath.Join(homeDir(), "bin", "gouse")); err != nil {
		log.Fatalf("failed to remove %s: %v", filepath.Join(homeDir(), "bin", "gouse"), err)
	}
	ec.execCommand(ctx, currentDir, "ln", "-s",
		filepath.Join(homeDir(), "go", "src", "github.com", "pankona", "gouse", "gouse"),
		filepath.Join(homeDir(), "bin", "gouse"),
	)

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing gouse succeeded")
}

func installGo(ctx context.Context) {
	ec := errContainer{}
	_ = withTempDir("install_go", func(workingDir string) error {
		ec.execCommand(ctx, workingDir, "wget", "https://go.dev/dl/go"+goVersion+".linux-amd64.tar.gz")
		ec.execCommand(ctx, workingDir, "tar", "zxf", "go"+goVersion+".linux-amd64.tar.gz")
		ec.execCommand(ctx, workingDir, filepath.Join("go", "bin", "go"), "install", "golang.org/dl/go"+goVersion+"@latest")
		return ec.err
	})
	ec.execCommand(ctx, currentDir, "go"+goVersion, "download")
	ec.execCommand(ctx, currentDir, "gouse", "go"+goVersion)
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing go succeeded")
}

func installGoImports(ctx context.Context) {
	ec := errContainer{}
	ec.execCommand(ctx, currentDir, "go", "install", "golang.org/x/tools/cmd/goimports@latest")
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing goimports succeeded")
}
