Golang implementation of the Striven API ( https://api.striven.com ) 
This project is in no way officially affiliated with Striven.

Example Code: 
```
package main

import (
  "fmt"
  
  "github.com/darimm/striven"
)

func main() {
  s := striven.New("MYCLIENTID", "MYCLIENTSECRET")
  fmt.println(s.Token.AccessToken)
 
  resp, err := s.EmployeesGet()
  
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(resp)
 }
 ```

Current Status: Incomplete.

APIs Implemented (Checkmark means Completely implemented)

- [X] Access Tokens  
`New(CustomerID,CustomerSecret)` implements https://api.striven.com/v1/apitoken  
- [X] Classes  
`ClassesGet()` implements https://api.striven.com/v1/classes  
- [ ] Customers  
`CustomersGetContentGroups(clientID int)` implements https://api.striven.com/v1/customers/{clientID}/hub/content-groups  
Implementation of https://api.striven.com/Help/Api/POST-v1-customers-id-hub-content-groups-groupId-documents exists, but needs some updating to be clear on it's use.  
- [X] CustomList  
`CustomListsGet()` implements https://api.striven.com/v1/custom-lists  
`CustomListItemsGet(ListID int)` implements https://api.striven.com/v1/custom-lists/{ListID}/list-items  
- [X] Employees  
`EmployeesGet()` implements https://api.striven.com/v1/employees  
- [X] Formats  
`InvoiceFormatsGet()` implements https://api.striven.com/Help/Api/GET-v1-invoice-formats  
- [X] GLCategories  
`GLCategoriesGet()` implements https://api.striven.com/Help/Api/GET-v1-glcategories
- [X] ItemTypes  
`ItemTypesGet()` implements https://api.striven.com/Help/Api/GET-v1-item-types  


APIs Not Implemented

- [ ] BillCredits  
- [ ] Bills  
- [ ] Categories  
- [ ] Contacts  
- [ ] CreditMemos  
- [ ] CustomerAssets  
- [ ] GLAcconuts  

- [ ] Industries  
- [ ] InventoryLocations  
- [ ] Invoices  
- [ ] Items  
- [ ] JournalEntries  
- [ ] Opportunities  
- [ ] Payments  
- [ ] PaymentTerms  
- [ ] Pools  
- [ ] Purchase Orders  
- [ ] ReferralSources  
- [ ] SalesOrders  
- [ ] SalesOrderTypes  
- [ ] ShippingMethods  
- [ ] Tasks  
- [ ] UserInfo  
- [ ] Vendors  
