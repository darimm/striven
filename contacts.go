package striven

import (
	"encoding/json"
	"fmt"
)

type contactsFunc struct{}

type customerContactAssociation struct {
	Customer     IDNamePair `json:"customer"`
	ContactTitle string     `json:"contactTitle"`
	IsPrimary    bool       `json:"isPrimary"`
	Active       bool       `json:"active"`
}

type vendorContactAssociation struct {
	Vendor       IDNamePair
	ContactTitle string `json:"contactTitle"`
	IsPrimary    bool   `json:"isPrimary"`
	Active       bool   `json:"active"`
}

// ContactsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-contacts-id
type ContactsAPIResult struct {
	CustomerAssociations []customerContactAssociation `json:"customerAssociations"`
	VendorAssociations   []vendorContactAssociation   `json:"vendorAssociations"`
	ID                   int                          `json:"id"`
	FirstName            string                       `json:"firstName"`
	LastName             string                       `json:"lastName"`
	Phones               []APIPhone                   `json:"phones"`
	Emails               []APIEmailAddress            `json:"emails"`
	Address              APIAddress                   `json:"address"`
	CustomFields         []APICustomField             `json:"customFields"`
}

// GetByID (Contacts) returns a single Contact Element
func (*contactsFunc) GetByID(contactID int) (ContactsAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/contacts/%d", contactID))
	if resp.StatusCode() != 200 || err != nil {
		return ContactsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Contact", resp.StatusCode())
	}
	var r ContactsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ContactsAPIResult{}, err
	}
	return r, nil
}
