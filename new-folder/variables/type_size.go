/*
int  int8  int16  int32  int64 // whole numbers

uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers

float32 float64 // decimal numbers

complex64 complex128 // imaginary numbers (rare)


The size (8, 16, 32, 64, 128, etc) indicates how many bits 
in memory will be used to store the variable. The default int and uint types are just aliases that
 refer to their respective 32 or 64 bit sizes depending on the environment of the user.

*/
package main

import "fmt"

func main() {
	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer

	accountAgeInt := int(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}