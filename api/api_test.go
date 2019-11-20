package api

import (
	"fmt"
	"github.com/huysamen/domains-go/reseller"
	"testing"
)

func TestDefault(t *testing.T) {
	api, err := Default()
	if err != nil {
		panic(err)
	}

	prices, err := api.Reseller.Prices(reseller.PricesReq{Currency: "ZAR"})
	if err != nil {
		panic(err)
	}

	fmt.Println(prices)
}
