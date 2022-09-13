package main

import "fmt"

type dog struct {
	name string
	age int
}

// constructor

func newDog(name string, age int) dog {
	d := dog{name, age}
	return d
}

func main() {
	dog1 := newDog("Rex", 3)
	dog2 := newDog("Max", 10)

	fmt.Println(&dog1)
	fmt.Println(dog1.age)

	fmt.Println(&dog2)
	fmt.Println(dog2.age)

}