package striven

import (
	"encoding/json"
	"fmt"
)

type glCategoryFunc struct{}

// GLCategoryAPIResult is the structure for an API return from https://api.striven.com/Help/Api/GET-v1-glcategories
type GLCategoryAPIResult struct {
	CategoryID         int    `json:"CategoryID"`
	CategoryName       string `json:"CategoryName"`
	CategoryFullName   string `json:"CategoryFullName"`
	ParentID           int    `json:"ParentID"`
	ParentCategoryName string `json:"ParentCategoryName"`
	Active             bool   `json:"active"`
}

// GLCategoriesAPIResult is the collection of InvoiceFormat needed to return all formats
type GLCategoriesAPIResult []GLCategoryAPIResult

// GetAll (GLCategoriesGet) is an implementition of https://api.striven.com/Help/Api/GET-v1-glcategories
func (*glCategoryFunc) GetAll() (GLCategoriesAPIResult, error) {

	resp, err := stv.apiGet("v1/glcategories")
	if resp.StatusCode() != 200 || err != nil {
		return GLCategoriesAPIResult{}, fmt.Errorf("Response Status Code: %d, Error retrieving GLCategories", resp.StatusCode())
	}
	var r GLCategoriesAPIResult
	err = json.Unmarshal([]byte(resp.Body()), &r)
	if err != nil {
		return GLCategoriesAPIResult{}, err
	}
	return r, nil
}
