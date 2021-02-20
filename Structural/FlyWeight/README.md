# FlyWeight Design Pattern
Flyweight is used to manage the state of an object with high variation. The pattern allows
us to share common parts of the object state among multiple objects, instead of each object
storing it. Variable object data is referred to as extrinsic state, and the rest of the object state
is intrinsic. Extrinsic data is passed to flyweight methods and will never be stored within it.
Flyweight pattern helps reduce the overall memory usage and the object initializing
overhead. The pattern helps create interclass relationships and lower memory to a
manageable level.

Flyweight objects are immutable. Value objects are a good example of the flyweight pattern.
Flyweight objects can be created in a single thread mode, ensuring one instance per value.
In a concurrent thread scenario, multiple instances are created. This is based on the equality
criterion of flyweight objects.

The participants of the **flyweight** pattern are the FlyWeight interface,
`ConcreteFlyWeight`, `FlyWeightFactory`, and the Client classes:

* The FlyWeight interface has a method through which flyweights can get and act
  on the extrinsic state.
* ConcreteFlyWeight implements the FlyWeight interface to represent
  flyweight objects.
* FlyweightFactory is used to create and manage flyweight objects. The client
  invokes FlyweightFactory to get a flyweight object. UnsharedFlyWeight can
  have a functionality that is not shared.
* Client classes

Let's say DataTransferObject is an interface with the getId method.
DataTransferObjectFactory creates a data transfer object through
getDataTransferObject by the DTO type. The DTO types are customer, employee,
manager, and address

```go
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
```

```
go run main.go

new DTO of dtoType: customer
Customer  1
new DTO of dtoType: employee
Employee  2
new DTO of dtoType: manager
Manager  3
new DTO of dtoType: address
Address  4
```