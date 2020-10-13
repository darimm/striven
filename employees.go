package striven

import (
	"encoding/json"
	"fmt"
	"time"
)

type employeesFunc struct{}

// rawEmployee is the structure for a single Employee returned from the API
type rawEmployee struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsSystemuser    bool   `json:"isSystemUser"`
	DateCreated     string `json:"dateCreated"`
	LastUpdatedDate string `json:"lastUpdatedDate"`
}

// EmployeeAPIResult is the structure for a single Employee
type EmployeeAPIResult struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsSystemuser    bool   `json:"isSystemUser"`
	DateCreated     time.Time
	LastUpdatedDate time.Time
}

// EmployeesAPIResult is the structure for employees in Striven.
type EmployeesAPIResult []rawEmployee

// GetAll (Employees) is an implementition of https://api.striven.com/Help/Api/GET-v1-employees Time is returned in UTC
func (*employeesFunc) GetAll() ([]EmployeeAPIResult, error) {

	resp, err := stv.apiGet("v1/employees")
	if resp.StatusCode() != 200 || err != nil {
		return []EmployeeAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving Employees", resp.StatusCode())
	}
	var e EmployeesAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &e)
	if err != nil {
		return []EmployeeAPIResult{}, err
	}

	var r []EmployeeAPIResult
	for _, v := range e {
		dc, _ := time.Parse(time.RFC3339Nano, fmt.Sprintf("%sZ", v.DateCreated))
		lu, _ := time.Parse(time.RFC3339Nano, fmt.Sprintf("%sZ", v.LastUpdatedDate))
		r = append(r, EmployeeAPIResult{
			ID:              v.ID,
			Name:            v.Name,
			Email:           v.Email,
			IsSystemuser:    v.IsSystemuser,
			DateCreated:     dc,
			LastUpdatedDate: lu,
		})

	}
	return r, nil
}
