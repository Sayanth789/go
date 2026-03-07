// Go is a statically typed language, which means that you need to declare the type of a variable when you create it. You can use the var keyword to declare a variable, followed by the variable name and the type. For example:

// this is a syntatic sugar for declaring a variable and assigning a value to it in one line. The type of the variable is inferred from the value that you assign to it. For example:
// var conferenceName = "Go conference" becomes conferenceName := "Go conference"
package main
import "fmt"


func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	const remainingTickets = 50

    fmt.Printf("Welcome to,   %v  booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int 


	// ask user for their name 
	fmt.Scan(&userName)  // pointers ... oh.. lovely 
	userTickets = 2
	fmt.Printf("user %v booked booked %v tickets\n", userName, userTickets)





}

