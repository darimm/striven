package striven

import (
	"encoding/json"
	"fmt"
)

// GLCategory is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-glcategories
type GLCategory struct {
	CategoryID         int    `json:"CategoryID"`
	CategoryName       string `json:"CategoryName"`
	CategoryFullName   string `json:"CategoryFullName"`
	ParentID           int    `json:"ParentID"`
	ParentCategoryName string `json:"ParentCategoryName"`
	Active             bool   `json:"active"`
}

// GLCategories is the collection of InvoiceFormat needed to return all formats
type GLCategories []GLCategory

// GLCategoriesGet is an implementition of https://api.striven.com/Help/Api/GET-v1-glcategories
func (s *Striven) GLCategoriesGet() (GLCategories, error) {

	resp, err := s.apiGet("v1/glcategories")
	if resp.StatusCode() != 200 || err != nil {
		return GLCategories{}, fmt.Errorf("Response Status Code: %d, Error retrieving Refresh Token", resp.StatusCode())
	}
	var r GLCategories
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return GLCategories{}, err
	}
	return r, nil
}
