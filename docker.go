package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/mattn/go-pipeline"
)

func installDocker(ctx context.Context) {
	ec := errContainer{}

	ec.execCommand(ctx, currentDir, "sudo", "apt", "update")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "ca-certificates", "curl", "gnupg", "lsb-release", "-y")
	ec.execCommand(ctx, currentDir, "sudo", "rm", "/usr/share/keyrings/docker-archive-keyring.gpg")
	_, ec.err = pipeline.Output(
		[]string{"curl", "-fsSL", "https://download.docker.com/linux/ubuntu/gpg"},
		[]string{"sudo", "gpg", "--dearmor", "-o", "/usr/share/keyrings/docker-archive-keyring.gpg"},
	)

	cmd := exec.CommandContext(ctx, "dpkg", "--print-architecture")
	dpkg, err := cmd.Output()
	if err != nil {
		log.Fatalf("dpkg --print-architecture failed: %v", err)
	}
	cmd = exec.CommandContext(ctx, "lsb_release", "-cs")
	lsbRelease, err := cmd.Output()
	if err != nil {
		log.Fatalf("lsb_release -cs failed: %v", err)
	}
	_, ec.err = pipeline.Output(
		[]string{"echo", fmt.Sprintf("deb [arch=%s signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu %s stable",
			strings.TrimSuffix(string(dpkg), "\n"),
			strings.TrimSuffix(string(lsbRelease), "\n"))},
		[]string{"sudo", "tee", "/etc/apt/sources.list.d/docker.list"},
	)

	ec.execCommand(ctx, currentDir, "sudo", "apt", "update")
	ec.execCommand(ctx, currentDir, "sudo", "apt", "install", "docker-ce", "docker-ce-cli", "containerd.io", "-y")

	user, err := user.Current()
	if err != nil {
		log.Fatal(ec.err)
	}
	ec.execCommand(ctx, currentDir, "sudo", "usermod", "-aG", "docker", user.Username)

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing docker succeeded")
}

const dockerComposeVersion = "v2.2.2"

func installDockerCompose(ctx context.Context) {
	ec := errContainer{}
	homeDir := homeDir()

	fmt.Println("https://github.com/docker/compose/releases/download/" + dockerComposeVersion + "/docker-compose-linux-x86_64")

	ec.execCommand(ctx, currentDir, "mkdir", "-p", filepath.Join(homeDir, ".docker", "cli-plugins"))
	ec.execCommand(ctx, currentDir, "curl", "-SL", "https://github.com/docker/compose/releases/download/"+dockerComposeVersion+"/docker-compose-linux-x86_64", "-o", filepath.Join(homeDir, ".docker", "cli-plugins", "docker-compose"))
	ec.execCommand(ctx, currentDir, "chmod", "+x", filepath.Join(homeDir, ".docker", "cli-plugins", "docker-compose"))

	if ec.err != nil {
		log.Fatal(ec.err)
	}

	log.Printf("installing docker-compose succeeded")
}
