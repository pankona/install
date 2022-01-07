package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
)

const currentDir = ""

func homeDir() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

type errContainer struct {
	err error
}

func (e *errContainer) execCommand(ctx context.Context, dir string, command ...string) {
	if e.err != nil {
		return
	}

	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	if dir != "" {
		err := os.Chdir(dir)
		if err != nil {
			e.err = fmt.Errorf("failed to chdir to %s: %w", dir, err)
			return
		}
	}
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Env = os.Environ()
	e.err = cmd.Run()
}

func withTempDir(pattern string, f func(workingDir string) error) error {
	tempDir, err := ioutil.TempDir("", pattern)
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer func() {
		os.RemoveAll(tempDir)
	}()

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current workng directory: %w", err)
	}

	if err := os.Chdir(tempDir); err != nil {
		return fmt.Errorf("failed to chdir: %w", err)
	}
	defer func() {
		if err := os.Chdir(currentDir); err != nil {
			log.Printf("failed to chdir: %v", err)
		}
	}()

	return f(tempDir)
}
