package main 
import "fmt"


func concat(s1 string, s2 string) string {
	return s1 + s2
}


func main() {
	test("Lane, ", " Happy birthday")
	test("Davinci, ", "desgined robots???")
	test("Go", "is not bad")
}


func test(s1 string, s2 string ) {
	fmt.Println(concat(s1, s2))
}
