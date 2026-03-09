/* 
While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver. A receiver is a special parameter that 
syntactically goes before the name of the function.


type rect struct {
  width int
  height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
  return r.width * r.height
}

r := rect{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
// prints 50

A receiver is just a special kind of function parameter. Receivers are important because they will,
allow us to define interfaces that our structs (and other types) can implement.
*/


package main 
import "fmt"


type authenticationInfo struct {

	username string 
	password string
}


func (a authenticationInfo ) getBasicAuth() string {
	return "Authoriztion: Basic " + a.username + ":" + a.password
}

func test(a authenticationInfo) {
	fmt.Println(a.getBasicAuth())
	fmt.Println("====================================")
}

func main() {
	test(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	test(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	test(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}