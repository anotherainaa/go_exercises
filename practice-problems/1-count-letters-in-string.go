package main 

import "fmt"

/*
Count letters in string
In this kata, you've to count lowercase letters in a given string and return the letter count in a hash with 'letter' as key and count as 'value'. 
The key must be 'symbol' instead of string in Ruby and 'char' instead of string in Crystal.

input: string
output: hashMap with number of letter count 

approach 
- iterate through each character 
- if the letter exists in the hash add += 1
- else set to 1 

return hashof letter count

*/

func letterCount(letters string) map [string]int {
	letterCount := make(map[string]int)

	for _, v := range letters {
		_, exists := letterCount[string(v)]

		if (exists) {
			letterCount[string(v)] += 1
		} else {
			letterCount[string(v)] = 1
		}
	}
	return letterCount
}

func main() {
	fmt.Println(letterCount("arithmetics"))
	fmt.Println(letterCount("codewars"))
	fmt.Println(letterCount("codewars"))
}