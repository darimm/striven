package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type paymentTermsFunc struct{}

type paymentTerm struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TermDays  int    `json:"termDays"`
	IsDefault bool   `json:"isDefault"`
	Active    bool   `json:"active"`
}

// PaymentTermsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-payment-terms_excludeDiscounts
type PaymentTermsAPIResult struct {
	TotalCount int           `json:"totalCount"`
	Data       []paymentTerm `json:"data"`
}

// GetAll (PaymentTermsGet) returns a list of Payment Terms, filtered to exclude discounts if the passed parameter is true
func (*paymentTermsFunc) GetAll(excludeDiscounts bool) (PaymentTermsAPIResult, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/payment-terms?excludeDiscounts=%s", strconv.FormatBool(excludeDiscounts)))
	if resp.StatusCode() != 200 || err != nil {
		return PaymentTermsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Payment Terms", resp.StatusCode())
	}
	var r PaymentTermsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return PaymentTermsAPIResult{}, err
	}
	return r, nil
}
