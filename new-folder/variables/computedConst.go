/* Constants must be known at compile time. More often than not they will be declared with a static value:

You cannot declare a constant that can only be computed at run-time.

*/

package main 

import "fmt"

func main() {

	const secondsInMinutes = 60 
	const minutesInHour = 60 
	const secondsInHour = secondsInMinutes * minutesInHour


	fmt.Println("number of seconds in an hour:", secondsInHour)
}