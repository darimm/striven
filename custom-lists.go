package striven

import (
	"encoding/json"
	"fmt"
)

// CustomLists is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-custom-lists
type CustomLists struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}

// CustomListItems is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-custom-lists-id-list-items
type CustomListItems struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		SortOrder int    `json:"sortOrder"`
		Active    bool   `json:"active"`
	}
}

// CustomListsGet returns a list of available Custom Lists used in the system.
func (s *Striven) CustomListsGet() (CustomLists, error) {

	resp, err := s.apiGet("v1/custom-lists")
	if err != nil {
		return CustomLists{}, err
	}

	var r CustomLists
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// CustomListItemsGet returns a list of items in a specific custom list specified by listID
func (s *Striven) CustomListItemsGet(listID int) (CustomListItems, error) {
	resp, err := s.apiGet(fmt.Sprintf("v1/custom-lists/%d/list-items", listID))
	if err != nil {
		return CustomListItems{}, err
	}

	var r CustomListItems
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
