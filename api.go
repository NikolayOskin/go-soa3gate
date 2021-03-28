package soa3gate

import (
	"errors"
	"fmt"
)

// GetPenaltyByBill: Получить бюджетное начисление (штраф, налог, ИПФССП и т.п.) по УИН
func (a3 *A3) GetPenaltyByBill(bill string, extStatus bool) (Penalty, bool, error) {
	resp, err := a3.getPenalties(penaltyByBillReq(bill, extStatus, a3.config.authKey))
	if err != nil {
		return Penalty{}, false, fmt.Errorf("GetPenaltyByBill error: %v", err)
	}
	if !resp.isOk() {
		return Penalty{}, false, fmt.Errorf("GetPenaltyByBill error: %v", resp.Response.Result.ResMessage)
	}
	if !resp.hasResult() {
		return Penalty{}, false, nil
	}
	return resp.Response.Penalties.Penalty[0], true, nil
}

/*
GetPenaltiesByType: Получить бюджетные начисления по документам определенного типа.

Доступный перечень типов: «vu» - номер ВУ; «sts» - номер СТС; «inn» - ИНН физлица; «snils» - СНИЛС;
«birthCert» - свидетельство о рождении; «pass» - паспорт; «ip» - исполнительное производство ФССП,
«rawID» - идентефикатор плательщика в формате ГИС.
*/
func (a3 *A3) GetPenaltiesByType(docType string, searchedDocs []string, extStatus bool) ([]Penalty, error) {
	if extStatus && len(searchedDocs) > 2 {
		// А3 позволяет запрашивать начисления со статусом квитирования не более, чем по двум документам в запросе
		return nil, errors.New("GetPenaltiesByType error: max 2 documents allowed with extStatus true")
	}
	if !isDocTypeValid(docType) {
		return nil, errors.New("GetPenaltiesByType error: invalid doctype")
	}

	var documents []Doc
	for _, doc := range searchedDocs {
		documents = append(documents, Doc{DocPayload{docType, doc}})
	}
	resp, err := a3.getPenalties(penaltiesByDocsReq(documents, extStatus, a3.config.authKey))
	if err != nil {
		return nil, fmt.Errorf("GetPenaltiesByType error: %v", err)
	}
	if !resp.isOk() {
		return nil, fmt.Errorf("GetPenaltiesByType error: %v", resp.Response.Result.ResMessage)
	}
	return resp.Response.Penalties.Penalty, nil
}
