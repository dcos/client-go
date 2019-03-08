package main

import (
	"context"
	"fmt"
	"log"

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

/*
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
*/

func installPackage(client *dcos.APIClient, packageName string) error {
	result, _, err := client.CosmosApi.PackageInstall(context.TODO(), dcos.InstallRequest{
		PackageName: packageName,
	})
	if err != nil {
		switch err := err.(type) {
		case dcos.GenericOpenAPIError:
			if err.Error() == "409 Conflict" {
				log.Printf("Package %s already installed", packageName)
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

	/*
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
	*/

	err = installPackage(client, "jenkins")
	if err != nil {
		log.Fatalf("Installing package failed: %s\n", err)
	}
}
