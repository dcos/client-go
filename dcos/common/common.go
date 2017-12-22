package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	upath "path"

	"github.com/mesosphere/dcos-go/dcos/config"
)

type Service struct {
	Config *config.Config
	Client *http.Client
	Path   string
}

type Job struct {
}

func (s *Service) GetServiceURL(path string) *url.URL {
	u, _ := s.Config.GetDCOSUrl()
	u.Path = upath.Join(u.Path, s.Path, path)
	return u
}

func (s *Service) GetJSON(path string, object interface{}) error {
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

	if resp.StatusCode >= 400 {
		return fmt.Errorf("Post respond with status: %d", resp.StatusCode)
	}

	recBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(recBody, receiveObject)

	if err != nil {
		return err
	}
	return nil
}
