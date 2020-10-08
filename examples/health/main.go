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
	if err != nil {
		log.Fatal(err)
	}

	h, _, err := client.Health.V1SystemHealth(context.TODO())
	if err != nil {
		fmt.Printf("Error requesting SystemHealth %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("DC/OS Version %s, Diagnostics Version %s\n", h.DcosVersion, h.DcosDiagnosticsVersion)

	nodes, _, err := client.Health.V1SystemHealthNodes(context.TODO())
	if err != nil {
		fmt.Printf("Error requesting SystemHealthNodes %v\n", err)
		os.Exit(1)
	}

	fmt.Println("========== NODES ============")

	for _, node := range nodes.Nodes {
		fmt.Printf("Node: %s, Role: %s, Health: %d\n", node.HostIp, node.Role, node.Health)
	}

	units, _, err := client.Health.V1SystemHealthUnits(context.TODO())
	if err != nil {
		fmt.Printf("Error requesting SystemHealthNodes %v\n", err)
		os.Exit(1)
	}

	fmt.Println("========== Units ============")

	for _, unit := range units.Units {
		fmt.Printf("Unit: %s, Health: %d\n", unit.Name, unit.Health)
	}
}
