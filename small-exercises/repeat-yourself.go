package main 

import "fmt"

func repeat(str string, num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(str)
	}
}

func main() {
	repeat("Hello", 3)
}