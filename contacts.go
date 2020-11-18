package striven

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
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

// ContactsSearchParam is the required Parameter for searching contacts with https://api.striven.com/v1/contacts/search. May be empty to return all contacts.
type ContactsSearchParam struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
	AccountName    string `json:"accountName,omitempty"`
	Active         bool   `json:"active,omitempty"`
	PageIndex      int    `json:"pageIndex,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	SortExperssion string `json:"sortExpression,omitempty"`
	SortOrder      int    `json:"sortOrder,omitempty"`
}

//ContactsSearchAPIResult is a collection of ContactsAPIResult wrapped with a TotalCount, structure is the return value for http://api.striven.com/Help/Api/POST-v1-contacts-search
type ContactsSearchAPIResult struct {
	TotalCount int                   `json:"totalCount"`
	Data       []ContactSearchResult `json:"data"`
}

//ContactSearchResult is the resultant data that's pulled from matching contacts
type ContactSearchResult struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateCreated Timestamp `json:"dateCreated"`
	LastUpdated Timestamp `json:"lastUpdatedDate"`
	Active      bool      `json:"active"`
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

// ContactsParam is the Parameter to be passed for Creating or updating a contact.
type ContactsParam struct {
	ID           int                `json:"id,omitempty"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName,omitempty"`
	Phones       *[]APIPhone        `json:"phones,omitempty"`
	Emails       *[]APIEmailAddress `json:"emails,omitempty"`
	Address      *APIAddress        `json:"address,omitempty"`
	CustomFields *[]APICustomField  `json:"customFields,omitempty"`
}

// UpdatedContact is the returned contact info from a contact creation request. It is functionally identical to a ContactsParam
type UpdatedContact struct {
	ContactsParam
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

// Search returns a collection of Contacts
func (*contactsFunc) Search(param ContactsSearchParam) (ContactsSearchAPIResult, error) {

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(jsonHeaders()).
		SetBody(param).
		Post(fmt.Sprintf("%s%s", StrivenURL, "/v1/contacts/search"))
	if err != nil {
		return ContactsSearchAPIResult{}, err
	}

	var r ContactsSearchAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return ContactsSearchAPIResult{}, err
	}
	return r, nil
}

// Update (Contacts) Updates an existing contact in the system.
func (*contactsFunc) Update(c ContactsParam) (UpdatedContact, error) {

	x, err := json.Marshal(c)
	fmt.Println(string(x))
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(jsonHeaders()).
		SetBody(c).
		Post(fmt.Sprintf("%s%s", StrivenURL, "/v1/contacts"))
	if resp.StatusCode() != 200 || err != nil {
		return UpdatedContact{}, fmt.Errorf("Response Code: %d, Error: %+v", resp.StatusCode(), err)
	}

	var r UpdatedContact
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return UpdatedContact{}, err
	}
	return r, nil
}
