package striven

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
)

//Bill Credit Status
const (
	BillCreditActive = iota + 165
	BillCreditVoided
)

//Bill Status
const (
	BillStatusPendingReview = iota + 124
	BillStatusToBePaid
	BillStatusPaid
	BillStatusCancelled
	BillStatusDenied
)

//Credit Memo Status
const (
	CreditMemoActive = iota + 163
	CreditMemoVoided
)

//Customer Asset Status
var (
	// AssetStatusOutOfServiceParam is the Static Type for an Out of Service Asset Status
	AssetStatusOutOfServiceParam = IDNamePair{
		ID:   15,
		Name: "Out of Service",
	}
	// AssetStatusInServiceParam is the Static Type for an In Service Asset Status
	AssetStatusInServiceParam = IDNamePair{
		ID:   16,
		Name: "In Service",
	}
	// AssetStatusRetiredParam is the Static Type for a Retired Asset Status
	AssetStatusRetiredParam = IDNamePair{
		ID:   17,
		Name: "Retired",
	}
	// AssetStatusUnsupportedParam is the Static Type for an Unsupported Asset Status
	AssetStatusUnsupportedParam = IDNamePair{
		ID:   148,
		Name: "Unsupported",
	}
)

//Customer/Vendor Status
const (
	StatusProspect = iota + 1
	StatusActive
	StatusDeleted
	StatusLost
)

//Customer/Vendor Types
const (
	TypeCustomer = iota + 1
	TypeVendor
)

//GL Account Types
const (
	GLAccountIncome = iota + 1
	GLAccountExpense
	GLAccountFixedAsset
	GLAccountBank
	GLAccountLoan
	GLAccountCreditCard
	GLAccountEquity
	GLAccountAccountsReceivable
	GLAccountAccountsPayable
	GLAccountCostOfGoodsSold
	GLAccountOtherAsset
	_ // 12 - Unused
	GLAccountOtherLiability
	_ // 14 - Unsused
	GLAccountOtherIncome
	GLAccountOtherExpense
)

//Invoice Status
const (
	InvoiceActive = iota + 160
	InvoiceVoided
)

//Journal Entry Status
const (
	JournalEntryActive = iota + 145
	JournalEntryVoided
)

//Order Status
const (
	OrderIncomplete = iota + 18
	OrderQuoted
	OrderPendingApproval
	OrderDeclined
	OrderApproved
	OrderCanceled
	OrderLost
	OrderInProgress
	_ // 26 - Unused
	OrderCompleted
)

//Order Invoice Status
const (
	OrderInvoiceNo = iota + 167
	OrderInvoicePartial
	OrderInvoiceFull
)

//Order Line Item Status
const (
	OrderLineItemPending = iota + 30
	OrderLineItemOrdered
	OrderLineItemShipped
	OrderLineItemCancelled
)

//Payment Status
const (
	PaymentPending = iota + 69
	PaymentApproved
	PaymentFailed
	PaymentVoided = 75
)

//PO Status
const (
	POStatusIncomplete = iota + 105
	POStatusPendingApproval
	POStatusDeclined
	POStatusApproved
	POStatusInProgress
	POStatusFulfilled
	POStatusCancelled
	POStatusClosed = 132
)

//PO Line Item Status
const (
	POLineItemStatusPending = iota + 113
	POLineItemStatusOrdered
	POLineItemStatusReceived
	POLineItemStatusCanceled
)

//Phone Number Types
const (
	PhoneNumberTypeMobile = iota + 1
	PhoneNumberTypeWork
	PhoneNumberTypeHome
	PhoneNumberTypeFax
)

//Subcontracted Work Status
const (
	SubcontractedWorkPending = iota + 140
	SubcontractedWorkInProgress
	SubcontractedWorkDone
)

//Task Status
const (
	TaskOpen = iota + 48
	_        // 49 - Unused
	TaskDone
	TaskCanceled
	TaskOnHold = 68
)

//Vendor PO Status
const (
	VendorPOPending = iota + 117
	VendorPOAccepted
	VendorPODeclined
	VendorPOPartiallyFulfilled
	VendorPOFulfilled
)

// IDNamePair is used pretty much everywhere in the API.
type IDNamePair struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// APICurrency is the standard structure for displaying currency in Striven
type APICurrency struct {
	CurrencyISOCode string  `json:"currencyISOCode"`
	ExchangeRate    float64 `json:"exchangeRate"`
}

// APICustomField is the structure for Custom Fields in the Striven API
type APICustomField struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FieldType  IDNamePair `json:"fieldType"`
	SourceID   int        `json:"sourceId"`
	Value      string     `json:"value"`
	IsRequired bool       `json:"isRequired"`
}

//APIAddress is the standard format for an API address in Striven
type APIAddress struct {
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
}

//APIEmailAddress is the format of an Email address return from Striven
type APIEmailAddress struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	IsPrimary bool   `json:"isPrimary"`
	Active    bool   `json:"active"`
}

//APIPhone is the format of phone information in the Striven API
type APIPhone struct {
	ID          int        `json:"id"`
	PhoneType   IDNamePair `json:"phoneType,omitempty"`
	Number      string     `json:"number"`
	Extension   string     `json:"extension"`
	IsPreferred bool       `json:"isPreferred"`
	Active      bool       `json:"active"`
}

//APIDateRange is an API date searching construct
type APIDateRange struct {
	DateFrom Timestamp `json:"dateFrom"`
	DateTo   Timestamp `json:"dateTo"`
}

//apiGet powers all of the Striven API read-only GET calls for all submodules
func (s *Striven) apiGet(URI string) (*resty.Response, error) {
	err := s.validateAccessToken()
	if err != nil {
		return nil, err
	}
	client := resty.New()

	resp, err := client.R().
		SetAuthToken(s.Token.AccessToken).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%s%s", StrivenURL, URI))

	returncode := resp.StatusCode()

	if returncode != 200 || err != nil {
		return resp, fmt.Errorf("An Error occurred. See the Response Body for details")
	}

	return resp, nil
}

// Timestamp is a custom time field that can be unmarshalled direct from striven because it's timestamp is not RFC3339
type Timestamp time.Time

const timestampFormat = time.RFC3339 // same as ISO8601

var (
	_ json.Unmarshaler = (*Timestamp)(nil)
	_ json.Marshaler   = (*Timestamp)(nil)
)

// NewTimestamp returns a new Timestamp formatted time based on a real golang time
func NewTimestamp(t time.Time) Timestamp {
	return Timestamp(t)
}

// NowTimestamp returns a new timestamp formatted time with the current time.
func NowTimestamp() Timestamp {
	return NewTimestamp(time.Now())
}

// UnmarshalJSON parses a nullable RFC3339 string into time.
func (t *Timestamp) UnmarshalJSON(v []byte) error {
	str := strings.Trim(string(v), `"`)
	if str == "null" {
		return nil
	}

	tz, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic("Cannot load Time America/New_York")
	}
	r, err := time.ParseInLocation("2006-01-02T15:04:05.999", str, tz)
	if err != nil {
		return err
	}

	*t = Timestamp(r)
	return nil
}

// MarshalJSON returns null if Timestamp is not valid (zero). It returns the
// time formatted in RFC3339 otherwise.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	if !t.IsValid() {
		return []byte("null"), nil
	}

	return json.Marshal(t.Format("2006-01-02T15:04:05.999"))
}

// IsValid just returns whether or not a time is valid.
func (t Timestamp) IsValid() bool {
	return !t.Time().IsZero()
}

// Format is an Implementaion of the built in time lib's Format function
func (t Timestamp) Format(fmt string) string {
	return t.Time().Format(fmt)
}

// Time is an Implementation of the built in time lib's Time function
func (t Timestamp) Time() time.Time {
	return time.Time(t)
}
