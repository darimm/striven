package striven

import (
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/resty.v1"
)

// Employee is the structure for a single Employee
type Employee struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	IsSystemuser    bool   `json:"isSystemUser"`
	DateCreated     time.Time
	LastUpdatedDate time.Time
}

// Employees is the structure for employees in Striven.
type Employees []Employee

// EmployeesGet is an implementition of https://api.striven.com/Help/Api/GET-v1-employees
func (s *Striven) EmployeesGet() (Employees, error) {
	err := s.validateAccessToken()
	if err != nil {
		fmt.Println("Failed to Validate Access Token")
		return Employees{}, err
	}
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(s.Token.AccessToken).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%sv1/employees", StrivenURL))

	if resp.StatusCode() != 200 || err != nil {
		fmt.Println("REST Request to Striven API failed")
		fmt.Printf("Error code %d", resp.StatusCode())
		return Employees{}, err
	}
	var r Employees
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		fmt.Println("JSON Unmarshal failed")
	}
	return r, nil
}
