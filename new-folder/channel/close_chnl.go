/* 
Channels can be explicitly closed by a sender:

ch := make(chan int)

// do some stuff with the channel

close(ch)
Similar to the ok value when accessing data in a map, receivers can 
check the ok value when receiving from a channel to test if a channel was closed.

v, ok := <-ch
ok is false if the channel is empty and closed.


*/

package main

import (
	"fmt"
	"time"
)


func countReports(numSentCh chan int) int {
	total := 0 

	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}


func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func main() {
	test(3)
	test(4)
	test(5)
	test(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}