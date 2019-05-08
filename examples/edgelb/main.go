package main

import (
	"context"
	"log"

	"github.com/dcos/client-go/dcos"
)

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// Ping EdgeLB
	pong, _, err := client.Edgelb.Ping(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)
}
