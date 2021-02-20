# Composite Design Pattern
A composite is a group of similar objects in a single object. Objects are stored in a tree form
to persist the whole hierarchy. The composite pattern is used to change a hierarchical
collection of objects. The composite pattern is modeled on a heterogeneous collection. New
types of objects can be added without changing the interface and the client code. You can
use the composite pattern, for example, for UI layouts on the web, for directory trees, and
for managing employees across departments. The pattern provides a mechanism to access
the individual objects and groups in a similar manner.

The composite pattern comprises the component interface, component class, composite,
and client:

* The component interface defines the default behavior of all objects and behaviors
  for accessing the components of the composite.
* The composite and component classes implement the component interface.
* The client interacts with the component interface to invoke methods in the
  composite.

Let's say there is an IComposite interface with the perform method and
BranchClass that implements IComposite and has the addLeaf, addBranch, and
perform methods. The Leaflet class implements IComposite with the perform
method. BranchClass has a one-to-many relationship with leafs and branches.
Iterating over the branch recursively, one can traverse the composite tree.

```go
package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

type Branch struct {
	leafs []Leaflet
	name string
	branches []Branch
}


func (leaf *Leaflet) perform()  {
	fmt.Println("Leaflet " + leaf.name)
}

// The perform method of the Branch class calls the perform method on branch and leafs, as seen
func (branch *Branch) perform()  {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) add(leaf Leaflet)  {
	branch.leafs = append(branch.leafs, leaf)
}

// The addBranch method of the Branch class adds a new branch
func (branch *Branch) addBranch(newbranch Branch)  {
	branch.branches = append(branch.branches, newbranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	var branch = &Branch{name: "My Branch 1"}
	var leaf1 = Leaflet{name: "My Leaf 1"}
	var leaf2 = Leaflet{name: "My Leaf 2"}
	var branch2 = Branch{name: "My Branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)
	branch.addBranch(branch2)
	branch.perform()
}
```

```
go run main.go

Branch: My Branch 1
Leaflet My Leaf 1
Leaflet My Leaf 2
Branch: My Branch 2
```