package striven

import (
	"encoding/json"
	"fmt"
)

type industriesFunc struct{}

type industrySearchResult struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// IndustriesAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-industries
type IndustriesAPIResult struct {
	TotalCount int                    `json:"totalCount"`
	Data       []industrySearchResult `json:"data"`
}

// GetAll (Industries) returns a list of available Industries
func (*industriesFunc) GetAll() (IndustriesAPIResult, error) {

	resp, err := stv.apiGet("v1/industries")
	if resp.StatusCode() != 200 || err != nil {
		return IndustriesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Industries", resp.StatusCode())
	}
	var r IndustriesAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return IndustriesAPIResult{}, err
	}
	return r, nil
}
