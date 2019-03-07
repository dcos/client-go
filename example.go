package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
	"github.com/dcos/client-go/dcos/secrets/client/secrets"
	secretsmodels "github.com/dcos/client-go/dcos/secrets/models"

	cosmosclient "github.com/dcos/client-go/dcos/cosmos/client"
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

func getSecret(client *dcos.Client, secretName string) error {
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

func installPackage(client *dcos.Client, packageName string, options map[string]map[string]interface{}) error {
	paramsPackageInstall := cosmosclient.InstallRequest{
		PackageName: packageName,
		Options:     options,
	}

	result, _, err := client.Cosmos.DefaultApi.PackageInstall(context.TODO(), paramsPackageInstall)
	if err != nil {
		switch e := err.(type) {
		case cosmosclient.GenericOpenAPIError:
			if e.Model() != nil {
				switch m := e.Model().(type) {
				case cosmosclient.ModelError:
					log.Printf(m.Message)
					return nil
				default:
					return err
				}
			}
			return err
		default:
			return err
		}
	}

	log.Printf("Package installed: %+v\n", result.AppId)

	return nil
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: go run example.go SERVER_URL USERNAME PASSWORD")
	}

	config := dcos.NewConfig(nil)
	config.SetURL(os.Args[1])

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

	token, err := client.Login(os.Args[2], os.Args[3])
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

	jenkinsOptions := map[string]map[string]interface{}{
		"security": map[string]interface{}{
			"strictMode": true,
		},
	}

	err = installPackage(authClient, "jenkins", jenkinsOptions)
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}
}
