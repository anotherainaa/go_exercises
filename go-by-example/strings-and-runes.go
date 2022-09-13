package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "おはよう"
	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		// fmt.Printf("%U %d", runeValue, idx)
	}

	// example of using utf8.decodeRuneInString to the same exact thing as above
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		
		examineRune(runeValue)
	}
}

// single quotes is for literal runes 
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'は'{
		fmt.Println("found ha")
	}
}