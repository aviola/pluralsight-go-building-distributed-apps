package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// service is a generic module that takes info about the service, starts it up and handles graceful shutdown.

func Start(ctx context.Context, serviceName, host, port string, registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	// 1) goroutine to start the server
	go func() {
		// ListenAndServe() blocks as long as the server is running.
		// the console printout occurs only when ListenAndServe() returns.
		log.Println(srv.ListenAndServe())

		// if ListenAndServe() returns, it means that an error has occurred, so we need to cancel the context.
		cancel()
	}()

	// 2) goroutine to give us a cancellation option
	go func() {
		fmt.Printf("%v started. Press any key to stop.\n", serviceName)
		var s string
		fmt.Scanln(&s)

		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
