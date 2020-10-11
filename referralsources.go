package striven

import (
	"encoding/json"
	"fmt"
)

// ReferralSources is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-referral-sources
type ReferralSources struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}

// ReferralSourcesGet is an implementition of https://api.striven.com/Help/Api/GET-v1-referral-sources
func (s *Striven) ReferralSourcesGet() (ReferralSources, error) {

	resp, err := s.apiGet("v1/referral-sources")
	if err != nil {
		return ReferralSources{}, fmt.Errorf("Response Status Code: %d, Error retrieving Referral Sources", resp.StatusCode())
	}
	var r ReferralSources
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ReferralSources{}, err
	}
	return r, nil
}
