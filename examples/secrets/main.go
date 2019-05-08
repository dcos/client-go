package main

import (
	"context"
	"log"

	"github.com/dcos/client-go/dcos"
)

func createSecret(client *dcos.APIClient, secretName, secretValue string) error {
	secret := dcos.SecretsV1Secret{Value: secretValue}

	_, err := client.Secrets.CreateSecret(context.TODO(), "default", secretName, secret)
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
	secret, _, err := client.Secrets.GetSecret(context.TODO(), "default", secretName, nil)
	if err != nil {
		return err
	}

	log.Printf("Secret fetched: %+v\n", secret)

	return nil
}

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	err = createSecret(client, "biggestsecret3", "i am a dog")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}

	err = getSecret(client, "biggestsecret3")
	if err != nil {
		log.Fatalf("Creating secrets failed: %s\n", err)
	}
}
