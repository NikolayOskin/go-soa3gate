package soa3gate

type penaltyByBillRequest struct {
	AuthKey
	ExtStatus bool `json:"extStatus"`
	Bill      Bill `json:"bill"`
}

type penaltiesByDocsRequest struct {
	AuthKey
	ExtStatus bool  `json:"extStatus"`
	Docs      []Doc `json:"docs"`
}

type Bill struct {
	Number string `json:"number"`
}

type DocPayload struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Doc struct {
	Payload DocPayload `json:"doc"`
}

func penaltyByBillReq(bill string, extStatus bool, authKey string) *Request {
	r := penaltyByBillRequest{
		AuthKey:   AuthKey{authKey},
		ExtStatus: extStatus,
		Bill:      Bill{bill},
	}
	return newA3Request(r, "getPenalties")
}

func penaltiesByDocsReq(docs []Doc, extStatus bool, authKey string) *Request {
	r := penaltiesByDocsRequest{
		AuthKey:   AuthKey{authKey},
		ExtStatus: extStatus,
		Docs:      docs,
	}
	return newA3Request(r, "getPenalties")
}