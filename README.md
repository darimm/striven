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
striven.New(CustomerID,CustomerSecret)
- [X] Classes  
ClassesGet() implements https://api.striven.com/v1/classes
- [ ] Customers  
CustomersGetContentGroups(clientID int) implements https://api.striven.com/v1/customers/{clientID}/hub/content-groups
- [X] CustomList  
CustomListsGet() implements https://api.striven.com/v1/custom-lists
CustomListItemsGet(ListID int) implements https://api.striven.com/v1/custom-lists/{ListID}/list-items
- [X] Employees  
EmployeesGet() implements https://api.striven.com/v1/employees

APIs Not Implemented

- [ ] BillCredits  
- [ ] Bills  
- [ ] Categories  
- [ ] Contacts  
- [ ] CreditMemos  
- [ ] CustomerAssets  
- [ ] Formats  
- [ ] GLAcconuts  
- [ ] GLCategories  
- [ ] Industries  
- [ ] InventoryLocations  
- [ ] Invoices  
- [ ] Items  
- [ ] ItemTypes  
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
