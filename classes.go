package striven

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

// Classes is the overall structure for an API return from https://api.striven.com/Help/Api/GET-v1-classes
type Classes struct {
	TotalCount int `json:"totalCount"`
	Data       struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"fullName"`
	}
}

// ClassesGet returns a list of Hub content groups for a given Client.
func (s *Striven) ClassesGet() (Classes, error) {
	err := s.validateAccessToken()
	if err != nil {
		return Classes{}, err
	}
	client := resty.New()
	resp, err := client.R().
		SetAuthToken(s.Token.AccessToken).
		SetHeader("Content-Type", "application/json").
		Get(fmt.Sprintf("%sv1/classes", StrivenURL))

	if resp.StatusCode() != 200 || err != nil {
		return Classes{}, err
	}
	var r Classes
	json.Unmarshal([]byte(resp.Body()), &r)
	return r, nil
}
