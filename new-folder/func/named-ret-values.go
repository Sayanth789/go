/* 
Return values may be given names, and if they are, then they are treated the same as if they were new variables defined at the top of the function.

Named return values are best thought of as a way to document the purpose of the returned values.

func getCoords() (x, y int){
  // x and y are initialized with zero values

  return // automatically returns x and y
}

--------------
is the same as
--------------
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}



*/
package main

import (
	"fmt"
)

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		fmt.Println("congrats you're already an adult")
		yearsUntilAdult = 0
		
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		fmt.Println("You can already drink")
	    yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		fmt.Println("You can already rent a car")
		yearsUntilCarRental = 0

	}
	return
}


func test(age int) {
	fmt.Println("Age:", age)
	yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(age)
	fmt.Println("You are an adult in", yearsUntilAdult, "years")
	fmt.Println("You can drink in", yearsUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRental, "years")
	fmt.Println("====================================")
}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}