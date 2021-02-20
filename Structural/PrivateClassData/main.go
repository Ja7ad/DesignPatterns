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