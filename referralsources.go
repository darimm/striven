package striven

import (
	"encoding/json"
	"fmt"
)

type referralSourcesFunc struct{}

// ReferralSourcesAPIResult is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-referral-sources
type ReferralSourcesAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}
}

// GetAll (ReferralSources) is an implementition of https://api.striven.com/Help/Api/GET-v1-referral-sources
func (*referralSourcesFunc) GetAll() (ReferralSourcesAPIResult, error) {

	resp, err := stv.apiGet("v1/referral-sources")
	if err != nil {
		return ReferralSourcesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Referral Sources", resp.StatusCode())
	}
	var r ReferralSourcesAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ReferralSourcesAPIResult{}, err
	}
	return r, nil
}
