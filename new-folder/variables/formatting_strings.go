package main 
import "fmt"

func main() {
	const name = "Ben Akwnaman"
	const openRate = 50.4
	
	msg := fmt.Sprintf("Hi \"%s\", your open rate is %v percent", name, openRate)
	
	fmt.Println(msg)
}