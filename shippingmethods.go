package striven

import (
	"encoding/json"
	"fmt"
)

// ShippingMethod is the structure for an item in the API return from https://api.striven.com/Help/Api/GET-v1-shipping-methods
type ShippingMethod struct {
	ShippingMethodID int    `json:"shippingMethodId"`
	ShippingMethod   string `json:"shippingMethod"`
	TrackingURL      string `json:"trackingURL"`
	Active           bool   `json:"active"`
}

// ShippingMethods is the collection of ShippingMethod needed to return all available ShippingMethods from the API
type ShippingMethods []ShippingMethod

// ShippingMethodsGet is an implementition of https://api.striven.com/Help/Api/GET-v1-shipping-methods
func (s *Striven) ShippingMethodsGet() (ShippingMethods, error) {

	resp, err := s.apiGet("v1/shipping-methods")
	if err != nil {
		return ShippingMethods{}, fmt.Errorf("Response Status Code: %d, Error retrieving Shipping Methods", resp.StatusCode())
	}
	var r ShippingMethods
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ShippingMethods{}, err
	}
	return r, nil
}
