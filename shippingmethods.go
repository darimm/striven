package striven

import (
	"encoding/json"
	"fmt"
)

type shippingMethodsFunc struct{}

// ShippingMethodAPIResult is the structure for an item in the API return from https://api.striven.com/Help/Api/GET-v1-shipping-methods
type ShippingMethodAPIResult struct {
	ShippingMethodID int    `json:"shippingMethodId"`
	ShippingMethod   string `json:"shippingMethod"`
	TrackingURL      string `json:"trackingURL"`
	Active           bool   `json:"active"`
}

// ShippingMethodsAPIResult is the collection of ShippingMethod needed to return all available ShippingMethods from the API
type ShippingMethodsAPIResult []ShippingMethodAPIResult

// GetAll (ShippingMethods) is an implementition of https://api.striven.com/Help/Api/GET-v1-shipping-methods
func (*shippingMethodsFunc) GetAll() (ShippingMethodsAPIResult, error) {

	resp, err := stv.apiGet("v1/shipping-methods")
	if err != nil {
		return ShippingMethodsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Shipping Methods", resp.StatusCode())
	}
	var r ShippingMethodsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ShippingMethodsAPIResult{}, err
	}
	return r, nil
}
