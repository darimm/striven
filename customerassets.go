package striven

import (
	"encoding/json"
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

type customerAssetsFunc struct {
	search              CustomerAssetSearchParam
	Types               customerAssetsTypesFunc
	CustomFields        customerAssetsCustomFieldsFunc
	Status              customerAssetsStatusFunc
	MaintenanceSchedule customerAssetsMaintenanceFunc
}

type customerAssetsMaintenanceFunc struct{}

type customerAssetsCustomFieldsFunc struct{}

type customerAssetsTypesFunc struct{}

type customerAssetsStatusFunc struct{}

//CustomerAssetSearchParam is a parameter structure that gets passed to api.striven.com/v1/customer-assets/search
type CustomerAssetSearchParam struct {
	AssetName            string       `json:"assetName,omitempty"`
	StatusID             int          `json:"statusId,omitempty"`
	AssetTypeID          int          `json:"assetTypeId,omitempty"`
	LastUpdatedDateRange APIDateRange `json:"lastUpdatedDateRange,omitempty"`
	PageIndex            int          `json:"pageIndex,omitempty"`
	PageSize             int          `json:"pageSize,omitempty"`
	SortExpression       string       `json:"sortExpression,omitempty"`
	SortOrder            int          `json:"sortOrder,omitempty"`
}

//CustomerAssetStatusParam is the parameter structure that gets passed to api.striven.com/v1/customer-assets/{id}/update-status
type CustomerAssetStatusParam struct {
	Status     IDNamePair `json:"status"`
	StatusNote string     `json:"statusNote"`
}

// CustomerAssetMaintenanceScheduleParam is the parameter structure that gets passed to api.striven.com/v1/customer-assets/{id}/maintenance-schedule
// The "ID" field here is the Maintenance Schedule ID. As far as I can tell there's no way to get it other than to query an asset, so while this
// endpoit can be used to create a Maintenance Schedule, there is a strong possibility that you would just update an existing one and break other
// assets if you're not extremely careful. Use at your own risk. I could also be completely wrong here, there's almost no documentation around this.
// it is also the entire Customer Asset Maintenance Schedule object that is returned as part of an asset.
type CustomerAssetMaintenanceScheduleParam struct {
	ID        int       `json:"id"`
	AssetID   int       `json:"assetId"`
	StartDate Timestamp `json:"startDate"`
	EndDate   Timestamp `json:"endDate"`
	Notes     string    `json:"notes,omitempty"`
	Active    bool      `json:"active"`
}

// CustomerAssetsAPIResult is the Return value of an API customer assets search
type CustomerAssetsAPIResult struct {
	TotalCount int                         `json:"totalCount"`
	Data       []CustomerAssetSearchResult `json:"data"`
}

//CustomerAssetSearchResult is a single customer asset search result
type CustomerAssetSearchResult struct {
	ID              int        `json:"id"`
	AssetName       string     `json:"assetName"`
	AssetTypeID     int        `json:"assetTypeId"`
	Customer        IDNamePair `json:"customer"`
	Status          IDNamePair `json:"sattus"`
	DateCreated     Timestamp  `json:"dateCreated"`
	LastUpdatedDate Timestamp  `json:"lastUpdatedDate"`
}

// CustomerAsset is a single Customer Asset, returned from a Get by ID
type CustomerAsset struct {
	ID                  int                                     `json:"id"`
	AssetName           string                                  `json:"assetName"`
	AssetType           IDNamePair                              `json:"assetType"`
	Customer            IDNamePair                              `json:"customer"`
	CustomerLocation    IDNamePair                              `json:"customerLocation"`
	Status              IDNamePair                              `json:"status"`
	StatusNote          string                                  `json:"statusNote,omitempty"`
	PurchasePrice       float64                                 `json:"purchasePrice,omitempty"`
	VisibleOnPortal     bool                                    `json:"visibleOnPortal"`
	PresentValue        float64                                 `json:"presentValue,omitempty"`
	DatePurchased       Timestamp                               `json:"datePurchased,omitempty"`
	ExpirationDate      Timestamp                               `json:"expirationDate,omitempty"`
	DateCreated         Timestamp                               `json:"dateCreated"`
	CreatedBy           IDNamePair                              `json:"createdBy"`
	LastUpdatedDate     Timestamp                               `json:"lastUpdatedDate,omitempty"`
	LastUpdatedBy       IDNamePair                              `json:"lastUpdatedBy,omitempty"`
	Currency            APICurrency                             `json:"currency,omitempty"`
	CustomFields        []APICustomField                        `json:"customFields,omitempty"`
	MaintenanceSchedule []CustomerAssetMaintenanceScheduleParam `json:"maintenanceSchedule,omitempty"`
}

//CustomerAssetType is a single Customer Asset Type return value
type CustomerAssetType struct {
	ID                     int        `json:"id"`
	AssetType              string     `json:"assetType"`
	HasExpirationDate      bool       `json:"hasExpirationDate"`
	RequiresExpirationDate bool       `json:"requiresExpirationDate"`
	HasDatePurchased       bool       `json:"hasDatePurchased"`
	RequiresDatePurchased  bool       `json:"requiresDatePurchased"`
	HasPresentValue        bool       `json:"hasPresentValue"`
	RequiresPresentValue   bool       `json:"requiresPresentValue"`
	DefaultStatus          IDNamePair `json:"defaultStatus"`
}

//CustomerAssetTypeAPIResult is the return type for a list of customer asset types.
type CustomerAssetTypeAPIResult struct {
	TotalCount int                 `json:"totalCount"`
	Data       []CustomerAssetType `json:"data"`
}

// CustomerAssetCustomFieldsAPIResult is the result of pulling a list of custom fields on a given asset.
type CustomerAssetCustomFieldsAPIResult []struct {
	APICustomField
}

// GetByID (CustomerAssets) returns a single Asset by ID
func (*customerAssetsFunc) GetByID(assetID int) (CustomerAsset, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/customer-assets/%d", assetID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerAsset{}, fmt.Errorf("Response Status Code: %d, Error retrieving Contact", resp.StatusCode())
	}
	var r CustomerAsset
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return CustomerAsset{}, err
	}
	return r, nil
}

// Search returns a collection of CustomerAssets
func (*customerAssetsFunc) Search(param CustomerAssetSearchParam) (CustomerAssetsAPIResult, error) {

	var headers map[string]string

	headers = (map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(headers).
		SetBody(param).
		Post(fmt.Sprintf("%s%s", StrivenURL, "/v1/customer-assets/search"))
	if err != nil {
		return CustomerAssetsAPIResult{}, fmt.Errorf("%+v", err)
	}

	var r CustomerAssetsAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// GetAll (CustomerAssetTypes) returns a list of customer asset types available in a given Striven instance
func (*customerAssetsTypesFunc) GetAll() (CustomerAssetTypeAPIResult, error) {
	resp, err := stv.apiGet("v1/customer-assets/types")
	if resp.StatusCode() != 200 || err != nil {
		return CustomerAssetTypeAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Contact", resp.StatusCode())
	}
	var r CustomerAssetTypeAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return CustomerAssetTypeAPIResult{}, err
	}
	return r, nil
}

// GetByID (CustomerAssetsCustomFields) returns a list of Custom Fields on a given asset
func (*customerAssetsCustomFieldsFunc) GetByID(assetID int) (CustomerAssetCustomFieldsAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/customer-assets/%d/custom-fields", assetID))
	if resp.StatusCode() != 200 || err != nil {
		return CustomerAssetCustomFieldsAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Contact", resp.StatusCode())
	}
	var r CustomerAssetCustomFieldsAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return CustomerAssetCustomFieldsAPIResult{}, err
	}
	return r, nil
}

// UpdateByID Updates the status of an asset using an approved CustomerAssetStatusParam. All are defined in the utility.go file for convenience
// With the naming scheme CustomerAsset<Status>Param
func (*customerAssetsStatusFunc) UpdateByID(ID int, param IDNamePair, note string) (interface{}, error) {

	var headers map[string]string

	headers = (map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	p := CustomerAssetStatusParam{
		Status:     param,
		StatusNote: note,
	}

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(headers).
		SetBody(p).
		Post(fmt.Sprintf("%s%s", StrivenURL, fmt.Sprintf("/v1/customer-assets/%d/update-status", ID)))
	if err != nil {
		return nil, fmt.Errorf("%+v", err)
	}

	var r interface{}
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}

// UpdateByID Updates the status of an asset using an approved CustomerAssetStatusParam. All are defined in the utility.go file for convenience
// With the naming scheme CustomerAsset<Status>Param
func (*customerAssetsMaintenanceFunc) UpdateByID(ID int, param CustomerAssetMaintenanceScheduleParam) (interface{}, error) {

	var headers map[string]string

	headers = (map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	//Enforce that the ID must be correct, regardless of the validity of the object passed in.
	param.AssetID = ID

	client := resty.New()
	resp, err := client.R().
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(headers).
		SetBody(param).
		Post(fmt.Sprintf("%s%s", StrivenURL, fmt.Sprintf("/v1/customer-assets/%d/maintenance-schedule", ID)))
	if err != nil {
		return nil, fmt.Errorf("%+v", err)
	}

	var r interface{}
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
