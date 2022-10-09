package main 

import (
	"fmt"
	"math"
)

/*

Longest vowel chain
The vowel substrings in the word codewarriors are o,e,a,io. 
The longest of these has a length of 2. 
Given a lowercase string that has alphabetic characters only and no spaces, return the length of the longest vowel substring. Vowels are any of aeiou.

vowel: aeiou
find the longest vowel chain 

approach 
- create constant of vowels
- iterate through each character 
- if the character is a vowel, increase vowel chain by 1
- if the character is not a vowel, reset vowel chain to 0 
- if we reach the end, return the chain value

*/


func vowelChain(str string) int {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	chain := 0
	max := 0

	for _, v := range str {
		_, exists := vowels[string(v)]
		if (exists) {
			chain += 1
		} else {
			chain = 0
		}

		max = int(math.Max(float64(chain), float64(max)))
	}
	return max
}

func main() {
	fmt.Println(vowelChain("codeawarrios"))
	fmt.Println(vowelChain("suoidea"))
	fmt.Println(vowelChain("iuuvgheaae"))
	fmt.Println(vowelChain("ultrarevolutionariees"))
	fmt.Println(vowelChain("strengthlessnesses"))
	fmt.Println(vowelChain("cuboideonavicuare"))
	fmt.Println(vowelChain("chrononhotonthuooaos"))
	fmt.Println(vowelChain("iiihoovaeaaaoougjyaw"))
}