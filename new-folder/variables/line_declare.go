// We're able to declare multiple variables on the same line 

// mileage , company := 90454, BMW

// mileage := 90454
// company := BMW 

package main 

import "fmt"


func main() {

	averageOpenRate, displayMessage := 123,"Hello beautiful..."

	fmt.Println(averageOpenRate, displayMessage)
}