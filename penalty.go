package soa3gate

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
