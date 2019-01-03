package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	upath "path"

	"github.com/mesosphere/dcos-api-go/dcos/config"
)

type Service struct {
	Config *config.Config
	Client *http.Client
	Path   string
}

type RequestError struct {
	err    string
	status int
	body   []byte
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("Code: %d - %s", e.status, e.err)
}

func (e *RequestError) Body() []byte {
	return e.body
}

func (e *RequestError) StatusCode() int {
	return e.status
}

type Job struct {
}

func (s *Service) GetServiceURL(path string) *url.URL {
	u, _ := s.Config.GetDCOSUrl()
	u.Path = upath.Join(u.Path, s.Path, path)
	return u
}

func (s *Service) GetJSON(path string, object interface{}) error {
	resp, err := s.Client.Get(s.GetServiceURL(path).String())
	if err != nil {
		return err
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return &RequestError{"HTTP request not successful", resp.StatusCode, recBody}
	}

	err = json.Unmarshal(recBody, object)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(path string) error {
	req := &http.Request{}
	req.URL = s.GetServiceURL(path)
	req.Method = "DELETE"
	resp, err := s.Client.Do(req)

	if err != nil {
		return err
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return &RequestError{"Not authorized", resp.StatusCode, recBody}
	}

	if resp.StatusCode == 404 {
		return &RequestError{"Job not found", resp.StatusCode, recBody}
	}

	if resp.StatusCode >= 400 {
		return &RequestError{"Unknown error", resp.StatusCode, recBody}
	}

	return nil
}

func (s *Service) PostJSON(path string, sentObject, receiveObject interface{}) error {
	body, err := json.Marshal(sentObject)
	if err != nil {
		return err
	}

	resp, err := s.Client.Post(s.GetServiceURL(path).String(), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return &RequestError{"Post respond with 4xx", resp.StatusCode, recBody}
	}

	err = json.Unmarshal(recBody, receiveObject)

	if err != nil {
		return err
	}
	return nil
}

func (s *Service) PutJSON(path string, sentObject, receiveObject interface{}) error {
	body, err := json.Marshal(sentObject)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", s.GetServiceURL(path).String(), bytes.NewBuffer(body))

	req.Header.Set("content-type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return &RequestError{"Post respond with 4xx", resp.StatusCode, recBody}
	}

	err = json.Unmarshal(recBody, receiveObject)

	if err != nil {
		return err
	}
	return nil
}
