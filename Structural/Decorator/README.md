# Decorator Design Pattern
In a scenario where class responsibilities are removed or added, the decorator pattern is
applied. The decorator pattern helps with subclassing when modifying functionality,
instead of static inheritance. An object can have multiple decorators and run-time
decorators. The single responsibility principle can be achieved using a
decorator. The decorator can be applied to window components and graphical object
modeling. The decorator pattern helps with modifying existing instance attributes and
adding new methods at run-time.

The decorator pattern participants are the component interface, the concrete component
class, and the decorator class:

* The concrete component implements the component interface.
* The decorator class implements the component interface and provides
  additional functionality in the same method or additional methods. The
  decorator base can be a participant representing the base class for all decorators.

Let 's say IProcess is an interface with the process method. ProcessClass implements
an interface with the process method. ProcessDecorator implements the process
interface and has an instance of ProcessClass. ProcessDecorator can add more
functionality than ProcessClass

```go
package main

import "fmt"

type IProcess interface {
	process()
}

type ProcessClass struct {}

func (process *ProcessClass) process()  {
	fmt.Println("ProcessClass process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

// the ProcessDecorator class process method invokes the process
// method on the decorator instance of ProcessClass

func (decorator *ProcessDecorator) process()  {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Println("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}

	decorator.processInstance = process
	decorator.process()
}
```

```
go run main.go

ProcessDecorator process and 
ProcessClass process
```