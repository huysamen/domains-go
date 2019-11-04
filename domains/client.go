package domains

import (
	"github.com/huysamen/domains-go/types"
)

const (
	checkDomainPath         = "/domain/domain/check"
	registerDomainPath      = "/domain/domain/create"
	transferDomainPath      = "/domain/domain/transfer"
	transferDomainCheckPath = "/domain/domain/transferCheck"
)

type Client struct {
	post types.PostRequest
}

func Create(post types.PostRequest) *Client {
	return &Client{post: post}
}
