package main 

import "fmt"

/*
You are given array of integers, your task will be to count all pairs in that array and return their count.

Notes:
- Array can be empty or contain only one value; in this case return 0
- If there are more pairs of a certain number, count each pair only once. E.g.: for [0, 0, 0, 0] the return value is 2 (= 2 pairs of 0s)
- Random tests: maximum array length is 1000, range of values in array is between 0 and 1000

Examples
[1, 2, 5, 6, 5, 2]  -->  2
...because there are 2 pairs: 2 and 5

[1, 2, 2, 20, 6, 20, 2, 6, 2]  -->  4
...because there are 4 pairs: 2, 20, 6 and 2 (again)

p pairs([1, 2, 5, 6, 5, 2]) == 2
p pairs([1, 2, 2, 20, 6, 20, 2, 6, 2]) == 4
p pairs([0, 0, 0, 0, 0, 0, 0]) == 3 
p pairs([1000, 1000]) == 1
p pairs([]) == 0
p pairs([54]) == 0

Input
- array of values
- number of pairs inside of the array 

Approach
- iterate through the array of numbers 
- if the number exists in hashmap 
	- pair += 1
	- delete the number from hashmap 
- save the number in hashmap 
- return the pair number 

algorithm
pair = 0 
pairCheck = make(map [int]bool)
start loop
	if the number exists in the pairCheck 
		- pair += 1
		- delete number
	- set number to true in pairCheck
end loop
return pair

*/
func pairs(nums []int) int {
	pair := 0
	pairCheck := make(map[int]bool)

	for _, v := range nums {
		_, exists := pairCheck[v]

		if (exists) {
			pair += 1
			delete(pairCheck, int(v))
			continue
		}

		pairCheck[v] = true
	}
	return pair
}

func main() {
	fmt.Println(pairs([]int{1, 2, 5, 6, 5, 2}))
	fmt.Println(pairs([]int{1, 2, 2, 20, 6, 20, 2, 6, 2}))
	fmt.Println(pairs([]int{0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(pairs([]int{1000, 1000}))
	fmt.Println(pairs([]int{}))
	fmt.Println(pairs([]int{54}))
}