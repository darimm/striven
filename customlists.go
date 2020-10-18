package striven

import (
	"encoding/json"
	"fmt"
)

type customListItemsFunc struct{}

type customListsFunc struct {
	ListItems customListItemsFunc
}

type customListAPIResult struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// CustomListsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-custom-lists
type CustomListsAPIResult struct {
	TotalCount int                   `json:"totalCount"`
	Data       []customListAPIResult `json:"data"`
}

type customListItemAPIResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
	Active    bool   `json:"active"`
}

// CustomListItemsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-custom-lists-id-list-items
type CustomListItemsAPIResult struct {
	TotalCount int                       `json:"totalCount"`
	Data       []customListItemAPIResult `json:"data"`
}

// GetAll (CustomLists) returns a list of available Custom Lists used in the system.
func (*customListsFunc) GetAll() (CustomListsAPIResult, error) {

	resp, err := stv.apiGet("v1/custom-lists")
	if resp.StatusCode() != 200 || err != nil {
		return CustomListsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Custom Lists", resp.StatusCode())
	}

	var r CustomListsAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// GetByID (CustomListItems) returns a list of items in a specific custom list specified by listID
func (*customListItemsFunc) GetByID(listID int) (CustomListItemsAPIResult, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/custom-lists/%d/list-items", listID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomListItemsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving List Items", resp.StatusCode())
	}

	var r CustomListItemsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return CustomListItemsAPIResult{}, err
	}
	return r, nil
}
