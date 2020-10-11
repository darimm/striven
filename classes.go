package striven

import (
	"encoding/json"
	"fmt"
)

// Classes is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-classes
type Classes struct {
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

// ClassesGet returns a list of available Classes
func (s *Striven) ClassesGet() (Classes, error) {

	resp, err := s.apiGet("v1/classes")
	if err != nil {
		return Classes{}, fmt.Errorf("Response Status Code: %d, Error retrieving Classes", resp.StatusCode())
	}
	var r Classes
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
