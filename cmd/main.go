package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
	"github.com/dcos/client-go/dcos/cosmos"
	cosmostypes "github.com/dcos/client-go/dcos/cosmos/types"
)

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	cosmosPackageDescribeParams := cosmos.NewPackageDescribeParams()
	cosmosPackageDescribeParams.Body = &cosmostypes.PackageDescribeRequest{PackageName: &os.Args[1]}

	cosmosPackageDescribeOK, err := client.Cosmos.PackageDescribe(cosmosPackageDescribeParams, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Version %s\n", *cosmosPackageDescribeOK.Payload.Package.Version)
}
