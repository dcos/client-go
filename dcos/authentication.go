package dcos

import (
	"fmt"
	"net/http"
	"time"
)

type Authentication interface {
	AddHeaders(req *http.Request)
	Login() error
	Expires(in time.Duration) bool
}

type TestAuthentication struct {
	Token string
}

func (a *TestAuthentication) Login() error {
	return nil
}

func (a *TestAuthentication) Expires(in time.Duration) bool {
	if in > time.Minute {
		return false
	}
	return true
}

func (a *TestAuthentication) AddHeaders(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("token=%s", a.Token))
}
