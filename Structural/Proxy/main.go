package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct {}

type VirtualProxy struct {
	realobject *RealObject
}

func (realobject *RealObject) performAction()  {
	fmt.Println("RealObject performAction...")
}

func (proxy *VirtualProxy) performAction()  {
	if proxy.realobject == nil {
		proxy.realobject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction...")
	proxy.realobject.performAction()
}

func main() {
	var object = VirtualProxy{}
	object.performAction()
}