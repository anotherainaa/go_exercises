package main

import "fmt"

func palindrome(str string) bool{
	var reverse string
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse == str
}

func main() {
	fmt.Println(palindrome("madam"))
	fmt.Println(palindrome("Madam"))
	fmt.Println(palindrome("madam i'm adam"))
	fmt.Println(palindrome("356653"))
}