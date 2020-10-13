package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type salesOrderTypesFunc struct{}

// SalesOrderTypesAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-sales-order-types_excludeContractManagedTypes
type SalesOrderTypesAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		IsDefault bool   `json:"isDefault"`
		Active    bool   `json:"active"`
	}
}

// GetAll (SalesOrderTypes) returns a list of Sales Order types, filtered to exclude contract managed order types if the passed parameter is true
func (*salesOrderTypesFunc) GetAll(excludeContractManagedTypes bool) (SalesOrderTypesAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/sales-order-types?excludeContractManagedTypes=%s", strconv.FormatBool(excludeContractManagedTypes)))
	if err != nil {
		return SalesOrderTypesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Sales Order Types", resp.StatusCode())
	}
	var r SalesOrderTypesAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
