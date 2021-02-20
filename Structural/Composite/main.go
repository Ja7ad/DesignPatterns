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