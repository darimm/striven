package striven

import (
	"encoding/json"
	"fmt"
)

type billCreditNotesFunc struct{}

type billCreditAttachmentsFunc struct{}

type billCreditFunc struct {
	Notes       billCreditNotesFunc
	Attachments billCreditAttachmentsFunc
}

//BillCreditNotes is the return structure for a call to https://api.striven.com/Help/Api/GET-v1-bill-credits-id-notes_PageIndex_PageSize
type BillCreditNotes struct {
	TotalCount int              `json:"totalCount"`
	Data       []billCreditNote `json:"data,omitempty"`
}

type billCreditNote struct {
	Notes     string     `json:"notes"`
	CreatedBy IDNamePair `json:"CreatedBy"`
}

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
	Currency            APICurrency          `json:"currency"`
}

type billCreditAttachment struct {
	ID               int        `json:"id"`
	FileName         string     `json:"fileName"`
	OriginalFileName string     `json:"originalFileName"`
	FilePath         string     `json:"filePath"`
	UploadedBy       IDNamePair `json:"uploadedBy"`
	VisibleOnPortal  bool       `json:"visibleOnPortal"`
	IsDefault        bool       `json:"isDefault"`
	DateCreated      string     `json:"dateCreated"`
}

// BillCreditAttachmentAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-bill-credits-id-attachments
type BillCreditAttachmentAPIResult struct {
	TotalCount int                    `json:"totalCount"`
	Data       []billCreditAttachment `json:"data"`
}

// GetByID (BillCredits) returns a single Bill Credit by ID
func (*billCreditFunc) GetByID(billCreditID int) (BillCreditAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/bill-credits/%d", billCreditID))
	if resp.StatusCode() != 200 || err != nil {
		return BillCreditAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Bill Credit ID: %d", resp.StatusCode(), billCreditID)
	}
	var r BillCreditAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return BillCreditAPIResult{}, err
	}
	return r, nil
}

// GetByID (BillCreditAttachments) returns a collection of Bill Credit Attachments attached to a Bill Credit by Bill Credit ID
func (*billCreditAttachmentsFunc) GetByID(billCreditID int) (BillCreditAttachmentAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/bill-credits/%d/attachments", billCreditID))
	if resp.StatusCode() != 200 || err != nil {
		return BillCreditAttachmentAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Bill Credit Attachments for ID: %d", resp.StatusCode(), billCreditID)
	}
	var r BillCreditAttachmentAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return BillCreditAttachmentAPIResult{}, err
	}
	return r, nil
}

// GetByID (BillCreditNotes) returns a list of Bill Credits Passing a single int specifies the BillCredit ID
// Passing 2 ints will also specify the page to retrieve (of default 100 size pages), a 3rd int will set the
// page size. Any subsequent integers are ignored.
func (*billCreditNotesFunc) GetByID(params ...int) (BillCreditNotes, error) {

	var url string

	switch len(params) {
	case 3:
		url = fmt.Sprintf("v1/bill-credits/%d/notes?PageIndex=%d&PageSize=%d", params[0], params[1], params[2])
	case 2:
		url = fmt.Sprintf("v1/bill-credits/%d/notes?PageIndex=%d&PageSize=%d", params[0], params[1], 100)
	case 1:
		url = fmt.Sprintf("v1/bill-credits/%d/notes?PageIndex=%d&PageSize=%d", params[0], 0, 100)
	default:
		url = fmt.Sprintf("v1/bill-credits/%d/notes?PageIndex=%d&PageSize=%d", params[0], params[1], params[2])
	}

	resp, err := stv.apiGet(url)
	if resp.StatusCode() != 200 || err != nil {
		return BillCreditNotes{}, fmt.Errorf("Response Status Code: %d, Error retrieving Bill Credits", resp.StatusCode())
	}
	var r BillCreditNotes
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return BillCreditNotes{}, err
	}
	return r, nil
}
