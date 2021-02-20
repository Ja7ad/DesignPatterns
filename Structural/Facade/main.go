package main

import "fmt"

type Account struct {
	id string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

// The account class has the deleteById method, which is used to delete an account with a given ID
func (account *Account) deleteById(id string)  {
	fmt.Println("delete account by id")
}

type Customer struct {
	id int
	name string
}

// the customer class has a method that creates a new customer with name
func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

type Transaction struct {
	id string
	amount float32
	srcAccountId string
	destAccountId string
}

// the transaction class has the create method for creating a transaction
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFacade struct {
	account *Account
	customer *Customer
	transaction *Transaction
}

func NewBranchManager() *BranchManagerFacade  {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// BranchManagerFacade has the createCustomerAccount method, which calls
// the create method on the customer class instance
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId,destAccountId,amount)
	return transaction
}

func main() {
	var facade = NewBranchManager()
	var customer *Customer
	var account *Account
	var transaction = facade.createTransaction("1122", "1133", 1000)

	customer, account = facade.createCustomerAccount("Javad", "Saving")

	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	fmt.Println(transaction.amount)

}