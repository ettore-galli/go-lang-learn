package main

import "fmt"

type Package struct {
	label  string
	weight float64
}

func NewPackage(label string, weight float64) (Package, error) {
	return Package{label: label, weight: weight}, nil
}

func (c *Package) SetWeight(weight float64) {
	c.weight = weight
}

func PackageTest() {
	var c Package
	c = Package{label: "asda", weight: 3.14}
	c.SetWeight(2.71)
	fmt.Println(c)
	c, err := NewPackage("g", 9.81)
	fmt.Println(c, err)
}

func main() {

	PackageTest()
	c, err := NewPackage("g", 9.81)
	fmt.Println(c, err)

}
