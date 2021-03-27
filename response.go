package main

type GetPenaltiesResponse struct {
	Response struct {
		Penalties struct {
			Penalty []penalty `json:"penalty"`
		} `json:"penalties"`
		Result struct {
			ResCode string `json:"resCode"`
		} `json:"result"`
	} `json:"response"`
}

// IsOk проверяет успешность выполнения запроса
func (resp *GetPenaltiesResponse) IsOk() bool {
	if resp.Response.Result.ResCode == "0" {
		return true
	}
	return false
}

// HasResults проверяет есть ли хотя бы один штраф в ответе на GetPenalties запрос
func (resp *GetPenaltiesResponse) HasResults() bool {
	if len(resp.Response.Penalties.Penalty) > 0 {
		return true
	}
	return false
}

type penalty struct {
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
