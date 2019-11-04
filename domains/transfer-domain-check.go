package domains

import "encoding/json"

type DomainTransferCheckReq struct {
	Key string `json:"key"`
	SLD string `json:"sld"`
	TLD string `json:"tld"`
}

type DomainTransferStatus struct {
	DomainName  string `json:"strDomainName"`
	RequestDate int    `json:"requestDate"`
	Status      string `json:"status"`
}

type DomainTransferCheckRsp struct {
	ReturnCode int                    `json:"intReturnCode"`
	Message    string                 `json:"strMessage"`
	Domains    []DomainTransferStatus `json:"arrDomains"`
}

func (c *Client) TransferDomainCheck(payload *DomainTransferCheckReq) (*DomainTransferCheckRsp, error) {
	body, err := c.post(transferDomainCheckPath, payload)

	if err != nil {
		return nil, err
	}

	rsp := new(DomainTransferCheckRsp)
	err = json.Unmarshal(body, rsp)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}
