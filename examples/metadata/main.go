package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
)

func main() {
	client, err := dcos.NewClient()
	ctx := context.TODO()
	if err != nil {
		log.Fatal(err)
	}

	metadata, _, err := client.Metadata.Metadata(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ClusterID: %s  Public IPv4 %s", metadata.CLUSTER_ID, metadata.PUBLICIPV4)
	os.Exit(0)
}
