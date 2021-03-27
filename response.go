package main

type GetPenaltiesResponse struct {
	Response struct {
		Penalties struct {
			Penalty []Penalty `json:"Penalty"`
		} `json:"penalties"`
		Result struct {
			ResCode    string `json:"resCode"`
			ResMessage string `json:"resMessage"`
		} `json:"result"`
	} `json:"response"`
}

func (resp *GetPenaltiesResponse) isOk() bool {
	if resp.Response.Result.ResCode == "0" {
		return true
	}
	return false
}

func (resp *GetPenaltiesResponse) hasAny() bool {
	if len(resp.Response.Penalties.Penalty) > 0 {
		return true
	}
	return false
}

type Penalty struct {
	AddInfo      string `json:"addInfo"`
	Amount       string `json:"amount"`
	AmountToPay  string `json:"amountToPay"`
	BillDate     string `json:"billDate"`
	BillNumber   string `json:"billNumber"`
	DiscountDate string `json:"discountDate"`
	DiscountSize string `json:"discountSize"`
	DocName      string `json:"docName"`
	DocNumber    string `json:"docNumber"`
	PayStatus    string `json:"payStatus"`
	PayerName    string `json:"payerName"`
	Quittance    string `json:"quittance"`
	RegCert      string `json:"regCert"`
	ValidUntil   string `json:"validUntil"`
}
