package striven

import (
	"encoding/json"
	"fmt"
)

type tasksFunc struct{}

// TasksAPIResult is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-Tasks-TaskID
type TasksAPIResult struct {
	ID                      int       `json:"id"`
	TaskName                string    `json:"taskName"`
	TaskTypeID              int       `json:"taskTypeId"`
	TaskTypeName            string    `json:"taskTypeName"`
	AccountID               int       `json:"accountId"`
	AccountName             string    `json:"accountName"`
	LocationID              int       `json:"locationId"`
	LocationName            string    `json:"locationName"`
	TaskDescription         string    `json:"taskDescription"`
	BudgetHours             float64   `json:"budgetHours"`
	TaskPercentComplete     float64   `json:"taskPercentComplete"`
	DateRequested           Timestamp `json:"dateRequested"`
	DesiredEndDate          Timestamp `json:"desiredEndDate"`
	DesiredStartDate        Timestamp `json:"desiredStartDate"`
	StatusID                int       `json:"statusId"`
	Status                  string    `json:"status"`
	PriorityID              int       `json:"priorityId"`
	PriorityName            string    `json:"priorityName"`
	DivisionID              int       `json:"divisionId"`
	AssignedTo              string    `json:"assignedTo"`
	IsRecurring             bool      `json:"isRecurring"`
	IsSticky                bool      `json:"isSticky"`
	RequestedBy             int       `json:"requestedBy"`
	RequestedByName         string    `json:"requestedByName"`
	OrderID                 int       `json:"orderId"`
	OrderName               string    `json:"orderName"`
	ProjectID               int       `json:"projectId"`
	ProjectName             string    `json:"projectName"`
	MilestoneID             int       `json:"milestoneId"`
	MilestoneName           string    `json:"milestoneName"`
	UseSubcontractor        bool      `json:"useSubcontractor"`
	AssignedVendorID        int       `json:"assignedVendorId"`
	AssignedVendorName      string    `json:"assignedVendorName"`
	PoID                    int       `json:"poId"`
	PoName                  string    `json:"poName"`
	PoStatus                string    `json:"poStatus"`
	SubContractedWorkStatus string    `json:"subContractedWorkStatus"`
	CreatedBy               int       `json:"createdBy"`
	DateCreated             Timestamp `json:"dateCreated"`
	ModifiedBy              int       `json:"modifiedBy"`
	DateModified            Timestamp `json:"dateModified"`
	AttachmentCount         int       `json:"attachmentCount"`
	NotesLogCount           int       `json:"notesLogCount"`
	CustomFields            []struct {
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
}

// GetByID (Tasks) returns a single Task
func (*tasksFunc) GetByID(taskID int) (TasksAPIResult, error) {

	resp, err := stv.apiGet(fmt.Sprintf("v1/tasks/%d", taskID))
	if resp.StatusCode() != 200 || err != nil {
		return TasksAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Task ID: %d", resp.StatusCode(), taskID)
	}
	var r TasksAPIResult
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
