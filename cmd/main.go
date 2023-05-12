package main

import (
	"grpc_logging/internal/server"
	"log"
)

func main() {
	log.Fatal(run())
}

func run() error {

	loggingSrv := server.NewLoggingServer()
	srv := server.New(loggingSrv)

	if err := srv.ListenAndServe(9000); err != nil {
		return err
	}
	return nil
}
