/*

There could be many reasons. For example, maybe a function called getCircle returns the center point and the radius, but you really only need the 
radius for your calculation. In that case, you would ignore the center point variable.

This is crucial to understand because the Go compiler will throw an error if you have unused variable declarations in your code, so you need to ignore anything you don't intend to use.

func getPoint() (x int, y int) {
  return 3, 4
}

// ignore y value
x, _ := getPoint()

*/
package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}


func getNames() (string, string) {
	return "John", "Doe"
}