package main 

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
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

// convert the number into string and then use the function I've created befoe

func palindromic_number(n int) bool{
	return realPalindrome(strconv.Itoa(n))
}

func main() {
	fmt.Println(palindromic_number(34543))
	fmt.Println(palindromic_number(123210))
	fmt.Println(palindromic_number(22))
	fmt.Println(palindromic_number(5))
}