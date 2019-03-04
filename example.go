package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dcos/client-go/dcos"
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
}
