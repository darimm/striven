package striven

import (
	"encoding/json"
	"fmt"
)

// ItemTypes is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-item-types
type ItemTypes struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
}

// ItemTypesGet returns a list of ItemTypes
func (s *Striven) ItemTypesGet() (ItemTypes, error) {

	resp, err := s.apiGet("v1/item-types")
	if err != nil {
		return ItemTypes{}, fmt.Errorf("Response Status Code: %d, Error retrieving Item Types", resp.StatusCode())
	}
	var r ItemTypes
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
