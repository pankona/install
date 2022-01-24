package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
)

const asdfVersion = "v0.9.0"

func installAsdf(ctx context.Context) {
	asdfInstallDir := filepath.Join(homeDir(), ".asdf")
	_, err := os.Stat(asdfInstallDir)
	if err == nil {
		log.Printf("It seems asdf is already installed since " + asdfInstallDir + " directory exists. skip.")
		return
	}

	ec := errContainer{}
	ec.execCommand(ctx, currentDir,
		"git", "clone", "https://github.com/asdf-vm/asdf.git", asdfInstallDir, "--branch", asdfVersion)
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing asdf succeeded")
}

type tool struct {
	name    string
	version string
}

func installToolsViaAsdf(ctx context.Context) {
	tools := []tool{
		{name: "ruby", version: "3.1.0"},
		{name: "nodejs", version: "16.13.0"},
		{name: "yarn", version: "latest"},
		{name: "ghq", version: "latest"},
	}
	ec := errContainer{}
	for _, tool := range tools {
		ec.execCommand(ctx, currentDir, "asdf", "plugin-add", tool.name)
		ec.err = nil // ignore error to ignore "the plugin is already added"
		ec.execCommand(ctx, currentDir, "asdf", "install", tool.name, tool.version)
		ec.execCommand(ctx, currentDir, "asdf", "global", tool.name, tool.version)
	}
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing tools via asdf succeeded")
}

func installGHQ(ctx context.Context) {
	tools := []tool{
		{name: "ghq", version: "latest"},
	}
	ec := errContainer{}
	for _, tool := range tools {
		ec.execCommand(ctx, currentDir, "asdf", "plugin-add", tool.name)
		ec.err = nil // ignore error to ignore "the plugin is already added"
		ec.execCommand(ctx, currentDir, "asdf", "install", tool.name, tool.version)
		ec.execCommand(ctx, currentDir, "asdf", "global", tool.name, tool.version)
	}
	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing ghq via asdf succeeded")
}
