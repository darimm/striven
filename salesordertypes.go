package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type salesOrderTypesFunc struct{}

type salesOrder struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
	Active    bool   `json:"active"`
}

// SalesOrderTypesAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-sales-order-types_excludeContractManagedTypes
type SalesOrderTypesAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []salesOrder
}

// GetAll (SalesOrderTypes) returns a list of Sales Order types, filtered to exclude contract managed order types if the passed parameter is true
func (*salesOrderTypesFunc) GetAll(excludeContractManagedTypes bool) (SalesOrderTypesAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/sales-order-types?excludeContractManagedTypes=%s", strconv.FormatBool(excludeContractManagedTypes)))
	if resp.StatusCode() != 200 || err != nil {
		return SalesOrderTypesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Sales Order Types", resp.StatusCode())
	}
	var r SalesOrderTypesAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return SalesOrderTypesAPIResult{}, err
	}
	return r, nil
}
