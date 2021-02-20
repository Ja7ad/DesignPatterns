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