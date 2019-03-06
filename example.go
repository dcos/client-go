package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
	"github.com/dcos/client-go/dcos/secrets/client/secrets"
	secretsmodels "github.com/dcos/client-go/dcos/secrets/models"

	"github.com/dcos/client-go/dcos/cosmos"
	cosmosoperations "github.com/dcos/client-go/dcos/cosmos/client/operations"
	cosmosmodels "github.com/dcos/client-go/dcos/cosmos/models"
)

const serverURL = "http://tball-client-go-0-23411561.us-east-1.elb.amazonaws.com"

func listUsers(client *dcos.Client) error {
	users, err := client.IAM.Users.GetUsers(nil)
	if err != nil {
		return err
	}

	log.Println("Listing users...")

	for _, u := range users.Payload.Array {
		fmt.Printf("\tUser %q, description=%q\n", *u.UID, *u.Description)
	}

	return nil
}

func createSecret(client *dcos.Client, secretName, secretValue string) error {
	paramsSecret := secrets.
		NewCreateSecretParams().
		WithStore("default").
		WithPathToSecret(secretName).
		WithBody(&secretsmodels.Secret{Value: secretValue})

	result, err := client.Secrets.Secrets.CreateSecret(paramsSecret)
	if err != nil {
		switch err.(type) {
		case *secrets.CreateSecretConflict:
			log.Printf("Secret %q already created", secretName)
		default:
			return err
		}
	}

	log.Printf("Secret created: %+v\n", result)

	return nil
}

func getSecret(client *dcos.Client, secretName string) error {
	paramsSecret := secrets.
		NewGetSecretParams().
		WithStore("default").
		WithPathToSecret(secretName)

	result, err := client.Secrets.Secrets.GetSecret(paramsSecret)
	if err != nil {
		return err
	}

	log.Printf("Secret fetched: %+v\n", result)

	return nil
}

func installPackage(client *dcos.Client, packageName string, options *cosmos.PackageOptions) error {
	paramsPackageInstall := cosmosoperations.
		NewPackageInstallParams().
		WithBody(&cosmosmodels.InstallRequest{
			PackageName: &packageName,
			Options:     options,
		})

	result, err := client.Cosmos.Operations.PackageInstall(paramsPackageInstall)
	if err != nil {
		switch err.(type) {
		case *cosmosoperations.PackageInstallConflict:
			log.Printf("Package %q is already installed", packageName)
			return nil
		default:
			return err
		}
	}

	log.Printf("Package installed: %+v\n", result.Payload.AppID)

	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: go run example.org USERNAME PASSWORD")
	}

	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		log.Fatalf("Creating new HTTP client failed: %s\n", err)
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		log.Fatalf("Creating new client failed: %s\n", err)
	}

	err = listUsers(client)
	if err != nil {
		log.Printf("Listing users WITHOUT token failed: %s\n", err)
	}

	token, err := client.Login(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("Login failed: %s\n", err)
	}

	config.SetACSToken(token)

	httpClient, err = dcos.NewHTTPClient(config)
	if err != nil {
		log.Fatalf("Creating new HTTP client failed: %s\n", err)
	}

	authClient, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		log.Fatalf("Creating new client failed: %s\n", err)
	}

	log.Printf("Login successful")

	err = listUsers(authClient)
	if err != nil {
		log.Fatalf("Listing users failed: %s\n", err)
	}

	err = createSecret(authClient, "biggestsecret3", "i am a dog")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = getSecret(authClient, "biggestsecret3")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = installPackage(authClient, "marathon-lb", nil)
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}

	jenkinsOptions := &cosmos.PackageOptions{
		Security: cosmos.PackageSecurityOptions{
			StrictMode: true,
		},
	}

	err = installPackage(authClient, "jenkins", jenkinsOptions)
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}
}
