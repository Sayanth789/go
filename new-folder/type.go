package main 

// Go has a lot of built in types, but you can also create your own custom types.
// Here we can't add different types of data together, but we can create a new type that is a combination of the two. This is called a struct, and it's a way to group together related data.


import "fmt"


func main() {

	var userName string = "Sayanth"
	var password string = "212112122112"

	// 
	fmt.Println("Authorization : Basic", userName+":"+password)
}