package main

import "fmt"

func shortLongShort(s1, s2 string) string{
	if len(s1) < len(s2) {
		return s1 + s2 + s1
	} else {
		return s2 + s1 + s2
	}
}

func main() {
	fmt.Println(shortLongShort("abc", "defgh"))
	fmt.Println(shortLongShort("abcde", "fgh"))
	fmt.Println(shortLongShort("", "xyz"))
}