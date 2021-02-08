package striven

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"
)

type tasksFunc struct{}

// TasksAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-Tasks-TaskID
type TasksAPIResult struct {
	ID                      int              `json:"id"`
	TaskName                string           `json:"taskName"`
	TaskTypeID              int              `json:"taskTypeId"`
	TaskTypeName            string           `json:"taskTypeName"`
	AccountID               int              `json:"accountId"`
	AccountName             string           `json:"accountName"`
	LocationID              int              `json:"locationId"`
	LocationName            string           `json:"locationName"`
	TaskDescription         string           `json:"taskDescription"`
	BudgetHours             float64          `json:"budgetHours"`
	TaskPercentComplete     float64          `json:"taskPercentComplete"`
	DateRequested           Timestamp        `json:"dateRequested"`
	DesiredEndDate          Timestamp        `json:"desiredEndDate"`
	DesiredStartDate        Timestamp        `json:"desiredStartDate"`
	StatusID                int              `json:"statusId"`
	Status                  string           `json:"status"`
	PriorityID              int              `json:"priorityId"`
	PriorityName            string           `json:"priorityName"`
	DivisionID              int              `json:"divisionId"`
	AssignedTo              string           `json:"assignedTo"`
	IsRecurring             bool             `json:"isRecurring"`
	IsSticky                bool             `json:"isSticky"`
	RequestedBy             int              `json:"requestedBy"`
	RequestedByName         string           `json:"requestedByName"`
	OrderID                 int              `json:"orderId"`
	OrderName               string           `json:"orderName"`
	ProjectID               int              `json:"projectId"`
	ProjectName             string           `json:"projectName"`
	MilestoneID             int              `json:"milestoneId"`
	MilestoneName           string           `json:"milestoneName"`
	UseSubcontractor        bool             `json:"useSubcontractor"`
	AssignedVendorID        int              `json:"assignedVendorId"`
	AssignedVendorName      string           `json:"assignedVendorName"`
	PoID                    int              `json:"poId"`
	PoName                  string           `json:"poName"`
	PoStatus                string           `json:"poStatus"`
	SubContractedWorkStatus string           `json:"subContractedWorkStatus"`
	CreatedBy               int              `json:"createdBy"`
	DateCreated             Timestamp        `json:"dateCreated"`
	ModifiedBy              int              `json:"modifiedBy"`
	DateModified            Timestamp        `json:"dateModified"`
	AttachmentCount         int              `json:"attachmentCount"`
	NotesLogCount           int              `json:"notesLogCount"`
	CustomFields            []APICustomField `json:"customFields"`
}

//TaskCreateParams are the parameters for creating a task in Striven
type TaskCreateParams struct {
	TaskName                        string    `json:"taskName,omitempty"`
	TaskTypeID                      int       `json:"taskTypeID,omitempty"`
	PriorityID                      int       `json:"priorityID,omitempty"`
	DueDate                         time.Time `json:"dueDate,omitempty"`
	RequestedByObjectID             int       `json:"requestedByObjectID,omitempty"`
	RequestedByKeyID                int       `json:"requestedByKeyID,omitempty"`
	AccountID                       int       `json:"accountID,omitempty"`
	OrderID                         int       `json:"orderID,omitempty"`
	ProjectID                       int       `json:"projectID,omitempty"`
	MilestoneID                     int       `json:"milestoneID,omitempty"`
	TaskDesc                        string    `json:"taskDesc,omitempty"`
	AssignedToObjectID              int       `json:"assignedToObjectID,omitempty"`
	AssignedToKeyID                 int       `json:"assignedToKeyID,omitempty"`
	AssignToUserByDefault           bool      `json:"assignToUserByDefault,omitempty"`
	DeriveRequestedByUsingEmailFrom bool      `json:"deriveRequestedByUsingEmailFrom,omitempty"`
	RequestedByEmail                string    `json:"requestedByEmail,omitempty"`
	StatusID                        int       `json:"statusID,omitempty"`
}

//TaskCreateResult is the return value when creating a task
type TaskCreateResult struct {
	TaskID             int       `json:"taskID,omitempty"`
	PriorityID         int       `json:"priorityID,omitempty"`
	StartDate          Timestamp `json:"startDate,omitempty"`
	DueDate            Timestamp `json:"dueDate,omitempty"`
	DateCreated        string    `json:"dateCreated,omitempty"`
	AssignedToObjectID int       `json:"assignedToObjectID,omitempty"`
	AssignedToKeyID    int       `json:"assignedToKeyID,omitempty"`
}

// GetByID (Tasks) returns a single Task
func (*tasksFunc) GetByID(taskID int) (TasksAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/tasks/%d", taskID))
	if resp.StatusCode() != 200 || err != nil {
		return TasksAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Task ID: %d", resp.StatusCode(), taskID)
	}
	var r TasksAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return TasksAPIResult{}, err
	}
	return r, nil
}

// Create (CustomerTask) Creates an existing task in the system.
func (*tasksFunc) Create(task TaskCreateParams) (TaskCreateResult, error) {

	err := stv.validateAccessToken()
	if err != nil {
		return TaskCreateResult{}, err
	}
	ctx, done := context.WithCancel(stv.Context)
	defer done()
	client := resty.New()
	resp, err := client.R().
		SetContext(ctx).
		SetAuthToken(stv.Token.AccessToken).
		SetHeaders(jsonHeaders()).
		SetBody(task).
		Post(fmt.Sprintf("%s%s", StrivenURL, "/v1/customer-assets"))
	if resp.StatusCode() != 200 || err != nil {
		return TaskCreateResult{}, fmt.Errorf("Response Code: %d, Error: %+v", resp.StatusCode(), err)
	}

	var r TaskCreateResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
