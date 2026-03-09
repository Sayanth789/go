/*

A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

*/

package main 

import  "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type emailBill struct {
	costInPennies int 
}

func test(bills []emailBill) {
	defer fmt.Println("=========================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}


func main() {
	test([]emailBill {
		{45},
		{32},
		{43},
		{12},
		{32},
		{54},
	})

	test([]emailBill{
		{12},
		{12},
		{978},
		{12},
		{542},
	})

	test([]emailBill{
		{743},
		{13},
		{8},
	})


}