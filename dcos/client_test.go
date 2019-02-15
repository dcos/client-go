package dcos

import (
	"net/http"
	"testing"
)

func TestClientNew(t *testing.T) {
	baseClient := &http.Client{}
	c, err := NewClient(baseClient)
	if err != nil {
		t.Fatalf("NewClient returned unexpected error: %s", err)
	}

	if c.HttpClient != baseClient {
		t.Errorf("client.HttpClient wrong. got=%+v, want=%+v",
			c.HttpClient, baseClient)
	}

	c, err = NewClient(nil)
	if err != nil {
		t.Fatalf("NewClient returned unexpected error: %s", err)
	}

	if c.HttpClient != http.DefaultClient {
		t.Errorf("client.HttpClient not http.Defaultclient. got=%+v",
			c.HttpClient)
	}
}
