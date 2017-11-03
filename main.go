package main

import (
	"context"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server"
)

func main() {
	c := context.Background()
	if err := server.Run(c); err != nil {
		panic(err)
	}
}
