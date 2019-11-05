package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/huysamen/domains-go/domains"
)

const baseUrl = "https://api-v3.domains.co.za/api"

type Api struct {
	apiKey string
	http   *http.Client

	Domains *domains.Client
}

func Default() (*Api, error) {
	apiKey := os.Getenv("DOMAINS_API_KEY")

	if apiKey == "" {
		return nil, errors.New("no api key present")
	}

	api := &Api{
		apiKey: apiKey,
		http: &http.Client{
			Timeout: time.Second * 60,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					MinVersion: tls.VersionTLS12,
				},
			},
		},
	}

	api.createServices()

	return api, nil
}

func (a *Api) createServices() {
	a.Domains = domains.Create(a.post)
}

func (a *Api) post(path string, payload interface{}) ([]byte, error) {
	p, err := convertToMap(payload)
	if err != nil {
		return nil, err
	}

	p["key"] = a.apiKey

	req, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	rsp, err := a.http.Post(baseUrl+path, "application/json", bytes.NewBuffer(req))

	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}

// not ideal, but we don't want people to think they need to set the api key
func convertToMap(payload interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}

	buff, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buff, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
