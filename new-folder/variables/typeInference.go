/*
To declare a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), 
the variable's type is inferred from the value on the right hand side.

When the right hadnd side of the declaration is typed, the new variable is of that same type.


var i int
j := i // j is also an int

*/

package main

import "fmt"

func main() {
	penniesPerText := 0.02
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)
}