package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/signalfx/signalfx-go-tracing/tracing"
)

func main() {
	l := log.New(os.Stdout, "", 0)

	tracing.Start()
	defer tracing.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	go func() {
		errCh <- ServeSplunk(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}
