//Inside a function (even the main function), the := short assignment statement can be used in place of a var declaration. The := operator infers the type of the new variable based on the value.
/*
var empty string 

is the same as 

empty := ""

numCars := 10 // infered to be an integer 
temperature := 0.0 // temperature is inferred to be a floating point value because it has a decimal point

var isFunny = true // isFunny is inferred to be a boolean

*/

package main 
import "fmt"


func main() {

	congrats := "happy birthday..!"

	fmt.Println(congrats)
}
