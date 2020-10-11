package striven

import (
	"encoding/json"
	"fmt"
)

// InventoryLocations is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-inventory-locations
type InventoryLocations struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
		Parent   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		Active bool `json:"active"`
	}
}

// InventoryLocationsGet returns a list of available Inventory Locations
func (s *Striven) InventoryLocationsGet() (InventoryLocations, error) {

	resp, err := s.apiGet("v1/inventory-locations")
	if err != nil {
		return InventoryLocations{}, fmt.Errorf("Response Status Code: %d, Error retrieving Inventory Locations", resp.StatusCode())
	}
	var r InventoryLocations
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
