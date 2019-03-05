package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
	"github.com/dcos/client-go/dcos/secrets/client/secrets"
	secretsmodels "github.com/dcos/client-go/dcos/secrets/models"

	cosmosoperations "github.com/dcos/client-go/dcos/cosmos/client/operations"
	cosmosmodels "github.com/dcos/client-go/dcos/cosmos/models"
)

const serverURL = "http://tball-client-go-0-1598387033.us-east-1.elb.amazonaws.com"

func Login(username, password string) (string, error) {
	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		return "", err
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		return "", err
	}

	token, err := client.Login(username, password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ListUsers(token string) error {
	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)
	config.SetACSToken(token)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		return err
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		return err
	}

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

func CreateSecret(token, secretName, secretValue string) error {
	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)
	config.SetACSToken(token)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		return err
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		return err
	}

	paramsSecret := secrets.
		NewPutSecretStorePathToSecretParams().
		WithStore("default").
		WithPathToSecret(secretName).
		WithBody(&secretsmodels.Secret{Value: secretValue})

	result, err := client.Secrets.Secrets.PutSecretStorePathToSecret(paramsSecret)
	if err != nil {
		switch err.(type) {
		case *secrets.PutSecretStorePathToSecretConflict:
			log.Printf("Secret %q already created", secretName)
		default:
			return err
		}
	}

	log.Printf("Secret created: %+v\n", result)

	return nil
}

func GetSecret(token, secretName string) error {
	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)
	config.SetACSToken(token)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		return err
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		return err
	}

	paramsSecret := secrets.
		NewGetSecretStorePathToSecretParams().
		WithStore("default").
		WithPathToSecret(secretName)

	result, err := client.Secrets.Secrets.GetSecretStorePathToSecret(paramsSecret)
	if err != nil {
		return err
	}

	log.Printf("Secret fetched: %+v\n", result)

	return nil
}

func InstallPackage(token, packageName string) error {
	config := dcos.NewConfig(nil)
	config.SetURL(serverURL)
	config.SetACSToken(token)

	httpClient, err := dcos.NewHTTPClient(config)
	if err != nil {
		return err
	}

	client, err := dcos.NewClientWithOptions(httpClient, config)
	if err != nil {
		return err
	}

	paramsPackageInstall := cosmosoperations.
		NewPackageInstallParams().
		WithBody(&cosmosmodels.InstallRequest{PackageName: &packageName})

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

	err := ListUsers("")
	if err != nil {
		log.Printf("Listing users WITHOUT token failed: %s\n", err)
	}

	token, err := Login(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("Login failed: %s\n", err)
	}

	log.Printf("Login successful")

	err = ListUsers(token)
	if err != nil {
		log.Fatalf("Listing users failed: %s\n", err)
	}

	err = CreateSecret(token, "biggestsecret3", "i am a dog")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = GetSecret(token, "biggestsecret3")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = InstallPackage(token, "marathon-lb")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}
}
