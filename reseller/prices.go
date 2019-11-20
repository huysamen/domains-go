package reseller

import "encoding/json"

type PricesReq struct {
	Currency string `json:"currency"`
}

type PricesRsp struct {
	ReturnCode int                    `json:"intReturnCode"`
	UUID       string                 `json:"strUUID"`
	Message    string                 `json:"strMessage"`
	Prices     map[string]DomainPrice `json:"arrPrices"`
}

type DomainPrice struct {
	Currency     string  `json:"currency"`
	Registration float64 `json:"registration,string"`
	Renewal      float64 `json:"renewal,string"`
	Transfer     float64 `json:"transfer,string"`
	Redemption   float64 `json:"redemption,string"`
	Premium      float64 `json:"premium,string"`
	OptedOut     bool    `json:"optedOut"`
}

func (c *Client) Prices(payload PricesReq) (*PricesRsp, error) {
	if len(payload.Currency) == 0 {
		payload.Currency = "ZAR"
	}

	body, err := c.post(prices, &payload)

	if err != nil {
		return nil, err
	}

	rsp := new(PricesRsp)
	err = json.Unmarshal(body, rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
