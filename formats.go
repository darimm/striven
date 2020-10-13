package striven

import (
	"encoding/json"
	"fmt"
)

type invoiceFormatsFunc struct{}

// InvoiceFormatAPIResult is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-invoice-formats
type InvoiceFormatAPIResult struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// InvoiceFormatsAPIResult is the collection of InvoiceFormat needed to return all formats
type InvoiceFormatsAPIResult []InvoiceFormatAPIResult

// GetAll (InvoiceFormats) is an implementition of https://api.striven.com/Help/Api/GET-v1-invoice-formats
func (*invoiceFormatsFunc) GetAll() (InvoiceFormatsAPIResult, error) {

	resp, err := stv.apiGet("v1/invoice-formats")
	if resp.StatusCode() != 200 || err != nil {
		return InvoiceFormatsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Invoice Formats", resp.StatusCode())
	}
	var r InvoiceFormatsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return InvoiceFormatsAPIResult{}, err
	}
	return r, nil
}
