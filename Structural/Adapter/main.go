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