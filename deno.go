package main

import (
	"context"
	"log"
	"os/exec"
)

func installDeno(ctx context.Context) {
	// curl -fsSL https://deno.land/x/install/install.sh | sh
	cmd := exec.CommandContext(ctx, "curl", "-fsSL", "https://deno.land/x/install/install.sh")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to get kubectl latest version number: %v", err)
	}

	log.Printf("installing deno succeeded")
}
