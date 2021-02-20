# Adapter Design Pattern
The adapter pattern provides a wrapper with an interface required by the API client to link
incompatible types and act as a translator between the two types. The adapter uses the
interface of a class to be a class with another compatible interface. When requirements
change, there are scenarios where class functionality needs to be changed because of
incompatible interfaces.

The dependency inversion principle can be adhered to by using the adapter pattern, when a
class defines its own interface to the next level module interface implemented by an
adapter class. Delegation is the other principle used by the adapter pattern. Multiple
formats handling source-to-destination transformations are the scenarios where the adapter
pattern is applied.


The adapter pattern comprises the target, adaptee, adapter, and client:
* Target is the interface that the client calls and invokes methods on the adapter
  and adaptee.
* The client wants the incompatible interface implemented by the adapter.
* The adapter translates the incompatible interface of the adaptee into an interface
  that the client wants.
  

```go
package main

import "fmt"

type IProcess interface {
  process()
}

type Adapter struct {
  adaptee Adaptee
}

// Adaptor class method process
func (adapter Adapter) process()  {
  fmt.Println("Adapter Process")
  adapter.adaptee.convert()
}

type Adaptee struct {
  adapterType int
}

// Adaptee class method convert
func (adaptee Adaptee) convert()  {
  fmt.Printf("Adaptee convert method")
}

func main() {
  var processor IProcess = Adapter{}
  processor.process()
}
```

```
go run main.go

Adapter Process
Adaptee convert method
```