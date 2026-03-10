/*

: : : Maps : : : 
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.

the zero value for a map is nil.

We can create a map by using a literal or by using the make() function:

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

ages = map[string]int{
  "John": 37,
  "Mary": 21,
}

The len() function works on a map, it returns the total number of key/value pairs.
ages = map[string]int{
  "John": 37,
  "Mary": 21,
}
fmt.Println(len(ages)) // 2
*/

package main

import (
	"fmt"
	"errors"
)


func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// if len(names) != len(phoneNumbers) {
	// 	return nil, fmt.Errorf("length of names and phoneNumbers must be the same")
	// }
	// users := make(map[string]user)
	// for i := 0; i < len(names); i++ {
	// 	users[names[i]] = user{
	// 		name:        names[i],
	// 		phoneNumber: phoneNumbers[i],
	// 	}
	// }
	// return users, nil
	users := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}
	for i := range names {
		name := names[i]
		users[name] = user {
			name: name, 
			phoneNumber: phoneNumbers[i],
		}
	}
	return users, nil


}


type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}