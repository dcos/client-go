package main

import (
	"context"
	"fmt"
	"log"

	"github.com/antihax/optional"
	"github.com/dcos/client-go/dcos"
)

func describePackage(client *dcos.APIClient, request dcos.CosmosPackageDescribeV1Request) error {
	result, _, err := client.Cosmos.PackageDescribe(context.TODO(), &dcos.PackageDescribeOpts{
		CosmosPackageDescribeV1Request: optional.NewInterface(request),
	})
	if err != nil {
		return err
	}
	log.Printf("Package %s description: %+v\n", request.PackageName, result.Package.Description)

	return nil
}

func installPackage(client *dcos.APIClient, request dcos.CosmosPackageInstallV1Request) error {
	result, _, err := client.Cosmos.PackageInstall(context.TODO(), request)
	if err != nil {
		switch err := err.(type) {
		case dcos.GenericOpenAPIError:
			if err.Model() != nil {
				if e, ok := err.Model().(dcos.CosmosError); ok && e.Type == "PackageAlreadyInstalled" {
					log.Printf("Package %s already installed", request.PackageName)
					return nil
				}
			}
			return err
		default:
			return err
		}
	}

	log.Printf("Package installed: %+v\n", result)

	return nil
}

func listPackages(client *dcos.APIClient) error {
	result, _, err := client.Cosmos.PackageList(context.TODO(), &dcos.PackageListOpts{
		CosmosPackageListV1Request: optional.NewInterface(dcos.CosmosPackageListV1Request{}),
	})
	if err != nil {
		return err
	}
	fmt.Println("Installed packages:")
	for _, pkg := range result.Packages {
		fmt.Println(" - " + pkg.PackageInformation.PackageDefinition.Name)
	}
	return nil
}

func uninstallPackage(client *dcos.APIClient, request dcos.CosmosPackageUninstallV1Request) error {
	result, _, err := client.Cosmos.PackageUninstall(context.TODO(), request, nil)
	if err != nil {
		return err
	}

	log.Printf("Package uninstalled: %+v\n", result)

	return nil
}

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	err = describePackage(client, dcos.CosmosPackageDescribeV1Request{PackageName: "hello-world"})
	if err != nil {
		log.Fatalf("Describing package failed: %s\n", err)
	}

	err = installPackage(client, dcos.CosmosPackageInstallV1Request{PackageName: "hello-world"})
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}

	err = listPackages(client)
	if err != nil {
		log.Fatalf("Listing packages failed: %s\n", err)
	}

	err = uninstallPackage(client, dcos.CosmosPackageUninstallV1Request{PackageName: "hello-world"})
	if err != nil {
		log.Fatalf("Uninstalling package failed: %s\n", err)
	}

	dcosMonitoringRequest := dcos.CosmosPackageInstallV1Request{
		PackageName: "beta-dcos-monitoring",
		Options: map[string]interface{}{
			"grafana": map[string]interface{}{
				"default_dashboards": false,
				"dashboard_config_repository": map[string]string{
					"url":  "https://github.com/masonse/grafana-dashboards",
					"path": "/dashboards",
				},
			},
		},
	}

	err = installPackage(client, dcosMonitoringRequest)
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}

	err = listPackages(client)
	if err != nil {
		log.Fatalf("Listing packages failed: %s\n", err)
	}

	err = uninstallPackage(client, dcos.CosmosPackageUninstallV1Request{PackageName: "beta-dcos-monitoring"})
	if err != nil {
		log.Fatalf("Uninstalling package failed: %s\n", err)
	}
}
