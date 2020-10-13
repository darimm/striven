package striven

import (
	"encoding/json"
	"fmt"
)

type itemTypesfunc struct{}

// ItemTypesAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-item-types
type ItemTypesAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
}

// GetAll (ItemTypes) returns a list of ItemTypes
func (*itemTypesfunc) GetAll() (ItemTypesAPIResult, error) {

	resp, err := stv.apiGet("v1/item-types")
	if err != nil {
		return ItemTypesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Item Types", resp.StatusCode())
	}
	var r ItemTypesAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
