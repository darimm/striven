package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// SalesOrderTypes is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-sales-order-types_excludeContractManagedTypes
type SalesOrderTypes struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		IsDefault bool   `json:"isDefault"`
		Active    bool   `json:"active"`
	}
}

// SalesOrderTypesGet returns a list of Sales Order types, filtered to exclude contract managed order types if the passed parameter is true
func (s *Striven) SalesOrderTypesGet(excludeContractManagedTypes bool) (SalesOrderTypes, error) {

	resp, err := s.apiGet(fmt.Sprintf("v1/sales-order-types?excludeContractManagedTypes=%s", strconv.FormatBool(excludeContractManagedTypes)))
	if err != nil {
		return SalesOrderTypes{}, fmt.Errorf("Response Status Code: %d, Error retrieving Sales Order Types", resp.StatusCode())
	}
	var r SalesOrderTypes
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
