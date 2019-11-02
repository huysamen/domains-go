package domains

import (
	"fmt"
	"github.com/huysamen/domains-go/api"
)

const (
	checkDomain         = "/domain/domain/check"
	registerDomain      = "/domain/domain/create"
	transferDomain      = "/domain/domain/transfer"
	transferDomainCheck = "/domain/domain/transferCheck"
)

var (
	checkDomainUrl         = fmt.Sprintf("%s%s", api.BaseUrl, checkDomain)
	registerDomainUrl      = fmt.Sprintf("%s%s", api.BaseUrl, registerDomain)
	transferDomainUrl      = fmt.Sprintf("%s%s", api.BaseUrl, transferDomain)
	transferDomainCheckUrl = fmt.Sprintf("%s%s", api.BaseUrl, transferDomainCheck)
)

type Client struct {
	api *api.Api
}

func Create(api *api.Api) *Client {
	return &Client{api: api}
}
