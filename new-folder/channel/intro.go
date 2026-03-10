/* 

the 'go' keyword is used to start a new goroutine, which is a lightweight thread of execution. When you call a function with 'go', it runs concurrently with the calling code. In the context of the ping-pong example, using 'go' allows the pinger and ponger functions to run simultaneously, enabling them to send and receive messages through channels without blocking each other. 
It is helping to have concurrent execution of the program and ponger functions, as it allows them to interact with each other in real-time, simulating a ping-pong game where both players can send and receive messages without waiting for the other to finish. This concurrency is essential for the proper functioning of the ping-pong example, as it allows for a more dynamic and responsive interaction between the two functions.

*/

package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}


func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Stacy!")
	test("Hi there John!")
	test("Hey there Jane!")
}