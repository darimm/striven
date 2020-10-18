package striven

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"
)

//StrivenURL is the global URL for Striven's API
const StrivenURL string = "https://api.striven.com/"

var stv *Striven = &Striven{}

type strivenToken struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	ExpiresIn      string `json:"expires_in"`
	ExpirationTime time.Time
}

//Striven is the core object in this module. It is an instance of the Striven API Token and the ID and Secret used to connect.
type Striven struct {
	ClientID           string
	ClientSecret       string
	Token              strivenToken
	BillCredits        billCreditFunc
	Classes            classesFunc
	Contacts           contactsFunc
	Customers          customersFunc
	CustomLists        customListsFunc
	Employees          employeesFunc
	InvoiceFormats     invoiceFormatsFunc
	GLCategories       glCategoryFunc
	Industries         industriesFunc
	InventoryLocations inventoryLocationsFunc
	ItemTypes          itemTypesfunc
	PaymentTerms       paymentTermsFunc
	Pools              poolsFunc
	ReferralSources    referralSourcesFunc
	SalesOrderTypes    salesOrderTypesFunc
	ShippingMethods    shippingMethodsFunc
	Tasks              tasksFunc
}

//New is the constructor for an Striven Object. Changing the default ClientID and Secret will also invalidate the token
func New(ID string, Secret string) *Striven {
	stv = &Striven{
		ClientID:     ID,
		ClientSecret: Secret,
		Token:        strivenToken{},
	}
	stv.initializeToken()
	return stv
}

//initializeToken reaches out to the Striven API and uses your ClientID and ClientSecret to Generate an Striven Token
func (s *Striven) initializeToken() error {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(s.ClientID, s.ClientSecret).
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
			"ClientId":   s.ClientID,
		}).
		Post(fmt.Sprintf("%saccesstoken", StrivenURL))

	if resp.StatusCode() != 200 || err != nil {
		s.Token = strivenToken{}
		return fmt.Errorf("Response Status Code: %d, Error retrieving APIToken", resp.StatusCode())
	}

	json.Unmarshal([]byte(resp.Body()), &s.Token)
	expiresIn, err := strconv.Atoi(s.Token.ExpiresIn)
	if err != nil {
		expiresIn = 0
	}
	s.Token.ExpirationTime = time.Now().Add(time.Second * time.Duration(expiresIn))
	return nil
}

//refreshAccessToken
func (s *Striven) refreshAccessToken() error {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBasicAuth(s.ClientID, s.ClientSecret).
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": s.Token.RefreshToken,
		}).
		Post(fmt.Sprintf("%saccesstoken", StrivenURL))

	if resp.StatusCode() != 200 || err != nil {
		return fmt.Errorf("Response Status Code: %d, Error retrieving Refresh Token", resp.StatusCode())
	}

	json.Unmarshal([]byte(resp.Body()), &s.Token)
	expiresIn, err := strconv.Atoi(s.Token.ExpiresIn)
	if err != nil {
		expiresIn = 0
	}
	s.Token.ExpirationTime = time.Now().Add(time.Second * time.Duration(expiresIn))
	return nil
}

//validateAccessToken makes sure that your token is valid
func (s *Striven) validateAccessToken() error {

	if !time.Now().After(s.Token.ExpirationTime.Add(-1 * time.Hour)) {
		return nil
	}
	err := s.refreshAccessToken()
	if err != nil {
		err := s.initializeToken()
		if err != nil {
			return err
		}
	}
	return nil
}
