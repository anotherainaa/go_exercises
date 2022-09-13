package main

import (
	"fmt"
	"regexp"
	"strings"
)

func palindrome(str string) bool{
	var reverse string
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse == str
}

func realPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := strings.ToLower(reg.ReplaceAllString(s, ""))
	return palindrome(processedString)
}

func main() {
	fmt.Println(realPalindrome("madam"))
	fmt.Println(realPalindrome("Madam"))
	fmt.Println(realPalindrome("Madam, I'm Adam"))
	fmt.Println(realPalindrome("356653"))
	fmt.Println(realPalindrome("356a653"))
	fmt.Println(realPalindrome("123ab321"))
}