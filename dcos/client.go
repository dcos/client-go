package dcos

import (
	"fmt"
	"net/http"
	"net/url"

	secretsclient "github.com/dcos/client-go/dcos/secrets/client"

	cosmosclient "github.com/dcos/client-go/dcos/cosmos/client"

	iamclient "github.com/dcos/client-go/dcos/iam/client"
	iamlogin "github.com/dcos/client-go/dcos/iam/client/login"
	iammodels "github.com/dcos/client-go/dcos/iam/models"

	"github.com/go-openapi/runtime"
	runtimeClient "github.com/go-openapi/runtime/client"
)

// Client is a client for DC/OS.
type Client struct {
	HTTPClient *http.Client
	Config     *Config

	IAM     *iamclient.IdentityAndAccessManagement
	Secrets *secretsclient.DCOSSecrets
	Cosmos  *cosmosclient.Cosmos
}

// NewClient returns a Client which will detect its Config through the well
// known process and use a default http.Client with DC/OS authentication.
func NewClient() (*Client, error) {
	config, err := NewConfigManager(nil).Current()
	if err != nil {
		return nil, err
	}

	httpClient, err := NewHTTPClient(config)
	if err != nil {
		return nil, err
	}

	return NewClientWithOptions(httpClient, config)
}

// NewClientWithOptions creates a Client from a given *http.Client and *Config
func NewClientWithOptions(httpClient *http.Client, config *Config) (*Client, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("httpClient cannot be nil")
	}

	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	iamClient, err := newIamClient(config.URL(), httpClient)
	if err != nil {
		return nil, fmt.Errorf("could not create IAM client: %s", err)
	}

	secretsClient, err := newSecretsClient(config.URL(), httpClient)
	if err != nil {
		return nil, fmt.Errorf("could not create Secrets client: %s", err)
	}

	cosmosClient, err := newCosmosClient(config.URL(), httpClient)
	if err != nil {
		return nil, fmt.Errorf("could not create Cosmos client: %s", err)
	}

	return &Client{
		HTTPClient: httpClient,
		Config:     config,
		IAM:        iamClient,
		Secrets:    secretsClient,
		Cosmos:     cosmosClient,
	}, nil
}

// Login uses the IAM client to create a new authentication token for the given
// username and password
func (c *Client) Login(username, password string) (string, error) {
	loginObject := &iammodels.LoginObject{UID: username, Password: password}
	params := iamlogin.NewPostAuthLoginParams().WithLoginObject(loginObject)

	result, err := c.IAM.Login.PostAuthLogin(params)
	if err != nil {
		return "", err
	}

	return *result.Payload.Token, nil
}

func newIamClient(clusterURL string, client *http.Client) (*iamclient.IdentityAndAccessManagement, error) {
	dcosURL, err := url.Parse(clusterURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid DC/OS cluster URL '%s': %v", clusterURL, err)
	}

	iamRuntime := runtimeClient.NewWithClient(
		dcosURL.Host,
		iamclient.DefaultBasePath,
		[]string{dcosURL.Scheme},
		client,
	)

	return iamclient.New(iamRuntime, nil), nil
}

func newSecretsClient(clusterURL string, client *http.Client) (*secretsclient.DCOSSecrets, error) {
	dcosURL, err := url.Parse(clusterURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid DC/OS cluster URL '%s': %v", clusterURL, err)
	}

	secretsRuntime := runtimeClient.NewWithClient(
		dcosURL.Host,
		secretsclient.DefaultBasePath,
		[]string{dcosURL.Scheme},
		client,
	)

	return secretsclient.New(secretsRuntime, nil), nil
}

func newCosmosClient(clusterURL string, client *http.Client) (*cosmosclient.Cosmos, error) {
	dcosURL, err := url.Parse(clusterURL)
	if err != nil {
		return nil, fmt.Errorf("Invalid DC/OS cluster URL '%s': %v", clusterURL, err)
	}

	cosmosRuntime := runtimeClient.NewWithClient(
		dcosURL.Host,
		cosmosclient.DefaultBasePath,
		[]string{dcosURL.Scheme},
		client,
	)

	cosmosRuntime.Consumers["application/vnd.dcos.package.install-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.package.install-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.package.uninstall-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.package.uninstall-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.package.error+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.service.update-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.service.update-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.package.repository.add-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.package.repository.add-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.package.repository.delete-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.package.repository.delete-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]
	cosmosRuntime.Consumers["application/vnd.dcos.service.describe-response+json"] = cosmosRuntime.Consumers[runtime.JSONMime]
	cosmosRuntime.Producers["application/vnd.dcos.service.describe-request+json;charset=utf-8;version=v1"] = cosmosRuntime.Producers[runtime.JSONMime]

	return cosmosclient.New(cosmosRuntime, nil), nil
}
