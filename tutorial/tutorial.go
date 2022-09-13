package main

import (
	"fmt"
	"bufio"
	"strconv"
	// "sort"
	// "strings"
	// "math"
	// "reflect"
	// "strings"
)

// func main() {
// 	var name string
// 	name = "Ainaa"
// 	name = "Bill"

// 	fmt.Println("Konnichiwa World" + " " + name)
// 	fmt.Println(reflect.TypeOf(name))
// }

// func main() {
// 	var number = 2000.98
// 	text := "Hello, 世界" // assignment expression
// 	for index, characters := range text {
// 		fmt.Printf("%d: %s\n", index, string(characters))
// 	}
// 	fmt.Println("length of \"Hello, 世界\": ", len(text))
// 	fmt.Println(reflect.TypeOf(text))
// 	fmt.Println(strings.SplitAfter(text, " "))
// 	fmt.Printf("%T", number)
// }

// func main() {
// 	var stringSlice = []string{"This", "is", "a", "string", "slice"}
// 	fmt.Println(stringSlice)  // prints [This is a string slice]
// }

// func main() {
// 	var arr = [4]int{1, 2, 3, 4}
// 	var arrSlice = arr[1:3] // slice from index, not including the last one
// 	fmt.Println(arrSlice)
// }

// func main() {
// var odds = [8]int{3, 5, 7, 9, 11, 13, 15, 17}
// slice1 := odds[1:6] // [5 7 9 11 13]
// slice2 := slice1[2:4] // [9 11]
// fmt.Println(slice1)
// fmt.Println(slice2)
// }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("What is your birth year? >> ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10 , 64)
	fmt.Printf("You will be %d years old at the end of 2022", 2022 - input)
}

type Student struct {
	name string
	grades []int
	age int
}

func (s Student) getAge() int{
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32{
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

// How to make the pointer to the next node 
type Node struct {
	data string
	next_node *Node
}

type LinkedList struct {
	firstNode Node
}

func reverse(slice []int) []int{
	midpoint := len(slice) / 2

	// iterate only to midpoint
	// swap the first place with last
	// swap the second place with second last
	// swap middle - not necessary 
	for i := 0; i < midpoint; i++ {
		slice[i], slice[len(slice) - 1 - i] = slice[len(slice) - 1 - i], slice[i]
	}
	return slice
}


func remove(slice []string, index int) []string {
	return append(slice[:index], slice[(index + 1):]...)
}


// // a function can read a variable in the outer scope
// func main() {
// 	// ==== Creating a Strudent Struct
// 	// s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
// 	// fmt.Println(s1.getAge())
// 	// s1.setAge(8)
// 	// fmt.Println(s1.getAge())
// 	// fmt.Println(s1.getAverageGrade())
// 	// fmt.Printf("%t", is_prime(7))

	
// 	// fmt.Printf("%t", is_leap_year(2000))
// }
