package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/huysamen/domains-go/domains"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"reflect"
	"time"
)

const BaseUrl = "https://api-v3.domains.co.za/api"

type Api struct {
	key  string
	http *http.Client
}

func Default() (*Api, error) {
	apiKey := os.Getenv("DOMAINS_API_KEY")

	if apiKey == "" {
		return nil, errors.New("no api key present")
	}

	return &Api{
		key: apiKey,
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
	}, nil
}

func (a *Api) Domains() *domains.Client {
	return domains.Create(a)
}

func (a *Api) DoPost(url string, payload interface{}) ([]byte, error) {
	pt := reflect.ValueOf(payload).Elem()
	f := pt.FieldByName("Key")

	if f.IsValid() {
		f.SetString(a.key)
	}

	req, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	rsp, err := a.http.Post(url, "application/json", bytes.NewBuffer(req))

	if err != nil {
		return nil, err
	}

	defer func() { _ = rsp.Body.Close() }()

	return ioutil.ReadAll(rsp.Body)
}
