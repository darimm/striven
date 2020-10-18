package striven

import (
	"encoding/json"
	"fmt"
)

type poolsFunc struct{}

// PoolAPIResult is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-pools
type PoolAPIResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
}

// PoolsAPIResult is the collection of Pool needed to return all available Pools
type PoolsAPIResult []PoolAPIResult

// GetAll (Pools) is an implementition of https://api.striven.com/Help/Api/GET-v1-pools
func (*poolsFunc) GetAll() (PoolsAPIResult, error) {

	resp, err := stv.apiGet("v1/pools")
	if resp.StatusCode() != 200 || err != nil {
		return PoolsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Pools", resp.StatusCode())
	}
	var r PoolsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return PoolsAPIResult{}, err
	}
	return r, nil
}
