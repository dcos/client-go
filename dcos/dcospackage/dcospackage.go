package dcospackage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mesosphere/dcos-api-go/dcos/common"
	"github.com/mesosphere/dcos-api-go/dcos/config"
)

// SearchContentType media type for Search
const SearchContentType = "application/vnd.dcos.package.search-request+json;charset=utf-8;version=v1"

// SearchResultContentType media type for SearchResult
const SearchResultContentType = "application/vnd.dcos.package.search-response+json;charset=utf-8;version=v1"

// PackageService represents access to package API
type PackageService struct {
	*common.Service
}

func (p *PackageService) post(url, contentType, accept string, input, output interface{}) (err error) {
	in, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(in))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", contentType)

	resp, err := p.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("Unauthorized. You're not allowed to request this enpoint")
	}

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Request error got: %d - body: %s", resp.StatusCode, b)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, output)
	if err != nil {
		return err
	}

	return nil
}

func (p *PackageService) get(url, output interface{}) (err error) {
	return nil
}

// NewPackageService creates a new PackageService from config and client with defaults
func NewPackageService(config *config.Config, client *http.Client) *PackageService {
	p := PackageService{
		&common.Service{
			Config: config,
			Client: client,
			Path:   "/package",
		},
	}

	return &p
}

// SearchPackage respresents a DC/OS package
type SearchPackage struct {
	Name           string `json:"name"`
	CurrentVersion string `json:"currentVersion"`
	Description    string `json:"description"`
	Framework      bool   `json:"framework"`
	Selected       bool   `json:"selected"`
	Images         struct {
		IconSmall   string `json:"icon-small"`
		IconMedium  string `json:"icon-medium"`
		IconLarge   string `json:"icon-large"`
		Screenshots []string
	} `json:"images"`
	Tags     []string       `json:"tags"`
	Versions map[string]int `json:"versions"` //FIXME: API definition states map[string]string
}

// SearchResult containes a slices of found Packages
type SearchResult struct {
	Packages []SearchPackage
}

// Search returns SearchResult for packages matching query
func (p *PackageService) Search(query string) *SearchResult {

	q := map[string]string{"Query": query}
	r := SearchResult{}

	err := p.post(p.GetServiceURL("/search").String(), SearchContentType, SearchResultContentType, &q, &r)
	if err != nil {
		fmt.Printf("%v", err)
	}

	return &r
}

// // List returns a list of installed packages
// func (p *PackageService) List() *SearchResult {
// 	err := p.Client.Post(p.GetServiceURL("/search").String(), SearchContentType, SearchResultContentType, &q, &r)
// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}
//
// 	return &r
// }
