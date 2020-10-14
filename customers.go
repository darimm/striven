package striven

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type customersFunc struct {
	ContentGroups contentGroupsFunc
	Contacts      customerContentContactsFunc
}

type contentGroupsFunc struct {
	Document CustomersHubDoc
}

type customerContentContactsFunc struct {
}

// CustomerDetailAPIResult is the structure of a single Customer from the customers APi
type CustomerDetailAPIResult struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Number            string `json:"number"`
	IsVendor          bool   `json:"isVendor"`
	IsConsumerAccount bool   `json:"isConsumerAccount"`
	PrimaryContact    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"primaryContact"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Categories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
	ReferralSource struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"referralSource"`
	Industry struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"industry"`
	CustomerSince         string  `json:"customerSince"`
	OnCreditHold          bool    `json:"onCreditHold"`
	CreditLimit           float64 `json:"creditLimit"`
	WebSite               string  `json:"webSite"`
	IsTaxExempt           bool    `json:"isTaxExempt"`
	IsFinanceChargeExempt bool    `json:"isFinanceChargeExempt"`
	PaymentTerm           struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"paymentTerm"`
	BillToLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"billToLocation"`
	ShipToLocation struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"shipToLocation"`
	Phones []struct {
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
	PrimaryAddress struct {
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
	} `json:"primaryAddress"`
	PriceList struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priceList"`
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
	DateCreated string `json:"dateCreated"`
	CreatedBy   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"createdBy"`
	LastUpdatedDate string `json:"lastUpdatedDate"`
	LastUpdatedBy   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"lastUpdatedBy"`
	Currency struct {
		CurrencyISOCode string  `json:"currencyISOCode"`
		ExchangeRate    float64 `json:"exchangeRate"`
	} `json:"currency"`
}

// GetById (Customers) will return a single customer structure.
func (*customersFunc) GetByID(customerID int) (CustomerDetailAPIResult, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/customers/%d", customerID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerDetailAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Client", resp.StatusCode())
	}
	var r CustomerDetailAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

//CustomerContactAPIResult is the structure of a request to Striven to /v1/customers/{ID}/contacts
type CustomerContactAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Title     string `json:"title"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		IsPrimary bool   `json:"isPrimary"`
		Active    bool   `json:"active"`
	} `json:"data"`
}

//GetByCustomerID (Contacts) Returns a list of contacts associated with a customer
func (*customerContentContactsFunc) GetByCustomerID(customerID int) (CustomerContactAPIResult, error) {
	resp, err := stv.apiGet(fmt.Sprintf("v1/customers/%d/contacts", customerID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerContactAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Client", resp.StatusCode())
	}
	var r CustomerContactAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// CustomersHubContentGroupAPIResult is the structure of a request to Striven to GetContentGroups()
type CustomersHubContentGroupAPIResult struct {
	TotalCount int `json:"totalCount"`
	Data       []struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		IsDefault bool   `json:"isDefault"`
	}
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
func (chd *CustomersHubDoc) Upload(remoteFileName string, localFilePath string, opts ...CustomersHubDocOption) (int, error) {
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

	httpResponseCode, err := chd.uploadClientHubFile(remoteFileName, localFilePath)
	return httpResponseCode, err

}

// uploadClientHubFile is the function to Upload a document to a Client Hub
func (chd *CustomersHubDoc) uploadClientHubFile(remoteFileName string, localFilePath string) (int, error) {
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
