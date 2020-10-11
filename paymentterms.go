package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// PaymentTerms is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-payment-terms_excludeDiscounts
type PaymentTerms struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		TermDays  int    `json:"termDays"`
		IsDefault bool   `json:"isDefault"`
		Active    bool   `json:"active"`
	}
}

// PaymentTermsGet returns a list of Payment Terms, filtered to exclude discounts if the passed parameter is true
func (s *Striven) PaymentTermsGet(excludeDiscounts bool) (PaymentTerms, error) {

	resp, err := s.apiGet(fmt.Sprintf("v1/payment-terms?excludeDiscounts=%s", strconv.FormatBool(excludeDiscounts)))
	if err != nil {
		return PaymentTerms{}, fmt.Errorf("Response Status Code: %d, Error retrieving Payment Terms", resp.StatusCode())
	}
	var r PaymentTerms
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
