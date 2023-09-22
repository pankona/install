package main

import (
	"context"
	"log"

	"github.com/mattn/go-pipeline"
)

func installDeno(ctx context.Context) {
	// curl -fsSL https://deno.land/x/install/install.sh | sh
	_, err := pipeline.Output(
		[]string{"curl", "-fsSL", "https://deno.land/x/install/install.sh"},
		[]string{"sh"},
	)
	if err != nil {
		log.Fatalf("failed install deno: %v", err)
	}

	log.Printf("installing deno succeeded")
}
