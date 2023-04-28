package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/aviola/pluralsight-go-building-distributed-apps/log"
	"github.com/aviola/pluralsight-go-building-distributed-apps/registry"
	"github.com/aviola/pluralsight-go-building-distributed-apps/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(), host, port, r, log.RegisterHandlers)
	if err != nil {
		// Here, the log service has crashed, so the only way to log is through the standard logger.
		stlog.Fatal(err)
	}

	// ctx is Done when cancel() is called in one of the goroutines of the service (see service.go).
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
