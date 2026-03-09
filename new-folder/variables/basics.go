/*

go's basic variable types are:

bool 

string 

int int8 int16 int32 int64 

uint uint8 uint16 uint32 uint64 

byte  // alias for unit8 

rune // alias for unit32

float32 float64

complex64 complex128


*/
package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string
	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
