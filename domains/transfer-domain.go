package domains

import "encoding/json"

type DomainTransferReq struct {
	Key                     string `json:"key"`
	SLD                     string `json:"sld"`
	TLD                     string `json:"tld"`
	RegistrantName          string `json:"registrantName"`
	RegistrantEmail         string `json:"registrantEmail"`
	RegistrantContactNumber string `json:"registrantContactNumber"`
	RegistrantAddress1      string `json:"registrantAddress1"`
	RegistrantAddress2      string `json:"registrantAddress2"`
	RegistrantPostalCode    string `json:"registrantPostalCode"`
	RegistrantCountry       string `json:"registrantCountry"`
	RegistrantCompany       string `json:"registrantCompany"`
	RegistrantCity          string `json:"registrantCity"`
	RegistrantProvince      string `json:"registrantProvince"`
	EPPKey                  string `json:"eppKey"`
	NS1                     string `json:"ns1"`
	NS2                     string `json:"ns2"`
}

type DomainTransferRsp struct {
	ReturnCode int    `json:"intReturnCode"`
	Message    string `json:"strMessage"`
	DomainName string `json:"strDomainName"`
	Registrant string `json:"strRegistrant"`
}

func (c *Client) TransferDomain(payload DomainTransferReq) (*DomainTransferRsp, error) {
	body, err := c.post(transferDomainPath, payload)

	if err != nil {
		return nil, err
	}

	rsp := new(DomainTransferRsp)
	err = json.Unmarshal(body, rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
