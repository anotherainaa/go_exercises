package main

import "fmt"

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func main() {
	i, j := 42, 2701
	
	// to see the address we can do this
	fmt.Println(&i, &j)

	p := &i

	// * is n operator that return the value p is pointing to
	// we refer to this as de-referencing
	fmt.Println(*p)

	// changing the value of i
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	a:= 4
	squareVal(a)
	fmt.Println(&a, a)

	squareAdd(&a)
	fmt.Println(&a, a)
}

func squareAdd(p *int) {
	*p *= *p
	fmt.Println(p, *p)
}