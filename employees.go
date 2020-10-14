package striven

import (
	"encoding/json"
	"fmt"
)

type employeesFunc struct{}

// EmployeeAPIResult is the structure for a single Employee
type EmployeeAPIResult struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	IsSystemuser    bool      `json:"isSystemUser"`
	DateCreated     Timestamp `json:"datecreated"`
	LastUpdatedDate Timestamp `json:"lastUpdatedDate"`
}

// EmployeesAPIResult is the structure for employees in Striven.
type EmployeesAPIResult []EmployeeAPIResult

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

	return e, nil
}
