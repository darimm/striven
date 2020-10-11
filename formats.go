package striven

import (
	"encoding/json"
	"fmt"
)

// InvoiceFormat is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-invoice-formats
type InvoiceFormat struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// InvoiceFormats is the collection of InvoiceFormat needed to return all formats
type InvoiceFormats []InvoiceFormat

// InvoiceFormatsGet is an implementition of https://api.striven.com/Help/Api/GET-v1-invoice-formats
func (s *Striven) InvoiceFormatsGet() (InvoiceFormats, error) {

	resp, err := s.apiGet("v1/invoice-formats")
	if resp.StatusCode() != 200 || err != nil {
		return InvoiceFormats{}, fmt.Errorf("Response Status Code: %d, Error retrieving Invoice Formats", resp.StatusCode())
	}
	var r InvoiceFormats
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return InvoiceFormats{}, err
	}
	return r, nil
}
