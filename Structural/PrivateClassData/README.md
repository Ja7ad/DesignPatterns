# Private Class Data Design Pattern
The private class data pattern secures the data within a class. This pattern encapsulates the
initialization of the class data. The write privileges of properties within the private class are
protected, and properties are set during construction. The private class pattern prints the
exposure of information by securing it in a class that retains the state. The encapsulation of
class data initialization is a scenario where this pattern is applicable.

`Account` is a class with account details and a customer name. AccountDetails is the
private attribute of Account , and CustomerName is the public attribute. JSON marshaling
of Account has CustomerName as a public property. AccountDetails is the package
property in Go (modeled as private class data)

```go
package main

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id string
	accountType string
}

type Account struct {
	details *AccountDetails
	CustomerName string
}


func (account *Account) setDetails(id string, accountType string)  {
	account.details = &AccountDetails{id, accountType}
}

// As shown in the following code, the Account class has the getId method, which returns
// the id private class attribute
func (account *Account) getId() string {
	return account.details.id
}

func (account *Account) getAccountType() string {
	return account.details.accountType
}

// The main method calls the Account initializer with CustomerName. The details of the
// account are set details with the setDetails method
func main() {
	var account *Account = &Account{CustomerName: "Javad"}
	account.setDetails("1", "admin")
	jsonAccount, _ :=json.Marshal(account)

	fmt.Println("Private Class Hidden", string(jsonAccount))
	fmt.Println("Account Id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
```

```
go run main.go

Private Class Hidden {"CustomerName":"Javad"}
Account Id 1
Account Type admin
```