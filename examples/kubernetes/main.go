package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/dcos/client-go/dcos"
)

var client *dcos.APIClient

func createServiceAccount(uid, pubKey, privKey string) {
	resp, err := client.IAM.CreateUser(context.TODO(), uid, dcos.IamUserCreate{
		Description: "Service account for " + uid,
		PublicKey:   pubKey,
	})
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Printf("user '%s' already exists\n", uid)
	}

	serviceAccountLoginSecret, err := json.Marshal(dcos.ServiceAccountOptions{
		UID:           uid,
		PrivateKey:    privKey,
		LoginEndpoint: "https://leader.mesos/acs/api/v1/auth/login",
		Scheme:        "RS256",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Secrets.CreateSecret(context.TODO(), "default", uid+"/sa", dcos.SecretsV1Secret{
		Value: string(serviceAccountLoginSecret),
	})
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Printf("login secret for '%s' service account already exists\n", uid)
	}
}

func addRepo(name, uri string) {
	_, _, err := client.Cosmos.PackageRepositoryAdd(context.TODO(), &dcos.PackageRepositoryAddOpts{
		CosmosPackageAddRepoV1Request: optional.NewInterface(dcos.CosmosPackageAddRepoV1Request{
			Name: name,
			Uri:  uri,
		}),
	})
	if err != nil {
		switch err := err.(type) {
		case dcos.GenericOpenAPIError:
			if err.Model() != nil {
				if e, ok := err.Model().(dcos.CosmosError); ok && e.Type == "RepositoryAlreadyPresent" {
					log.Println("Repo already present")
					return
				}
			}
		}
		log.Fatal(err)
	}
}

func addPerm(rid, uid, action string) {
	rid = strings.Replace(rid, "/", "%252F", -1)
	resp, err := client.IAM.CreateResourceACL(context.TODO(), rid, dcos.IamaclCreate{})
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Printf("permission '%s:%s' for user '%s' already exists", rid, action, uid)
	}

	resp, err = client.IAM.PermitResourceUserAction(context.TODO(), rid, uid, action)
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Printf("permission '%s:%s' for user '%s' already exists", rid, action, uid)
	}
}

func installMKE() {
	createServiceAccount("kubernetes", mkePubKeyPEM, mkePrivKeyPEM)

	addPerm("dcos:mesos:master:reservation:role:kubernetes-role", "kubernetes", "create")
	addPerm("dcos:mesos:master:framework:role:kubernetes-role", "kubernetes", "create")
	addPerm("dcos:mesos:master:task:user:nobody", "kubernetes", "create")

	_, resp, err := client.Cosmos.PackageInstall(context.TODO(), dcos.CosmosPackageInstallV1Request{
		PackageName: "kubernetes",
		Options: map[string]map[string]interface{}{
			"service": {
				"service_account":        "kubernetes",
				"service_account_secret": "kubernetes/sa",
			},
		},
	})
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Println("package 'kubernetes' already installed")
	}
}

func createKubernetesCluster() {
	createServiceAccount("kubernetes-cluster1", k8sPubKeyPEM, k8sPrivKeyPEM)

	// Master node permissions
	addPerm("dcos:mesos:master:framework:role:kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:task:user:root", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:agent:task:user:root", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:reservation:role:kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:reservation:principal:kubernetes-cluster1", "kubernetes-cluster1", "delete")
	addPerm("dcos:mesos:master:volume:role:kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:volume:principal:kubernetes-cluster1", "kubernetes-cluster1", "delete")

	// Secrets permissions
	addPerm("dcos:secrets:default:/kubernetes-cluster1/*", "kubernetes-cluster1", "full")
	addPerm("dcos:secrets:list:default:/kubernetes-cluster1", "kubernetes-cluster1", "read")

	// Adminrouter permissions
	addPerm("dcos:adminrouter:ops:ca:rw", "kubernetes-cluster1", "full")
	addPerm("dcos:adminrouter:ops:ca:ro", "kubernetes-cluster1", "full")

	// Public agent permissions
	addPerm("dcos:mesos:master:framework:role:slave_public/kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:framework:role:slave_public/kubernetes-cluster1-role", "kubernetes-cluster1", "read")
	addPerm("dcos:mesos:master:reservation:role:slave_public/kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:volume:role:slave_public/kubernetes-cluster1-role", "kubernetes-cluster1", "create")
	addPerm("dcos:mesos:master:framework:role:slave_public", "kubernetes-cluster1", "read")
	addPerm("dcos:mesos:agent:framework:role:slave_public", "kubernetes-cluster1", "read")

	_, resp, err := client.Cosmos.PackageInstall(context.TODO(), dcos.CosmosPackageInstallV1Request{
		PackageName: "kubernetes-cluster",
		Options: map[string]map[string]interface{}{
			"service": {
				"name":                   "kubernetes-cluster1",
				"service_account":        "kubernetes-cluster1",
				"service_account_secret": "kubernetes-cluster1/sa",
			},
		},
	})

	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Println("package 'kubernetes-cluster' already installed")
	}
}

func configureEdgeLB() {
	// Add EdgeLB repositories
	addRepo("edgelb", "https://downloads.mesosphere.com/edgelb/v1.3.0/assets/stub-universe-edgelb.json")
	addRepo("edgelb-pool", "https://downloads.mesosphere.com/edgelb-pool/v1.3.0/assets/stub-universe-edgelb-pool.json")

	// Create dcos-edglelb service account and login secret
	createServiceAccount("dcos-edgelb", edgeLBPubKeyPEM, edgeLBPrivKeyPEM)

	// Add dcos-edgelb to the superusers group
	resp, err := client.IAM.CreateGroupUser(context.TODO(), "superusers", "dcos-edgelb")
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Println("'dcos-edgelb' is already in the superusers group")
	}

	// Install EdgeLB
	_, resp, err = client.Cosmos.PackageInstall(context.TODO(), dcos.CosmosPackageInstallV1Request{
		PackageName: "edgelb",
		Options: map[string]map[string]interface{}{
			"service": {
				"secretName":    "dcos-edgelb/sa",
				"principal":     "dcos-edgelb",
				"mesosProtocol": "https",
			},
		},
	})

	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Println("package 'edgelb' already installed")
	}

	// Wait for EdgeLB
	for i := 0; ; i++ {
		_, _, err := client.Edgelb.Ping(context.TODO())
		if err == nil {
			break
		}
		if i == 10 {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}

	// Create the EdgeLB pool
	edgelbPool := dcos.EdgelbV2Pool{
		ApiVersion:      dcos.EdgelbApiVersion("V2"),
		Name:            "edgelb-kubernetes-cluster-proxy-basic",
		Count:           1,
		AutoCertificate: true,
		Haproxy: dcos.EdgelbV2Haproxy{
			Frontends: []dcos.EdgelbV2Frontend{
				{
					BindPort:     6443,
					Protocol:     "HTTPS",
					Certificates: []string{"$AUTOCERT"},
					LinkBackend: dcos.EdgelbV2FrontendLinkBackend{
						DefaultBackend: "kubernetes-cluster1",
					},
				},
			},
			Backends: []dcos.EdgelbV2Backend{
				{
					Name:     "kubernetes-cluster1",
					Protocol: "HTTPS",
					Services: []dcos.EdgelbV2Service{
						{
							Mesos: dcos.EdgelbV2ServiceMesos{
								FrameworkName:   "kubernetes-cluster1",
								TaskNamePattern: "kube-control-plane",
							},
							Endpoint: dcos.EdgelbV2Endpoint{
								PortName: "apiserver",
							},
						},
					},
				},
			},
			Stats: dcos.EdgelbV2Stats{
				BindPort: 6090,
			},
		},
	}

	_, resp, err = client.Edgelb.V2CreatePool(context.TODO(), edgelbPool)
	if err != nil {
		if resp == nil || resp.StatusCode != 409 {
			log.Fatal(err)
		}
		log.Println("edgelb pool already created")
	}
}

func main() {
	var err error
	client, err = dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	installMKE()
	createKubernetesCluster()
	configureEdgeLB()
}
