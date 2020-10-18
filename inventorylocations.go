package striven

import (
	"encoding/json"
	"fmt"
)

type inventoryLocationsFunc struct{}

type inventoryLocationSearchResult struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	FullName string     `json:"fullName"`
	Parent   IDNamePair `json:"parent"`
	Active   bool       `json:"active"`
}

// InventoryLocationsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-inventory-locations
type InventoryLocationsAPIResult struct {
	TotalCount int                             `json:"totalCount"`
	Data       []inventoryLocationSearchResult `json:"data"`
}

// GetAll (InventoryLocations) returns a list of available Inventory Locations
func (*inventoryLocationsFunc) GetAll() (InventoryLocationsAPIResult, error) {

	resp, err := stv.apiGet("v1/inventory-locations")
	if resp.StatusCode() != 200 || err != nil {
		return InventoryLocationsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Inventory Locations", resp.StatusCode())
	}
	var r InventoryLocationsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return InventoryLocationsAPIResult{}, err
	}
	return r, nil
}
