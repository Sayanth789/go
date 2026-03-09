/*

Constants are declared like variables but use the const keyword. Constants can't use the := short declaration syntax.

Constants can be character, string, boolean, or numeric values. They can'y be more complex
types lile slices, maps and structs. 

As the name implies, the value of a constant can't be changed after it has been declared.
*/

package main 


import "fmt"


func main() {
	const  premiumPlanName = "Premium Plan"
	const basicPlanName = "Basic Plan"
	
	
	fmt.Println("plan:", premiumPlanName)
	fmt.Println("plan:", basicPlanName)
}