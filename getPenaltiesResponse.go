package soa3gate

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

func (resp *GetPenaltiesResponse) hasResult() bool {
	if len(resp.Response.Penalties.Penalty) > 0 {
		return true
	}
	return false
}
