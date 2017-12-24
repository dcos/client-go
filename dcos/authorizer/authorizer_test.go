package authorizer

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/dcos/dcos-go/dcos/http/transport"
	"github.com/google/uuid"
	"github.com/renstrom/dedent"
)

var (
	testIAMURL   = flag.String("iam-url", "", "The IAM base URL (eg., http://172.17.0.2/acs/api/v1)")
	testUsername = flag.String("username", "admin", "The DC/OS username to use when setting up a test service account")
	testPassword = flag.String("password", "admin", "The DC/OS password to use when setting up a test service account")
	testCABundle = flag.String("ca-bundle", "", "A CA certificate bundle to use when using TLS")
)

func init() {
	flag.Parse()
}

var testPrivateKey = dedent.Dedent(`
        -----BEGIN PRIVATE KEY-----
        MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCyRQCnWnTYAqIK
        pUp8Y8sP5kq3zjYHm1DkdUXYuukBguRjkitqms46+V6MYEGhI+PfzCXsSODFfs2h
        oOrtKEDkhLaiyIt7/JiNVXKqAzCH0Tmz3CeuM5dZT+ZZM3R0Z3xOdTCY9z2FpUA1
        j79URO9UziG1VI9sOXwI87QyjPB/a6HJmsnvkCxy4Co/gRM56hTR4S0vuQhfOX0U
        XhGEwb3hzdNqv4yrbmLNrz6X/3CLnUtl5kj1Ea0vJ98OdmWKAH/wH6J54v4w64Z7
        6zx4nO8ec97GkNbaeol6C6KXZ2Hh7YzMZYgkDKQkbLchAl2B51A/57FJghp3e6mg
        MZ0lEHONAgMBAAECggEADdItt8vajTi/CWZBPR2bi7MGDfQN2k8fWWcCMEhlcjmM
        rW+SsCQqYwYcX3PDAtQ+gYZZVlVcH9Ox65sHO7XIOi9T3ZEAx28RALwkNKwkohMF
        jz09UlMro7//EuFbYP748zhAuUtJ2Me0D7MCWW0yxPdqQ523tONMC4Ghcd3dc3Nq
        qK/gE6TLXYENdKtjz9+RLi5JyXXSA3yEsDTcCz8Mq7/AV59B8kGdqU8CSkXEaqNg
        fHqVCtFRkkKv1poVrrKYHR91VtpLeJ0hom9DqYPW/h71bXaMXEn3p3X6rNojZ+3W
        9d48dHJrUf4b8CDK4d8h4nIzpWOh15W3Jc6aLN0q3QKBgQDRFQ5TqgYU7y2vheCI
        pHeoM1HfkMFYxg3vJSeplanz1GTW3/ZCPruWb8Sa/TMkvw9jjNTcCVgGenwb+Q7i
        9toZq5SdD6cJ2AfUy3jW1t0vfF+9gJAVXyAiYmt6SfK7zroSBWGYOFwWfzbws/mU
        U844xAt6bC9goHQ5643MK6jZJwKBgQDaReEElGWDRTNDKnG15ou22O5UIhX+O2EN
        nuxuE1BswHPqSM2/CYPLXOztM4IllnPZ2Tc8yL2cpnMgLaSPlbSkdKF4EqKezNew
        tBRgQHunjTjh63bqJlbUDQNbMtzHbPj9ViVTSx91sTrq8TReB9shyIo8OdBSMIhu
        q++lEQR2KwKBgDptUeWbW7u54uBr3wUQQMfmH3kEMKOoJPixP1YqlqWmss9BIsX8
        VduCPES7gPFB2wEMt+cNTWXyEICYpspSBTyrMAp790fO8NFMUuZmgMTmUrLYHGt8
        e7RgIHy58XT+Hg0vJ8ir5z/6KNY21QXYLYa7ciGx3cN400HkS2nEwZh9AoGBANQD
        0GAMbNYnX3W8Qpt3hvpimI0sVsSNlOYQGoOJSzN3dKB7ey425cehSZnV0HMGQx21
        +guWyWbh3y3AZ/n4ZnDfwP/Kyh8JxfCj2q7rs7C6Yig/+U/TUb/DrCDntMAZP3mE
        NdlEPRiqwBn5lO5JHzcfffamCY3sCnnFlI93iiflAoGAKYGlDSg8GcBPl935O6gy
        3eRGc7DO1LO16gY+VC+TgGN/OSVX1EtE5Z6uIWSgcday7KNzarQpyOf9FQcoWVf1
        z6dZ80cCfBZG61PE7zHmTCnq97uZi+cqvM00kbQmKjB77gNtPhIKbdRKsQUpx1GJ
        3cUeg0nUHT5cg6LMl/we9ys=
        -----END PRIVATE KEY-----`)
var testPublicKey = dedent.Dedent(`
        -----BEGIN PUBLIC KEY-----
        MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAskUAp1p02AKiCqVKfGPL
        D+ZKt842B5tQ5HVF2LrpAYLkY5IraprOOvlejGBBoSPj38wl7EjgxX7NoaDq7ShA
        5IS2osiLe/yYjVVyqgMwh9E5s9wnrjOXWU/mWTN0dGd8TnUwmPc9haVANY+/VETv
        VM4htVSPbDl8CPO0Mozwf2uhyZrJ75AscuAqP4ETOeoU0eEtL7kIXzl9FF4RhMG9
        4c3Tar+Mq25iza8+l/9wi51LZeZI9RGtLyffDnZligB/8B+ieeL+MOuGe+s8eJzv
        HnPexpDW2nqJeguil2dh4e2MzGWIJAykJGy3IQJdgedQP+exSYIad3upoDGdJRBz
        jQIDAQAB
        -----END PUBLIC KEY-----`)

func TestAuthorized(t *testing.T) {
	if *testIAMURL == "" {
		t.Skip("-iam-url not provided")
	}
	// Disable TLS peer verification unless a CA bundle is provided.
	skipVerifyTLS := &tls.Config{InsecureSkipVerify: true}
	var tr http.RoundTripper = &http.Transport{TLSClientConfig: skipVerifyTLS}
	if *testCABundle != "" {
		var err error
		tr, err = transport.NewTransport(transport.OptionCaCertificatePath(*testCABundle))
		if err != nil {
			t.Fatal(err)
		}
	}
	// Ignore TLS verification for tests.
	client := &http.Client{
		Transport: tr,
	}
	testid := uuid.New().String()
	// Perform login so we can modify IAM state.
	authToken, err := login(client)
	if err != nil {
		t.Fatal(err)
	}
	// Create service account to use for authorizer
	authzUid, err := createServiceAccount(testid, authToken, client)
	if err != nil {
		t.Fatal(err)
	}
	// Add service account to superusers group
	if err := addToSuperusers(authzUid, authToken, client); err != nil {
		t.Fatal(err)
	}
	// Create user account to use for testing
	uid, err := createUserAccount(testid, authToken, client)
	if err != nil {
		t.Fatal(err)
	}
	// Create an ACL for use in testing
	rid, err := createACL(testid, authToken, client)
	if err != nil {
		t.Fatal(err)
	}
	// Setup authorizer roundtripper
	loginURL := fmt.Sprintf("%s/auth/login", *testIAMURL)
	rtcreds := transport.OptionCredentials(authzUid, testPrivateKey, loginURL)
	rt, err := transport.NewRoundTripper(tr, rtcreds)
	if err != nil {
		t.Fatal(err)
	}
	// Setup authorizer
	const cacheTTL = time.Second
	const timeout = 5 * time.Second
	authz := NewAuthorizer(rt, *testIAMURL, "testauthz", cacheTTL, timeout)
	// Check that the user does not have permission.
	ok, err := authz.Authorized(uid, rid, "read")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("expected to not be authorized")
	}
	// Grant 'read' permission to uid for rid.
	if err := grantPermissionToUser(uid, rid, "read", authToken, client); err != nil {
		t.Fatal(err)
	}
	// Check that the user still has no permission (due to permissions cache).
	ok, err = authz.Authorized(uid, rid, "read")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("expected to not be authorized")
	}
	// Wait for the cached permissions to expire
	time.Sleep(cacheTTL)
	// Check that the user now has permission.
	ok, err = authz.Authorized(uid, rid, "read")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected to be authorized")
	}

}

func login(client *http.Client) (string, error) {
	loginURL := fmt.Sprintf("%s/auth/login", *testIAMURL)
	loginParams, err := json.Marshal(map[string]string{"uid": *testUsername, "password": *testPassword})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", loginURL, bytes.NewReader(loginParams))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	tokenResponse := new(struct{ Token string })
	if err := json.NewDecoder(resp.Body).Decode(tokenResponse); err != nil {
		return "", err
	}
	return tokenResponse.Token, nil
}

func createServiceAccount(testid, authToken string, client *http.Client) (string, error) {
	userParams, err := json.Marshal(map[string]string{
		"description": "test user",
		"public_key":  testPublicKey,
	})
	if err != nil {
		return "", err
	}
	uid := fmt.Sprintf("testuser-%s", testid)
	userURL := fmt.Sprintf("%s/users/%s", *testIAMURL, uid)
	req, err := http.NewRequest("PUT", userURL, bytes.NewReader(userParams))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "token="+authToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return "", errors.New("failed to create service account")
	}
	return uid, nil
}

func addToSuperusers(uid, authToken string, client *http.Client) error {
	superuserURL := fmt.Sprintf("%s/groups/superusers/users/%s", *testIAMURL, uid)
	req, err := http.NewRequest("PUT", superuserURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token="+authToken)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return errors.New("failed to assign superuser privileges")
	}
	return nil
}

func createUserAccount(testid, authToken string, client *http.Client) (string, error) {
	userParams, err := json.Marshal(map[string]string{
		"description": "test user",
		"password":    "some password",
	})
	if err != nil {
		return "", err
	}
	uid := fmt.Sprintf("regular-testuser-%s", testid)
	userURL2 := fmt.Sprintf("%s/users/%s", *testIAMURL, uid)
	req, err := http.NewRequest("PUT", userURL2, bytes.NewReader(userParams))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "token="+authToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return "", errors.New("failed to create user account")
	}
	return uid, nil
}

func createACL(testid, authToken string, client *http.Client) (string, error) {
	reg := regexp.MustCompile("[0-9]+")
	alphaTestid := reg.ReplaceAllString(testid, "-")
	rid := fmt.Sprintf("dcos:acs:testrid-%s", alphaTestid)
	aclURL := fmt.Sprintf("%s/acls/%s", *testIAMURL, rid)
	aclParams, err := json.Marshal(map[string]string{
		"description": "test rid",
	})
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("PUT", aclURL, bytes.NewReader(aclParams))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "token="+authToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return "", errors.New("failed to create ACL")
	}
	return rid, nil
}

func grantPermissionToUser(uid, rid, action, authToken string, client *http.Client) error {
	allowURL := fmt.Sprintf("%s/acls/%s/users/%s/read", *testIAMURL, rid, uid)
	req, err := http.NewRequest("PUT", allowURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token="+authToken)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 204 {
		return errors.New("failed to grant permission")
	}
	return nil
}
