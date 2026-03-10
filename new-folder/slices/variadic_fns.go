/* 
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. 
This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func sum(nums ...int) int {
    // nums is just a slice
    for i := 0; i < len(nums); i++{
        num := nums[i]
    }
}

func main() {
    total := sum(1, 2, 3)
    fmt.Println(total)
    // prints "6"
}

The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with space delimiters and a newline at the end.

func Println(a ...interface{}) (n int, err error)

===============
Spread Operator 
===============
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}

*/

package main

import "fmt"

func sum(nums ...float64) float64 {
	// ?
	sum := 0.0
	for i := 0; i < len(nums); i++{
        num := nums[i]
        sum += num
    }
	return  sum
}


func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}


