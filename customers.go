package striven

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type customersFunc struct {
	ContentGroups contentGroupsFunc
}

type contentGroupsFunc struct{}

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
func SetClientID(c int) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.ClientID = c
	}
}

//SetGroupID Functional Option for CHD Constructor GroupID
func SetGroupID(g int) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.GroupID = g
	}
}

//SetContentGroupName Option for CHD Constructor Group Name
func SetContentGroupName(g string) CustomersHubDocOption {
	return func(h *CustomersHubDoc) {
		h.ContentGroupName = g
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

//NewCustomersHubDoc is the Constructor for a default CustomersHubDoc
func NewCustomersHubDoc(opts ...CustomersHubDocOption) *CustomersHubDoc {
	const (
		defaultClientID               = 1
		defaultGroupID                = 0
		defaultContentGroupName       = "Uncategorized"
		defaultOverwriteExistingFiles = false
		defaultVisibleOnPortal        = false
	)

	n := &CustomersHubDoc{
		ClientID:               defaultClientID,
		GroupID:                defaultGroupID,
		ContentGroupName:       defaultContentGroupName,
		OverwriteExistingFiles: defaultOverwriteExistingFiles,
		VisibleOnPortal:        defaultVisibleOnPortal,
	}

	// Iterate over each option provided
	for _, opt := range opts {
		// Call the option giving the above instance of CustomersHubDoc as the arguement
		opt(n)
	}

	return n
}

//TODO: This should be a method of the CustomersHubDoc struct.
//UploadClientHubFile is the function to Upload a document to a Client Hub
func (chd *CustomersHubDoc) UploadClientHubFile(remoteFileName string, localFilePath string) (int, error) {
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

	err := s.validateAccessToken()
	if err != nil {
		return 401, err
	}

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(s.Token.AccessToken).
		SetHeaders(headers).
		SetFile(remoteFileName, localFilePath).
		Post(URL)
	if resp.StatusCode() != 200 || err != nil {
		return resp.StatusCode(), fmt.Errorf("Response Status Code: %d, Error uploading document", resp.StatusCode())
	}
	return resp.StatusCode(), nil
}
