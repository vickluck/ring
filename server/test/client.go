package main

import (
	"context"
	"log"

	"github.com/tnngo/nog"
	"github.com/tnngo/ring/server/proto/provider"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1500", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := provider.NewProviderClient(conn)

	r, err := c.Regist(context.Background(), &provider.Request{})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	nog.SL.Debug(r)
}
