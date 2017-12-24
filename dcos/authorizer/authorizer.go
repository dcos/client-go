// This package provides a generic authorizer component to applications that
// need to perform authorization against the DC/OS Enterprise IAM. It does so
// by querying the DC/OS Enterprise IAM for the permissions assigned to the
// user whose action needs authorization.
//
// The authorizer needs to be initialized with DC/OS service account
// credentials to use when querying the IAM for a client's permissions. The
// authorizer periodically refreshes it's DC/OS Authentication Token using the
// service account credentials.
//
// Once a given principal's permissions are retrieved from the IAM, they are
// cached for a configurable amount of time in accordance with DC/OS Enterprise
// Authorizer best practices.
package authorizer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

var actionNames = map[string]struct{}{
	"create": struct{}{},
	"read":   struct{}{},
	"update": struct{}{},
	"delete": struct{}{},
	"full":   struct{}{},
}

type simpleError string

func (e simpleError) Error() string {
	return string(e)
}

const UnknownActionErr = simpleError("authorizer: action must be one of 'create', 'read', 'update', 'delete', 'full'.")

type Authorizer interface {
	Authorized(principal, resource, action string) (bool, error)
}

type authorizer struct {
	// An instance of "dcos/dcos-go/dcos/http/transport".RoundTripper that
	// is configured with the correct IAM login URL and service account
	// credentials.
	rt http.RoundTripper
	// The IAM base URL (ie., https://<master>/acs/api/v1)
	iamURL string
	// The User-Agent to perform login and authorization requests as.
	userAgent string
	// The *http.Client used to query the /permissions endpoint.
	client *http.Client
	// The permissions cache
	permscache *cache.Cache
}

// NewAuthorizer returns a authorizer that can be used to retrueve a user's
// permissions from the IAM.
//
// It expects a `"dcos/dcos-go/dcos/http/transport".RoundTripper` that is
// configured with the correct IAM login URL and service account credentials.
//
// Permissions are cached for `cacheTTL` duration and stale entries are evicted
// every second.
func NewAuthorizer(rt http.RoundTripper, iamURL, userAgent string, cacheTTL, timeout time.Duration) Authorizer {
	client := &http.Client{
		Transport: rt,
		Timeout:   timeout,
	}
	// purge old items every second
	const purgeInterval = time.Second
	permscache := cache.New(cacheTTL, purgeInterval)
	return &authorizer{rt, iamURL, userAgent, client, permscache}
}

// Authorized determines whether or not the `principal` (ie., uid) is permitted
// to perform the `action` on the specified `resource`. For example,
//
// Authorized("alice", "dcos:service:marathon:marathon:admin:config", "read")
func (a *authorizer) Authorized(principal, resource, action string) (bool, error) {
	var perms *permissionsResponse
	if x, ok := a.permscache.Get(principal); ok {
		perms = x.(*permissionsResponse)
	} else {
		permissionsURL := fmt.Sprintf("%s/users/%s/permissions", a.iamURL, principal)
		req, err := http.NewRequest("GET", permissionsURL, nil)
		if err != nil {
			return false, err
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		resp, err := a.client.Do(req)
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()
		perms = new(permissionsResponse)
		if err := json.NewDecoder(resp.Body).Decode(perms); err != nil {
			return false, err
		}
		a.permscache.Set(principal, perms, cache.DefaultExpiration)
	}
	for _, perm := range perms.Direct {
		if perm.Rid != resource {
			continue
		}
		for _, a := range perm.Actions {
			if a.Name == action {
				return true, nil
			}
		}
	}
	for _, perm := range perms.Groups {
		if perm.Rid != resource {
			continue
		}
		for _, a := range perm.Actions {
			if a.Name == action {
				return true, nil
			}
		}
	}
	return false, nil
}

type permissionsResponse struct {
	Direct []struct {
		Rid         string `json:"rid"`
		Description string `json:"description"`
		AclURL      string `json:"aclurl"`
		Actions     []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"actions"`
	}
	Groups []struct {
		Rid           string `json:"rid"`
		Gid           string `json:"gid"`
		Description   string `json:"description"`
		AclURL        string `json:"aclurl"`
		MembershipURL string `json:"membershipurl"`
		Actions       []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"actions"`
	}
}
