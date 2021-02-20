package main

import "fmt"

type DataTransferObject interface {
	getId() string
}

type DataTransferObjectFactory struct {
	pool map[string] DataTransferObject
}

type Customer struct {
	id string // sequence generator
	name string
	ssn string
}

type Employee struct {
	id string
	name string
}

type Manager struct {
	id string
	name string
	dept string
}

type Address struct {
	id string
	streetLine1 string
	streetLine2 string
	state string
	city string
}

// Class Method getId
func (customer Customer) getId() string {
	return customer.id
}

func (employee Employee) getId() string {
	return employee.id
}

func (manager Manager) getId() string {
	return manager.id
}

func (address Address) getId() string {
	return address.id
}

// DataTransferObjectFactory class method getDataTransferObject
func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}

	var customer = factory.getDataTransferObject("customer")
	fmt.Println("Customer ", customer.getId())

	var employee = factory.getDataTransferObject("employee")
	fmt.Println("Employee ", employee.getId())

	var manager = factory.getDataTransferObject("manager")
	fmt.Println("Manager ", manager.getId())

	var address = factory.getDataTransferObject("address")
	fmt.Println("Address ", address.getId())
}