package domains

import "encoding/json"

type DomainRegisterReq struct {
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
	NS1                     string `json:"ns1"`
	NS2                     string `json:"ns2"`
}

type DomainRegisterRsp struct {
	ReturnCode int    `json:"intReturnCode"`
	Message    string `json:"strMessage"`
	DomainName string `json:"strDomainName"`
	Registrant string `json:"strRegistrant"`
	CreateDate int    `json:"intCrDate"`
	ExpiryDate int    `json:"intExDate"`
}

func (c *Client) RegisterDomain(payload *DomainRegisterReq) (*DomainRegisterRsp, error) {
	body, err := c.api.DoPost(registerDomainUrl, payload)

	if err != nil {
		return nil, err
	}

	rsp := new(DomainRegisterRsp)
	err = json.Unmarshal(body, rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
