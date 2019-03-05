package dcos

import (
	"fmt"
	"net/http"
	"net/url"

	secretsclient "github.com/dcos/client-go/dcos/secrets/client"

	iamclient "github.com/dcos/client-go/dcos/iam/client"
	iamlogin "github.com/dcos/client-go/dcos/iam/client/login"
	iammodels "github.com/dcos/client-go/dcos/iam/models"

	runtimeClient "github.com/go-openapi/runtime/client"
)

// Client is a client for DC/OS.
type Client struct {
	HTTPClient *http.Client
	Config     *Config

	IAM     *iamclient.IdentityAndAccessManagement
	Secrets *secretsclient.DCOSSecrets
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

	return &Client{
		HTTPClient: httpClient,
		Config:     config,
		IAM:        iamClient,
		Secrets:    secretsClient,
	}, nil
}

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
