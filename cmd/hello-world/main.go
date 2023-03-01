package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Version injected at compile time.
var version = "No version provided"

type Config struct {
	Addr string `envconfig:"HELLO_WORLD_ADDR" default:":8080"`
}

func main() {
	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	err := realMain(ctx)
	done()
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatal(err)
	}
}

func realMain(ctx context.Context) error {
	log.Printf("Starting hello-world %s", version)
	var conf Config
	if err := envconfig.Process("HELLO_WORLD", &conf); err != nil {
		return err
	}

	srv := &http.Server{
		Addr:         conf.Addr,
		Handler:      handler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	errCh := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		if err := srv.Shutdown(context.Background()); err != nil {
			return err
		}
		log.Println("HTTP server stopped gracefully")
		return ctx.Err()
	case err := <-errCh:
		log.Printf("HTTP server stopped unexpectedly: %v", err)
		return err
	}
}

func handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Printf("Unable to split remote address: %s\n", r.RemoteAddr)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		log.Printf("Received request from %s\n", clientIP)
		fmt.Fprintf(w, "Hello %s @ %s", clientIP, time.Now().Format(time.RFC3339))
	}
}
