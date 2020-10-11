package striven

import (
	"encoding/json"
	"fmt"
)

// Pool is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-pools
type Pool struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"active"`
}

// Pools is the collection of Pool needed to return all available Pools
type Pools []Pool

// PoolsGet is an implementition of https://api.striven.com/Help/Api/GET-v1-pools
func (s *Striven) PoolsGet() (Pools, error) {

	resp, err := s.apiGet("v1/pools")
	if err != nil {
		return Pools{}, fmt.Errorf("Response Status Code: %d, Error retrieving Pools", resp.StatusCode())
	}
	var r Pools
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return Pools{}, err
	}
	return r, nil
}
