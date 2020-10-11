package striven

import (
	"encoding/json"
)

// Industries is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-industries
type Industries struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}

// IndustriesGet returns a list of available Industries
func (s *Striven) IndustriesGet() (Industries, error) {

	resp, err := s.apiGet("v1/industries")
	if err != nil {
		return Industries{}, err
	}
	var r Industries
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
