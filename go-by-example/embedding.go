package main

import "fmt"

// creating a base struct
type base struct {
	num int
}

// a method for base called describe 
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container is a struct that embeds a base - embedding looks like a field without a name
type container struct {
	base // the base is the struct type?
	str string
}


func main() {
	co := container{
		base:base{
			num: 1,
		},
		str: "some name",
	}

	// directly accessing field name num on base
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// here, spelling out the full path of the embedded type name
	fmt.Println("also num:", co.base.num)

	// container embeds base, therefore methods of base also become methods of container
	// invoking method describe which is method for base directly on co (container)
	fmt.Println("describe:", co.describe())
	
	// container now implements describer interface because it embeds base
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}