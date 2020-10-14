Golang implementation of the Striven API ( https://api.striven.com ) 
This project is in no way officially affiliated with Striven.

Example Code (Note - this is very bad form. Do not put your API IDs and Secret in your code): 
```
package main

import (
  "fmt"
  
  "github.com/darimm/striven"
)

func main() {
  s := striven.New("MYCLIENTID", "MYCLIENTSECRET")
  fmt.println(s.Token.AccessToken)
 
  resp, err := s.Employees.GetAll()
  
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
`Classes.GetAll()` implements https://api.striven.com/v1/classes  
- [ ] Contacts  
`Contacts.GetByID(contactID int)` implements https://api.striven.com/Help/Api/GET-v1-contacts-id
- [ ] Customers  
`Customers.GetByID(customerID int)` implements https://api.striven.com/v1/customers/{customerID}
`Customers.Contacts.GetByCustomerID(customerID int)` implements https://api.striven.com/v1/customers/{customerID}/contacts
`Customers.ContentGroups.GetByID(customerID int)` implements https://api.striven.com/v1/customers/{customerID}/hub/content-groups  
`Customers.ContentGroups.Document.Upload(remoteFileName string, localFilePath string, opts ...CustomersHubDocOption)` implements https://api.striven.com/Help/Api/POST-v1-customers-id-hub-content-groups-groupId-documents available options are `striven.SetClientID(ClientID int)`, `striven.SetGroupID(GroupID int)`, `striven.SetContentGroupName(GroupName string)`, `striven.IsOverwriteEnabled()`, and `striven.IsVisibleOnPortal()` this function is suitable for single file uploads, if your application needs concurrent uploads, create variables of the type CustomersHubDoc and call the Upload method as above.  
- [X] CustomList  
`CustomLists.GetAll()` implements https://api.striven.com/v1/custom-lists  
`CustomLists.ListItems.Get(ListID int)` implements https://api.striven.com/v1/custom-lists/{ListID}/list-items  
- [X] Employees  
`Employees.GetAll()` implements https://api.striven.com/v1/employees  
- [X] Formats  
`InvoiceFormats.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-invoice-formats  
- [X] GLCategories  
`GLCategories.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-glcategories  
- [X] Industries  
`Industries.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-industries  
- [X] InventoryLocations  
`InventoryLocations.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-inventory-locations  
- [X] ItemTypes  
`ItemTypes.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-item-types  
- [X] PaymentTerms  
`PaymentTerms.GetAll(excludeDiscounts bool)` implements https://api.striven.com/Help/Api/GET-v1-payment-terms_excludeDiscounts  
- [X] Pools  
`Pools.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-pools  
- [X] ReferralSources  
`ReferralSources.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-referral-sources  
- [X] SalesOrderTypes 
`SalesOrderTypes.GetAll(excludeContractManagedTypes bool)` implements https://api.striven.com/Help/Api/GET-v1-sales-order-types_excludeContractManagedTypes  
- [X] ShippingMethods  
`ShippingMethods.GetAll()` implements https://api.striven.com/Help/Api/GET-v1-shipping-methods  


APIs Not Implemented

- [ ] BillCredits  
- [ ] Bills  
- [ ] Categories  
- [ ] CreditMemos  
- [ ] CustomerAssets  
- [ ] GLAcconuts  
- [ ] Invoices  
- [ ] Items  
- [ ] JournalEntries  
- [ ] Opportunities  
- [ ] Payments  
- [ ] Purchase Orders  
- [ ] SalesOrders  
- [ ] Tasks  
- [ ] UserInfo  
- [ ] Vendors  
