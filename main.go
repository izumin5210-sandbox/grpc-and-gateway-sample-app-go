package main

import (
	"context"

	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/server"
	"github.com/izumin5210-sandbox/grpc-and-gateway-sample-app-go/type/system"
)

func main() {
	c := context.Background()
	appC, err := system.CreateAppContext()

	if err != nil {
		panic(err)
	}

	if err := server.Run(c, appC); err != nil {
		panic(err)
	}
}
