package client

import (
	"net/http"
	"testing"

	"github.com/mesosphere/dcos-go/dcos/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

// Lets do an samle request and see if we find the token header
func TestNewClient(t *testing.T) {
	testurl := "http://foobar.com/baz"
	testtoken := "Testtoken"
	conf := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
		},
	}
	gock.New(testurl).
		MatchHeader("Authorization", "^token="+testtoken+"$").
		Reply(200).
		BodyString("success")

	// return 401 if auth is not present
	gock.New(testurl).
		Reply(401).
		BodyString("denied")

	c := NewClient(&conf)
	resp, err := c.Get(testurl)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	//just test our mock
	c = http.DefaultClient
	resp, err = c.Get(testurl)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestNewClientInvalidConf(t *testing.T) {
	testurl := "http://foobar.com/baz"

	gock.New(testurl).
		Reply(201).
		BodyString("success")

	conf := config.Config{}

	assert.False(t, conf.Valid())

	c := NewClient(&conf)
	resp, err := c.Get(testurl)
	assert.Error(t, err)
	assert.Nil(t, resp)

}
