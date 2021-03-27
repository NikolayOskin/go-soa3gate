package main

import "fmt"

// GetPenaltiesByBill: Получить штраф по УИН
func (a3 *A3) GetPenaltyByBill(bill string, extStatus bool) (Penalty, bool, error) {
	resp, err := a3.getPenalties(penaltyByBillReq(bill, extStatus, a3.config.authKey))
	if err != nil {
		return Penalty{}, false, fmt.Errorf("GetPenaltiesByBill error: %v", err)
	}
	if !resp.isOk() {
		return Penalty{}, false, fmt.Errorf(resp.Response.Result.ResMessage)
	}
	if !resp.hasAny() {
		return Penalty{}, false, nil
	}
	return resp.Response.Penalties.Penalty[0], true, nil
}

// FetchBySTS: Получить штрафы по СТС
func (a3 *A3) FetchBySTS(sts []string, extStatus bool) (*GetPenaltiesResponse, error) {
	if extStatus && len(sts) > 2 {
		return nil, fmt.Errorf("FetchBySTS error: max 2 sts allowed with extStatus true")
	}

	var stsDocs []Doc
	for _, s := range sts {
		stsDocs = append(stsDocs, Doc{DocPayload{"sts", s}})
	}
	resp, err := a3.getPenalties(penaltiesByDocsReq(stsDocs, extStatus, a3.config.authKey))
	if err != nil {
		return nil, fmt.Errorf("FetchBySTS error: %v", err)
	}
	return resp, nil
}

// FetchByDriverLicenses: Получить штрафы по ВУ
func (a3 *A3) FetchByDriverLicenses(driverLicenses []string, extStatus bool) (*GetPenaltiesResponse, error) {
	if extStatus && len(driverLicenses) > 2 {
		return nil, fmt.Errorf("FetchByDriverLicenses error: max 2 driverLicenses allowed with extStatus true")
	}

	var vuDocs []Doc
	for _, dl := range driverLicenses {
		vuDocs = append(vuDocs, Doc{DocPayload{"vu", dl}})
	}
	resp, err := a3.getPenalties(penaltiesByDocsReq(vuDocs, extStatus, a3.config.authKey))
	if err != nil {
		return nil, fmt.Errorf("FetchByDriverLicenses error: %v", err)
	}
	return resp, nil
}

// GetPenaltiesByDocs: Получить штрафы по документам (см. Doc)
func (a3 *A3) GetPenaltiesByDocs(docs []Doc, extStatus bool) (*GetPenaltiesResponse, error) {
	if extStatus && len(docs) > 2 {
		return nil, fmt.Errorf("GetPenaltiesByDocs error: max 2 docs allowed with extStatus true")
	}

	resp, err := a3.getPenalties(penaltiesByDocsReq(docs, extStatus, a3.config.authKey))
	if err != nil {
		return nil, fmt.Errorf("GetPenaltiesByDocs error: %v", err)
	}
	return resp, nil
}
