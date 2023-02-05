package main

import (
	"context"
	"fmt"
	"net/http"

	httptrace "github.com/signalfx/signalfx-go-tracing/contrib/net/http"
	"github.com/tullo/otel-workshop/web/fib"
)

func ServeSplunk(ctx context.Context) error {
	mux := httptrace.NewServeMux()
	mux.Handle("/", http.HandlerFunc(fib.RootHandler))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/fib", http.HandlerFunc(fib.FibHandler))
	mux.Handle("/fibinternal", http.HandlerFunc(fib.FibHandler))

	fmt.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
