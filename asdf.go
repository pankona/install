package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
)

const asdfVersion = "v0.18.0"

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

	// v0.18.0+ requires building binary instead of sourcing shell scripts
	ec.execCommand(ctx, asdfInstallDir, "make")

	binDir := filepath.Join(homeDir(), "bin")
	os.MkdirAll(binDir, 0755)

	asdfBinary := filepath.Join(asdfInstallDir, "asdf")
	asdfBinaryDest := filepath.Join(binDir, "asdf")
	ec.execCommand(ctx, currentDir, "cp", asdfBinary, asdfBinaryDest)

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing asdf succeeded")
}

type tool struct {
	name    string
	version string
}

func installRuby(ctx context.Context) {
	tools := []tool{
		{name: "ruby", version: "latest"},
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

	log.Printf("installing ruby via asdf succeeded")

}

func installNodejs(ctx context.Context) {
	tools := []tool{
		{name: "nodejs", version: "latest"},
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

	log.Printf("installing nodejs via asdf succeeded")

}

func installYarn(ctx context.Context) {
	tools := []tool{
		{name: "yarn", version: "latest"},
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

	log.Printf("installing yarn via asdf succeeded")

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
