package striven

import (
	"encoding/json"
	"fmt"
)

type billCreditFunc struct{}

//vendorRef is an IDName pair with an added Number field that is a string which is used for synchronization with external systems.
type vendorRef struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Name   string `json:"name"`
}

//longIDNameExtended applies to Anything with a longer than average escription, and contains an extended name field.
type longIDNameExtended struct {
	ExtendedName string `json:"extendedName"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
}

//billCreditLineItem is the structure of a single line item in a BillCreditAPIResult.
type billCreditLineItem struct {
	ID                int                `json:"id"`
	IsExpense         bool               `json:"isExpense"`
	Item              IDNamePair         `json:"item"`
	GLAccount         IDNamePair         `json:"glAccount"`
	Qty               float64            `json:"qty"`
	UnitOfMeasure     longIDNameExtended `json:"unitOfMeasure,omitempty"`
	Cost              float64            `json:"cost"`
	Description       string             `json:"description"`
	Customer          IDNamePair         `json:"customer"`
	Order             IDNamePair         `json:"order,omitempty"`
	Billable          bool               `json:"billable"`
	Billed            bool               `json:"billed"`
	GlCategory        IDNamePair         `json:"glCategory,omitempty"`
	InventoryLocation IDNamePair         `json:"inventoryLocation,omitempty"`
}

//AppliedToTxn is the format for a single bill that a credit was applied to. Almost always a slice value since you can apply a credit to multiple bills
type AppliedToTxn struct {
	TxnID       int       `json:"txnid"` //TODO SS: I need to get a sample of this field to verify the case of the json return.
	Amount      float64   `json:"amount"`
	DateApplied Timestamp `json:"dateApplied"`
}

// BillCreditAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-bill-credits-id
type BillCreditAPIResult struct {
	ID                  int                  `json:"id"`
	TxnNumber           string               `json:"txnNumber"`
	TxnDate             Timestamp            `json:"txnDate"`
	Vendor              vendorRef            `json:"vendor"`
	VendorLocation      IDNamePair           `json:"vendorLocation,omitempty"`
	Memo                string               `json:"memo"`
	Status              IDNamePair           `json:"status"`
	CreditTotal         float64              `json:"creditTotal"`
	UnappliedCredit     float64              `json:"unappliedCredit"`
	ApglAccount         IDNamePair           `json:"apglAccount"`
	LineItemsGLCategory IDNamePair           `json:"lineItemsGLCategory,omitempty"`
	LineItems           []billCreditLineItem `json:"lineItems"`
	AppliedToBills      []AppliedToTxn       `json:"appliedToBills,omitempty"`
	NotesLogCount       int                  `json:"notesLogCount"`
	AttachmentCount     int                  `json:"attachmentCount"`
	DateCreated         string               `json:"dateCreated"`
	CreatedBy           IDNamePair           `json:"createdBy"`
	LastUpdatedDate     Timestamp            `json:"lastUpdatedDate,omitempty"`
	LastUpdatedBy       IDNamePair           `json:"lastUpdatedBy,omitempty"`
	Reviewed            bool                 `json:"reviewed"`
	DateReviewed        Timestamp            `json:"dateReviewed,omitempty"`
	ReviewedBy          IDNamePair           `json:"reviewedBy,omitempty"`
	Currency            StrivenCurrency      `json:"currency"`
}

// GetByID (Tasks) returns a single Task
func (*billCreditFunc) GetByID(billCreditID int) (BillCreditAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/bill-credits/%d", billCreditID))
	if resp.StatusCode() != 200 || err != nil {
		return BillCreditAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Task ID: %d", resp.StatusCode(), billCreditID)
	}
	var r BillCreditAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
