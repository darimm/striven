package striven

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/segmentio/ksuid"
)

type customersFunc struct {
	ContentGroups contentGroupsFunc
	Contacts      customersContentContactsFunc
}

type contentGroupsFunc struct {
	Document CustomersHubDoc
}

type customersContentContactsFunc struct {
}

// CustomerDetail is the structure of a single Customer from the customers APi
type CustomerDetail struct {
	ID                    int              `json:"id"`
	Name                  string           `json:"name"`
	Number                string           `json:"number,omitempty"`
	IsVendor              bool             `json:"isVendor,omitempty"`
	IsConsumerAccount     bool             `json:"isConsumerAccount,omitempty"`
	PrimaryContact        IDNamePair       `json:"primaryContact,omitempty"`
	Status                IDNamePair       `json:"status"`
	Categories            []IDNamePair     `json:"categories,omitempty"`
	ReferralSource        IDNamePair       `json:"referralSource,omitempty"`
	Industry              IDNamePair       `json:"industry,omitempty"`
	CustomerSince         string           `json:"customerSince,omitempty"` //date?
	OnCreditHold          bool             `json:"onCreditHold,omitempty"`
	CreditLimit           float64          `json:"creditLimit,omitempty"`
	WebSite               string           `json:"webSite,omitempty"`
	IsTaxExempt           bool             `json:"isTaxExempt,omitempty"`
	IsFinanceChargeExempt bool             `json:"isFinanceChargeExempt,omitempty"`
	PaymentTerm           IDNamePair       `json:"paymentTerm,omitempty"`
	BillToLocation        IDNamePair       `json:"billToLocation,omitempty"`
	ShipToLocation        IDNamePair       `json:"shipToLocation,omitempty"`
	Phones                []APIPhone       `json:"phones,omitempty"`
	PrimaryAddress        APIAddress       `json:"primaryAddress,omitempty"`
	PriceList             IDNamePair       `json:"priceList"`
	CustomFields          []APICustomField `json:"customFields,omitempty"`
	DateCreated           string           `json:"dateCreated,omitempty"`
	CreatedBy             IDNamePair       `json:"createdBy,omitempty"`
	LastUpdatedDate       string           `json:"lastUpdatedDate,omitempty"`
	LastUpdatedBy         IDNamePair       `json:"lastUpdatedBy,omitempty"`
	Currency              APICurrency      `json:"currency,omitempty"`
}

// New (Customers) will create a new Customer
// Use ID = 0 to create a customer, use existing ID to update. Customer Name, Status are required.
// Available Categories, ReferralSource, Industry can be pulled using different API's listed in API Help Page.
// Please ensure all the properties are correctly supplied in the request body when performing and update,
// This is an full update so existing customer will be updated fully with the object being supplied with this request
// Also ensure correct PhoneID, EmailID are being used when updating a customer.
// If incorrect ID's are used you will get a validation error. If you wish to add new Phone/Email use 0 as ID.
func (*customersFunc) New(customer CustomerDetail) (interface{}, error) { //(CustomerDetail, error) {
	/*
		fmt.Println("About to Marshal JSON")
		reqBody, err := json.Marshal(customer)
		if err != nil {
			return CustomerDetail{}, fmt.Errorf("Unable to Marshal object to JSON, error %v", err)
		}
		fmt.Println("Successfully Marshalled JSON")
	*/
	var headers map[string]string

	headers = (map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(headers).
		SetBody(customer).
		Post(fmt.Sprintf("%s%s", StrivenURL, "/v1/customers"))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerDetail{}, fmt.Errorf("Response Status Code: %d, Error uploading document", resp.StatusCode())
	}

	var r interface{}
	//var r CustomerDetail
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// GetById (Customers) will return a single customer structure.
func (*customersFunc) GetByID(customerID int) (CustomerDetail, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/customers/%d", customerID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerDetail{}, fmt.Errorf("Response Status Code: %d, Error retrieving Client", resp.StatusCode())
	}
	var r CustomerDetail
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

type customerContactAPIResult struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Title     string `json:"title"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	IsPrimary bool   `json:"isPrimary"`
	Active    bool   `json:"active"`
}

//CustomersContactAPIResult is the structure of a request to Striven to /v1/customers/{ID}/contacts
type CustomersContactAPIResult struct {
	TotalCount int                        `json:"totalCount"`
	Data       []customerContactAPIResult `json:"data"`
}

//GetByCustomerID (Contacts) Returns a list of contacts associated with a customer
func (*customersContentContactsFunc) GetByCustomerID(customerID int) (CustomersContactAPIResult, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/customers/%d/contacts", customerID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomersContactAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Client", resp.StatusCode())
	}
	var r CustomersContactAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

type customerHubContentGroupAPIResult struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
}

// CustomersHubContentGroupAPIResult is the structure of a request to Striven to GetContentGroups()
type CustomersHubContentGroupAPIResult struct {
	TotalCount int                                `json:"totalCount"`
	Data       []customerHubContentGroupAPIResult `json:"data"`
}

// CustomersGetContentGroups returns a list of Hub content groups for a given Client.
func (*contentGroupsFunc) GetByID(clientID int) (CustomersHubContentGroupAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/customers/%d/hub/content-groups", clientID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomersHubContentGroupAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Hub Content Groups", resp.StatusCode())
	}
	var r CustomersHubContentGroupAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

//CustomersHubDoc is a structure for file upload headers and options
type CustomersHubDoc struct {
	ClientID               int
	GroupID                int
	ContentGroupName       string
	OverwriteExistingFiles bool
	VisibleOnPortal        bool
}

//CustomersHubDocOption prototype definiton
type CustomersHubDocOption func(*CustomersHubDoc)

//SetClientID Functional Option for CHD Constructor ClientID
func SetClientID(clientID int) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.ClientID = clientID
	}
}

//SetGroupID Functional Option for CHD Constructor GroupID
func SetGroupID(groupID int) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.GroupID = groupID
	}
}

//SetContentGroupName Option for CHD Constructor Group Name
func SetContentGroupName(groupName string) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.ContentGroupName = groupName
	}
}

//IsOverwriteEnabled CHD Option to allow overwriting files
func IsOverwriteEnabled() CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.OverwriteExistingFiles = true
	}
}

//IsVisibleOnPortal CHD Option to allow customers to see these documents
func IsVisibleOnPortal() CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.VisibleOnPortal = true
	}
}

//Upload is the Constructor and uploader for a default CustomersHubDoc
func (chd *CustomersHubDoc) Upload(localFilePath string, opts ...CustomersHubDocOption) (int, error) {
	const (
		defaultClientID               = 1
		defaultGroupID                = 0
		defaultContentGroupName       = "Uncategorized"
		defaultOverwriteExistingFiles = false
		defaultVisibleOnPortal        = false
	)

	chd = &CustomersHubDoc{
		ClientID:               defaultClientID,
		GroupID:                defaultGroupID,
		ContentGroupName:       defaultContentGroupName,
		OverwriteExistingFiles: defaultOverwriteExistingFiles,
		VisibleOnPortal:        defaultVisibleOnPortal,
	}

	// Iterate over each option provided
	for _, opt := range opts {
		// Call the option giving the above instance of CustomersHubDoc as the arguement
		opt(chd)
	}

	httpResponseCode, err := chd.uploadClientHubFile(localFilePath)
	return httpResponseCode, err

}

// uploadClientHubFile is the function to Upload a document to a Client Hub
func (chd *CustomersHubDoc) uploadClientHubFile(localFilePath string) (int, error) {

	var overwrite string = "true"
	if !chd.OverwriteExistingFiles {
		overwrite = "false"
	}

	var visible string = "true"
	if !chd.VisibleOnPortal {
		visible = "false"
	}
	var headers map[string]string
	var URL string
	if chd.GroupID != 0 {
		URL = fmt.Sprintf("%sv1/customers/%d/hub/documents?groupId=%d", StrivenURL, chd.ClientID, chd.GroupID)
		headers = (map[string]string{
			"Content-Type":             "application/json",
			"Accept":                   "application/json",
			"Overwrite-Existing-Files": overwrite,
			"Visible-On-Portal":        visible,
		})
	} else {
		URL = fmt.Sprintf("%sv1/customers/%d/hub/documents", StrivenURL, chd.ClientID)
		headers = (map[string]string{
			"Content-Type":             "application/json",
			"Accept":                   "application/json",
			"Overwrite-Existing-Files": overwrite,
			"Visible-On-Portal":        visible,
			"Content-Group-Name":       chd.ContentGroupName,
		})
	}

	err := stv.validateAccessToken()
	if err != nil {
		return 401, err
	}

	// The Remote File Name in Striven must be a unique identifier for the database. KSUIDs are guarenteed to be unique,
	// so use two of them with a source-string tacked on in case Striven support needs to investigate something sourced from the API
	var remoteFileName = fmt.Sprintf("%s-%s-%s", "striven-go", ksuid.New().String(), ksuid.New().String())
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(headers).
		SetFile(remoteFileName, localFilePath).
		Post(URL)
	if resp.StatusCode() != 200 || err != nil {
		return resp.StatusCode(), fmt.Errorf("Response Status Code: %d, Error uploading document", resp.StatusCode())
	}
	return resp.StatusCode(), nil
}
