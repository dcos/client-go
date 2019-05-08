package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
)

func listUsers(client *dcos.APIClient) error {
	users, _, err := client.IAM.GetUsers(context.TODO(), nil)
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

	loginObject := dcos.IamLoginObject{Uid: username, Password: password}

	authToken, _, err := client.IAM.Login(context.TODO(), loginObject)
	if err != nil {
		return "", err
	}

	return authToken.Token, nil
}

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// List all users
	err = listUsers(client)
	if err != nil {
		log.Printf("Listing users WITHOUT token failed: %s\n", err)
	}

	// Login if username/password are passed as arguments
	if len(os.Args) > 3 {
		token, err := login(os.Args[1], os.Args[2], os.Args[3])
		if err != nil {
			log.Fatalf("Login failed: %s\n", err)
		}

		log.Printf("Logged in and received token: %s", token)
	}
}
