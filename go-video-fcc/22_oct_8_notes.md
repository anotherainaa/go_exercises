Interesting discovery about Golang

- How Go handles variable shadowing

```go
package main 

import "fmt"

var a = 20 // This is a global variable (only visible within the current file)

func main() {
  fmt.Println(a); // this will output 20 
  var a = 13
  fmt.Println(a); // here global variable a is shadowed while local variable a on line 14 is referenced. 
  a := 1 // However this will not work because we cannot re-declare a variable that has been declared in the block scope  
}
```

Something that I haven't seen before in JS nor Ruby is the fact that the function can refer to global variable up until the point where a new variable is initialized. 
Unlike JS, a is not hoisted, therefore an error is not thrown. 

Re-call notes 

- variables that start with Uppercase letter is available everywhere including other files 
- how type conversion works roughly 
- the difference between rune and bytes
  - bytes is utf8? 
  - rune is utf32. Rune is an integer representing the character. 
- byte is useful for some things, like when exporting the Go application to be used by something else also using bytes 
