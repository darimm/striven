package striven

import (
	"encoding/json"
	"fmt"
)

type contactsFunc struct{}

// ContactsAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-contacts-id
type ContactsAPIResult struct {
	CustomerAssociations []struct {
		Customer struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"customer"`
		ContactTitle string `json:"contactTitle"`
		IsPrimary    bool   `json:"isPrimary"`
		Active       bool   `json:"active"`
	} `json:"customerAssociations"`
	VendorAssociations []struct {
		Vendor struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		ContactTitle string `json:"contactTitle"`
		IsPrimary    bool   `json:"isPrimary"`
		Active       bool   `json:"active"`
	} `json:"vendorAssociations"`
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phones    []struct {
		ID        int `json:"id"`
		PhoneType struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"phoneType"`
		Number      string `json:"number"`
		Extension   string `json:"extension"`
		IsPreferred bool   `json:"isPreferred"`
		Active      bool   `json:"active"`
	} `json:"phones"`
	Emails []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		IsPrimary bool   `json:"isPrimary"`
		Active    bool   `json:"active"`
	} `json:"emails"`
	Address struct {
		Address1    string  `json:"address1"`
		Address2    string  `json:"address2"`
		Address3    string  `json:"address3"`
		City        string  `json:"city"`
		State       string  `json:"state"`
		PostalCode  string  `json:"postalCode"`
		Country     string  `json:"country"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
		FullAddress string  `json:"fullAddress"`
	} `json:"address"`
	CustomFields []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		FieldType struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"fieldType"`
		SourceID   int    `json:"sourceId"`
		Value      string `json:"value"`
		IsRequired bool   `json:"isRequired"`
	} `json:"customFields"`
}

// GetByID (Contacts) returns a single Contact Element
func (*contactsFunc) GetByID(contactID int) (ContactsAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/contacts/%d", contactID))
	if err != nil {
		return ContactsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Contact", resp.StatusCode())
	}
	var r ContactsAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
