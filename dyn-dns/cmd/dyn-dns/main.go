package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Version injected at compile time.
var version = "No version provided"

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := realMain(ctx)
	done()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	// TODO: Add runner once implemented.
	return nil
}
