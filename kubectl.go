package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func installKubectl(ctx context.Context) {
	cmd := exec.CommandContext(ctx, "curl", "-s", "https://storage.googleapis.com/kubernetes-release/release/stable.txt")
	version, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to get kubectl latest version number: %v", err)
	}

	cmd = exec.CommandContext(ctx,
		"curl", "-L", fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/linux/amd64/kubectl", string(version)),
		"-o", filepath.Join(homeDir(), "bin", "kubectl"))
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to get kubectl: %v", err)
	}

	if err = os.Chmod(filepath.Join(homeDir(), "bin", "kubectl"), 0o755); err != nil {
		log.Fatalf("failed to chmod kubectl: %v", err)
	}

	log.Printf("installing kubectl succeeded")
}

func installArgoRollouts(ctx context.Context) {
	cmd := exec.CommandContext(ctx,
		"curl", "-L", "https://github.com/argoproj/argo-rollouts/releases/latest/download/kubectl-argo-rollouts-linux-amd64",
		"-o", filepath.Join(homeDir(), "bin", "kubectl-argo-rollouts"))
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to get argo rollouts: %v", err)
	}
	if err := os.Chmod(filepath.Join(homeDir(), "bin", "kubectl-argo-rollouts"), 0o755); err != nil {
		log.Fatalf("failed to chmod argo rollouts: %v", err)
	}

	log.Printf("installing argo rollouts succeeded")
}
