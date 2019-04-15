# DC/OS: Go Client

[![GoDoc](https://godoc.org/github.com/dcos/client-go/dcos?status.svg)](https://godoc.org/github.com/dcos/client-go/dcos) [![Build Status](https://travis-ci.org/dcos/client-go.svg?branch=master)](https://travis-ci.org/dcos/client-go)

**⚠️ This project is still experimental and shouldn't be used in production yet. ⚠️**

dcos-client-go is a library accessing the DC/OS API to provide clients for cluster operators. A working and up-to-date example using the clients is available in the file `example.go`.

# IAM client

```
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
````

# Secrets client

```
func getSecret(client *dcos.APIClient, secretName string) error {
	secret, _, err := client.Secrets.GetSecret(context.TODO(), "default", secretName, nil)
	if err != nil {
		return err
	}

	log.Printf("Secret fetched: %+v\n", secret)

	return nil
}
```

# Cosmos client

```
func describePackage(client *dcos.APIClient, request dcos.PackageDescribeRequest) error {
	result, _, err := client.Cosmos.PackageDescribe(context.TODO(), &dcos.PackageDescribeOpts{
		PackageDescribeRequest: optional.NewInterface(request),
	})
	if err != nil {
		return err
	}
	log.Printf("Package %s description: %+v\n", request.PackageName, result.Package.Description)

	return nil
}
```

# EdgeLB client

```
func edgeLBPing(client *dcos.APIClient) error {
	pong, _, err := client.Edgelb.Ping(context.TODO())
	if err != nil {
		return err
	}
	log.Println(pong)
	return nil
}
```

# Marathon client

We currently do not use our own implementation of a Marathon client and rely on [go-marathon](https://github.com/gambol99/go-marathon).

```
func listMarathonApps(marathonClient marathon.Marathon) error {
	applications, err := marathonClient.Applications(nil)
	if err != nil {
		return err
	}

	log.Printf("Found %d application(s) running", len(applications.Apps))
	for _, application := range applications.Apps {
		log.Printf("Application: %+v\n", application)
	}
	return nil
}
```
