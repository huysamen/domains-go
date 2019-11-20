package reseller

import (
	"github.com/huysamen/domains-go/types"
)

const (
	prices = "/domain/reseller/prices/format/tld"
)

type Client struct {
	post types.PostRequest
}

func Create(post types.PostRequest) *Client {
	return &Client{post: post}
}
