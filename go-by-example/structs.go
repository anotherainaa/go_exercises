package main

import "fmt"

// kind of like the blueprint, "class"
type person struct {
	name string
	age int
}

// newPerson is kind of like a constructor method
// it returns a pointer> 

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// a simple example of initializing a struct
	fmt.Println(person{"Bob", 20})

	// naming the fields when intiaializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// & prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// access struct fields using dot notation
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// use dots with struct pointers - the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp)
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}