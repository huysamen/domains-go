package domains

import "encoding/json"

type DomainCheckReq struct {
	SLD string `json:"sld"`
	TLD string `json:"tld"`
}

type DomainCheckRsp struct {
	ReturnCode  int    `json:"intReturnCode"`
	UUID        string `json:"strUUID"`
	Message     string `json:"strMessage"`
	UsesEPPKey  string `json:"usesEppKey"`
	IsAvailable bool   `json:"isAvailable,string"`
	Reason      string `json:"strReason"`
	MaxRegTerm  int    `json:"intMaxRegTerm"`
	TLD         string `json:"tld"`
	SLD         string `json:"sld"`
	IsPremium   bool   `json:"isPremium,string"`
}

func (c *Client) CheckDomain(payload DomainCheckReq) (*DomainCheckRsp, error) {
	body, err := c.post(checkDomainPath, &payload)

	if err != nil {
		return nil, err
	}

	rsp := new(DomainCheckRsp)
	err = json.Unmarshal(body, rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
