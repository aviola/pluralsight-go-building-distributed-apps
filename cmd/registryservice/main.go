package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aviola/pluralsight-go-building-distributed-apps/registry"
)

/*
Here, we're gonna use an implementation similar to the service.go of the logservice.
However, we won't reuse that service.go because it's eventually going to have functionality
specifically designed to handle client services. So the registry service itself won't be able
to take advantage of that.
*/
func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

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
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)

		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()

	fmt.Println("Shutting down registry service")
}
