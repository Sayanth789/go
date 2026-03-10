/* Go programs express errors with error values. An Error is any type that implements the simple built-in error interface:

type error interface {
	Error() string
}
The fmt package's error formatting tools work on any error type, and the errors.New function is a convenient way to create a basic error value with a static message. 

When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

// Atoi converts a stringified number to an interger
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then
// i was converted successfully


*/
package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}

	costSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return costSpouse + cost, nil
}


func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = .0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func test(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("========")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)
	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Total cost: $%.4f\n", totalCost)
}

func main() {
	test(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	test(
		"Thanks for joining us!",
		"Have a good day.",
	)
	test(
		"Thank you.",
		"Enjoy!",
	)
	test(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)
}