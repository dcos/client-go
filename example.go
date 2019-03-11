package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/dcos/client-go/dcos"
)

func listUsers(client *dcos.APIClient) error {
	users, _, err := client.IAMApi.GetUsers(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("Listing users...")

	for _, u := range users.Array {
		fmt.Printf("\tUser %q, description=%q\n", u.Uid, u.Description)
	}

	return nil
}

func login(clusterURL, username, password string) (string, error) {
	// Empty config, without auth token
	config := dcos.NewConfig(nil)
	config.SetURL(clusterURL)

	client, err := dcos.NewClientWithConfig(config)
	if err != nil {
		return "", err
	}

	loginObject := dcos.LoginObject{Uid: username, Password: password}

	authToken, _, err := client.IAMApi.Login(context.TODO(), loginObject)
	if err != nil {
		return "", err
	}

	return authToken.Token, nil
}

func createSecret(client *dcos.APIClient, secretName, secretValue string) error {
	secret := dcos.Secret{Value: secretValue}

	_, err := client.SecretsApi.CreateSecret(context.TODO(), "default", secretName, secret)
	if err != nil {
		switch err := err.(type) {
		case dcos.GenericOpenAPIError:
			if err.Error() == "409 Conflict" {
				log.Printf("Secret %s already created", secretName)
				return nil
			}
			return err
		default:
			return err
		}
	}

	log.Printf("Secret created: %+v\n", secretName)

	return nil
}

func getSecret(client *dcos.APIClient, secretName string) error {
	secret, _, err := client.SecretsApi.GetSecret(context.TODO(), "default", secretName, nil)
	if err != nil {
		return err
	}

	log.Printf("Secret fetched: %+v\n", secret)

	return nil
}

func describePackage(client *dcos.APIClient, request dcos.PackageDescribeRequest) error {
	result, _, err := client.CosmosApi.PackageDescribe(context.TODO(), &dcos.PackageDescribeOpts{
		PackageDescribeRequest: optional.NewInterface(request),
	})
	if err != nil {
		return err
	}
	log.Printf("Package %s description: %+v\n", request.PackageName, result.Package.Description)

	return nil
}

func installPackage(client *dcos.APIClient, request dcos.InstallRequest) error {
	result, _, err := client.CosmosApi.PackageInstall(context.TODO(), request)
	if err != nil {
		switch err := err.(type) {
		case dcos.GenericOpenAPIError:
			if err.Error() == "409 Conflict" {
				log.Printf("Package %s already installed", request.PackageName)
				return nil
			}
			return err
		default:
			return err
		}
	}

	log.Printf("Package installed: %+v\n", result)

	return nil
}

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	err = listUsers(client)
	if err != nil {
		log.Printf("Listing users WITHOUT token failed: %s\n", err)
	}

	if len(os.Args) > 3 {
		token, err := login(os.Args[1], os.Args[2], os.Args[3])
		if err != nil {
			log.Fatalf("Login failed: %s\n", err)
		}

		log.Printf("Logged in and received token: %s", token)
	}

	err = createSecret(client, "biggestsecret3", "i am a dog")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = getSecret(client, "biggestsecret3")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = describePackage(client, dcos.PackageDescribeRequest{PackageName: "hello-world"})
	if err != nil {
		log.Fatalf("Describing package failed: %s\n", err)
	}

	err = installPackage(client, dcos.InstallRequest{PackageName: "hello-world"})
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}

	dcosMonitoringRequest := dcos.InstallRequest{
		PackageName: "beta-dcos-monitoring",
		Options: map[string]map[string]interface{}{
			"grafana": {
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
}
