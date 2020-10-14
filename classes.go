package striven

import (
	"encoding/json"
	"fmt"
)

type classesFunc struct{}

// ClassesAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-classes
type ClassesAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
		Parent   struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"parent"`
		Active bool `json:"active"`
	}
}

// ClassesGet returns a list of available Classes
func (*classesFunc) GetAll() (ClassesAPIResult, error) {

	resp, err := stv.apiGet("v1/classes")
	if err != nil {
		return ClassesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Classes", resp.StatusCode())
	}
	var r ClassesAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
