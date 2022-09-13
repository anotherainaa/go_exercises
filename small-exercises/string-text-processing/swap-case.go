package main

import (
	"fmt"
	"strings"
)

func swapCase(str string) string{
	var swapCased string

	for _, r := range str {
		letter := string(r)
		if strings.ToLower(letter) == letter {
			swapCased += strings.ToUpper(letter)
		} else if strings.ToUpper(letter) == letter {
			swapCased += strings.ToLower(letter)
		} else {
			swapCased += letter
		}
	}
	return swapCased
}

func main() {
	fmt.Println(swapCase("CamelCase"))
	fmt.Println(swapCase("Tonight on XYZ-TV"))
}