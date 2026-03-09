/*
Anonymous functions are true to form in that they have no name.
Anonymous functions are useful when defining a function that will only be used once or to create a quick closure.

*/

package main 

import "fmt"


func printReports(messages []string) {

	for _, message := range messages {
		printCostReport(func(m string) int {
			return len(m) * 2

		}, message) 
	}

}

func test(messages []string) {
	defer fmt.Println("===============================")
	printReports(messages)
}


func main() {

	test([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",

	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",

	})
}

func printCostReport(costCalcualtor func(string) int, message string) {
	cost := costCalcualtor(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}