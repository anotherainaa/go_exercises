package main

import (
	"fmt"
	"sort"
	"strings"
	"math"
	// "bufio"
	// "os"
	// "strconv"
	// "reflect"
	// "strings"
)

func test(number int) {
	fmt.Println(number)
}

func truiey() bool {
	return true
}

func is_prime(number int) bool{
	for i := 2; i < number; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func is_leap_year(year int) bool {
	if year % 100 == 0 {
		return false
	} else if year % 400 == 0 {
		return true
	} else {
		return year % 4 == 0
	}
}

func has_duplicate_value(array []int) bool {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			if (i != j && array[i] == array[j]) {
				return true
			}
		}
	}
	return false
}

// note that this doesn't work!
// func has_duplicate_value_2(array []int) bool {
// 	var existing_numbers []int
// 	for i := 0; i < len(array); i++ {
// 		current_num := array[i]
// 		if existing_numbers[current_num] == 1 {
// 			return true
// 		} else {
// 			existing_numbers[current_num] = 1
// 		}
// 	}
// 	return false
// }

func array_sum(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func bubble_sort(list []int) []int{
	unsorted_until_index := len(list) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < unsorted_until_index; i++ {
			if list[i] > list[i + 1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		unsorted_until_index -= 1
	}
	return list
}

func insertion_sort(array []int) []int{
	for index := 1; index < len(array); index++ {
		temp_value := array[index]
		position := index - 1

		for position >= 0 {
			if array[position] > temp_value {
				array[position + 1] = array[position]
				position = position - 1
			} else {
				break
			}
		}
		array[position + 1] = temp_value
	}

	return array
}

func intersection(array1 []int, array2 []int) []int{
	result := []int{}
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				result = append(result, array1[i])
				break
			}
		}
	}
	return result
}

// must have complexity of O(N)
// use hash as a lookup table? 
func intersection2(array1 []int, array2 []int) []int {
	var results []int
	mp := make(map[int]bool)
	// 1:true
	// 2:false
	// map[int]bool{
	// 	1:true,
	// 	2:true
	// }

	// determine the larger array 
	// keep the values of the larger array in a hash table
	// iterate through the shorter array
	// check if it exists in the hash table 
	// if it does push it into the results array 

	var longerArray []int
	var shorterArray []int

	if (len(array1) >= len(array2)) {
		longerArray = array1
		shorterArray = array2
	} else {
		longerArray = array2
		shorterArray = array1
	}

	for i := 0; i < len(longerArray); i++ {
		currentEl := longerArray[i]
		mp[currentEl] = true
	}

	for j := 0; j < len(shorterArray); j++ {
		currentEl := shorterArray[j]
		if mp[currentEl] {
			results = append(results, currentEl)
		}
	}

	return results
}

func has_duplicate_char(array []string) string{
	dict := make(map[string]int)

	for i := 0; i < len(array); i++ {
		currentChar := array[i]
		dict[currentChar] += 1
	}

	for key, element := range dict {
		if element >= 2 {
			return key
		}
	}

	return "no duplicates"
}

func wordBuilder(slice []string) []string{
	collection := []string{};

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if i != j {
				collection = append(collection, slice[i] + slice[j])
			} 
		}
	}
	return collection
}

func find_missing_alphabet(words string) string{
	alphabet := make(map[string]bool)

	var ch byte
	for ch = 'a'; ch <= 'z'; ch++ {
		alphabet[string(ch)] = true
	}

	characters := make(map[string]int)
	wordsArr := strings.Split(words, "")

	for i := 0; i < len(wordsArr); i++ {
		currentChar := wordsArr[i]

		if currentChar == " " {
			continue
		} else {
			characters[currentChar] += 1
		}
	}

	for key := range alphabet {
		if characters[key] == 0 {
			return key
		}
	}

	return "no missing characters"
}

func non_duplicated_character(word string) string{
	characterCount := make(map[string]int)

	wordChars := strings.Split(word, "")

	for i := 0; i < len(wordChars); i++ {
		char := wordChars[i]
		characterCount[char] += 1
	}

	for i := 0; i < len(wordChars); i++ {
		char := wordChars[i]
		if characterCount[char] == 1 {
			return char
		}
	}
	return "no non duplicated characters"
}

func sum_recursion(array []int) int{
	if len(array) == 1 {
		return array[0]
	}

	return array[0] + sum_recursion(array[1:(len(array))])
}

//
func reverse_string_recrusion(str string) string{
	arr := strings.Split(str, "")

	if len(arr) == 1 {
		return arr[0]
	}

	return reverse_string_recrusion(strings.Join(arr[1:(len(arr))], "")) + arr[0]
}

func max(array []int) int{
	fmt.Println("recursion")

	if len(array) == 1 {
		return array[0]
	}

	max_of_remainder := max(array[1:(len(array))])

	if array[0] > max_of_remainder {
		return array[0]
	} else {
		return max_of_remainder
	}
}

func fib(n int, memo map[int]int) int{
	if n == 0 || n == 1 {
		return n
	}

	_, exists := memo[n]
	if !exists {
		memo[n] = fib(n - 2, memo) + fib(n - 1, memo)
	}

	return memo[n]
}


func add_until_100(array []int) int{
	if len(array) == 0 {
		return 0
	}

	sum_of_remaining_numbers := add_until_100(array[1:(len(array))])

	if array[0] + sum_of_remaining_numbers > 100 {
		return sum_of_remaining_numbers
	} else {
		return array[0] + sum_of_remaining_numbers
	}
}

func golomb(n int, memo map[int]int) int{
	if n == 1 {
		return 1 
	}

	_, exists := memo[n]
	if !exists {
		memo[n] = golomb(n - golomb(golomb(n - 1, memo), memo), memo)
	}

	return 1 + memo[n]
	// return 1 + golomb(n - golomb(golomb(n - 1)))
}

func unique_paths(rows int, columns int, memo map[[2]int]int) int{
	if rows == 1 || columns == 1 {
		return 1
	}

	_, exists := memo[[2]int{rows, columns}]
	if !exists {
		memo[[2]int{rows, columns}] = unique_paths(rows - 1, columns, memo) + unique_paths(rows, columns - 1, memo)
	}

	return memo[[2]int{rows, columns}]
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

func makeUpperCase(slice []string) []string{
	var newSlice []string
	for i := 0; i < len(slice); i++ {
		fmt.Println(i)
		newSlice = append(newSlice, strings.ToUpper(slice[i]))
	}

	return newSlice
}

func makeUpperCase2(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		slice[i] = strings.ToUpper(slice[i])
	}
	return slice
}

func hasDuplicateValue3(slice []int) bool{
	sort.Ints(slice)

	for i := 0; i < len(slice) - 1; i++ {
		if slice[i] == slice[i + 1] {
			return true
		}
	}
	return false
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

var authors [5]Author
// var output_map = map[string]interface{}{
// 	"author_id": c.Int("author_id"),
// 	"name": c.String("name"),
// }

type Author struct {
	author_id int
	name string
}

type Book struct {

}

var author_hash_table = map[int]string {
	1: "Virginia Woolf",
	2: "Leo Tolstoy",
	3: "Dr. Seuss",
	4: "J. K. Rowling",
	5: "Mark Twain",
}

var author_1 = Author{1, "Virginia Woolf"} 
var author_2 = Author{2, "Leo Tolstoy"} 
var author_3 = Author{3, "Dr. Seuss"} 
var author_4 = Author{4, "J. K. Rowling"} 
var author_5 = Author{5, "Mark Twain"} 
// {"author_id" => 2, "name" => "Leo Tolstoy"}, 
// {"author_id" => 3, "name" => "Dr. Seuss"}, 
// {"author_id" => 4, "name" => "J. K. Rowling"}, 
// {"author_id" => 5, "name" => "Mark Twain"}

func connect_books_with_authors(books []Book, authors map[int]string) {
	fmt.Println(authors)
}

func twoSum(slice []int) bool{
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if (i != j && slice[i] + slice[j] == 10) {
				return true
			}
		}
	}
	return false
}

func twoSum2(slice []int) bool {
	hashTable := make(map[int]bool)

	for i := 0; i < len(slice); i++ {
		if hashTable[10 - slice[i]] {
			return true
		}
		hashTable[slice[i]] = true
	}
	return false
}

func game_winner(number_of_coins int, current_player string) string {
	var next_player string = ""

	if number_of_coins <= 0 {
		return current_player
	}

	if current_player == "you" {
		next_player = "them"
	} else if current_player == "them" {
		next_player = "you"
	}

	if game_winner(number_of_coins - 1, next_player) == current_player || game_winner(number_of_coins - 2, next_player) == current_player {
		return current_player
	} else {
		return next_player
	}
}

func game_winner2(number_of_coins int) string{
	if (number_of_coins - 1) % 3 == 0 {
		return "them"
	} else {
		return "you"
	}
}

func sum_swap(slice1 []int, slice2 []int) []int {
	hash_table := make(map[int]int)
	sum_1 := 0
	sum_2 := 0

	for i := 0; i < len(slice1); i++ {
		sum_1 += slice1[i]
		hash_table[slice1[i]] = i
	}

	for i := 0; i < len(slice2); i++ {
		sum_2 += slice2[i]
	}

	shift_amount := (sum_1 - sum_2) / 2


	for i := 0; i < len(slice2); i++ {
		_, exists := hash_table[(slice2[i] + shift_amount)]
		if exists {
			return []int{hash_table[slice2[i] + shift_amount], i}
		}
	}
	return []int{}
}

func max_sum(slice []int) int{
	current_sum := 0
	greatest_sum := 0

	for i := 0; i < len(slice); i++ {
		if current_sum + slice[i] < 0 {
			current_sum = 0
		} else {
			current_sum += slice[i]
		}

		if current_sum > greatest_sum {
			greatest_sum = current_sum
		}
	}
	return greatest_sum
}

func increasing_triplet(slice []int) bool{
	lowest_price := slice[0]
	middle_price := int(math.Inf(1))

	for i:= 0; i < len(slice); i++ {
		price := slice[i]
		if price <= lowest_price {
			lowest_price = price
		} else if price <= middle_price {
			middle_price = price
		} else {
			return true
		}
	}
	return false
}

func areAnagrams(string1 string, string2 string) bool {
	second_string_array := strings.Split(string2, "")

	for i := 0; i < len(string1); i++ {
		if len(second_string_array) == 0 {
			return false
		}

		for j := 0; j < len(second_string_array); j++ {
			fmt.Println(second_string_array)
			if string(string1[i]) == second_string_array[j] {
				second_string_array = remove(second_string_array, j)
				break
			}
		}

	}

	return len(second_string_array) == 0
}

func are_anagrams2(string1 string, string2 string) bool{
	first_word_map := make(map[string]int)
	second_word_map := make(map[string]int)

	for _, char := range string1 {
		first_word_map[string(char)] += 1 
	}

	for _, char2 := range string2 {
		second_word_map[string(char2)] += 1
	}

	for key, value := range first_word_map {
		v, exists := second_word_map[key]
		if !exists || value != v {
			return false
		}
	}
	return true
}

func group_array(slice []string) []string {
	hash_table := make(map[string]int)
	new_slice := []string{}

	for _, el := range slice {
		hash_table[el] += 1
	}

	for key, value := range hash_table {
		for i := 0; i < value; i++ {
			new_slice = append(new_slice, key)
		}
	}
	return new_slice
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[(index + 1):]...)
}

// time complexity O(N + M)
// space complexity O(N)
func missing_integer(slice []int) int {
	hash := make(map[int] bool)

	for _, elem := range slice {
		hash[elem] = true
	}

	for index := range slice {
		_, exists:= hash[slice[index] + 1]

		if !exists {
			current_missing := slice[index] + 1
			
			_, exists:= hash[current_missing + 1]
			if exists {
				return current_missing
			}
		}
	}
	return -1
}

// time complexity O(N + M)
// space complexity O(N)
func missing_integer2(slice []int, integer int) int {
	hash := make(map[int] bool)

	for _, elem := range slice {
		hash[elem] = true
	}

	for i := 0; i <= integer; i++ {
		_, exists:= hash[i]

		if !exists {
			return i
		}
	}
	return -1
}

// Time complexity O(N)
// Space complexity O(1) ? 
func greatest_profit(slice []int) int {
	lowest_price := slice[0]
	greatest_profit := 0

	for i := 0; i < len(slice) - 1; i++ {
		// check for lowest price
		if (slice[i] < lowest_price) {
			lowest_price = slice[i]
		}

		if (slice[i] - lowest_price > greatest_profit) {
			greatest_profit = slice[i] - lowest_price
		}
	}
	return greatest_profit
}

// a function can read a variable in the outer scope
func main() {
	// ==== Writing examples for DS&A book
	// result := intersection2([]int{55, 45, 35, 25, 15, 10}, []int{45, 3, 25})
	// result := has_duplicate_char([]string{"a", "b", "c", "d", "c", "e", "f"});
	// result := find_missing_alphabet("the quick brown box jumps over a lazy dog")
	// result := non_duplicated_character("minimum")
	// result := sum_recursion([]int{1, 2, 3, 4, 5})
	// result := reverse_string_recrusion("abcd")
	// result := max([]int{1, 2, 3, 4})
	// result := fib(6, make(map[int]int))
	// result := add_until_100([]int{1, 2, 3, 4, 100, 5})
	// arr := map[[2]int{0, 0}]int
	// hash := map[[2]int]int{
	// 	{0, 0}: 0,
	// }
	// result := unique_paths(3, 7, hash)

	// ==== Linked List
	// node_4 := Node{"time", nil}
	// node_3 := Node{"a", &node_4}
	// node_2 := Node{"upon", &node_3}
	// node_1 := Node{"once", &node_2}
	// fmt.Println(node_1)
	// fmt.Println(node_2)
	// fmt.Println(node_3)
	// fmt.Println(node_4)
	// fmt.Println(reflect.TypeOf(&node_2))

	// result := makeUpperCase2([]string{"tuvi", "leah", "shaya", "rami"})

	// result := hasDuplicateValue3([]int {3, 4, 2, 1, 6, 7, 8})
	// result := wordBuilder([]string{"a", "b", "c", "d", "e", "f", "g"})

	// ==== Book author problem chapter 20
	// authors[0] = author_1
	// authors[1] = author_2
	// authors[2] = author_3
	// authors[3] = author_4
	// authors[4] = author_5
	// result := reverse([]int {7, 8, 3, 4})

	// ==== Two sum Problem
	// result := twoSum2([]int{2, 0, 4, 5, 3, 9})

	// === The coin problem
	// result := game_winner(10, "you")
	// result2 := game_winner2(10)
	// fmt.Println(result)
	// fmt.Println(result2)

	// === Sum swap
	// result := sum_swap([]int{5, 3, 2, 9, 1}, []int{1, 12, 5})
	// fmt.Println(result)

	// === greatest sum
	// result := max_sum([]int{3, -4, 4, -3, 5, -9})

	// fmt.Println(result)

	// === increasing prices 
	// result := increasing_triplet([]int{5, 2, 8, 4, 3, 7})
	// fmt.Println(result)

	// === anagram
	// result := are_anagrams2("madam", "damam")
	// fmt.Println(result)

	// === group slice
	// result := group_array([]string{"a", "b", "c", "c", "b", "a"})
	// fmt.Println(result)

	// === exercise 20.2
	// result := missing_integer2([]int{8, 2, 3, 9, 4, 7, 5, 0, 6}, 9)
	// result2 := missing_integer2([]int{2, 3, 0, 6, 1, 5}, 5)
	// fmt.Println(result)
	// fmt.Println(result2)

	// === greatest profit
	result := greatest_profit([]int{10, 7, 5, 8, 11, 2, 6})
	fmt.Println(result)

	// ==== Creating a Strudent Struct
	// s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	// fmt.Println(s1.getAge())
	// s1.setAge(8)
	// fmt.Println(s1.getAge())
	// fmt.Println(s1.getAverageGrade())
	// fmt.Printf("%t", is_prime(7))

	
	// fmt.Printf("%t", is_leap_year(2000))
}
