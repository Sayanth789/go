/* 
The default case in a select statement executes immediately if no other 
channel has a value ready. A default case stops the select statement from blocking.

select {
  case v := <-ch:
    // use v
  default:
    // receiving from ch would block
    // so do something else
}

A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:

func main(){
  ch := make(chan int)
  readCh(ch)
}

func readCh(ch <-chan int) {
  // ch can only be read from
  // in this function
}

Write-only channels 

The same goes for write-only channels, but the arrow's position moves.
func writeCh(ch chan<- int) {
  // ch can only be written to
  // in this function
}
 
*/

package main 
import (
	"fmt"
	"time"
)


func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	// ?
	for {
		select {
			case <- snapshotTicker:
				takeSnapshot()
			case <- saveAfter:
				saveSnapshot()
				return
			default:
				waitForData()	
				time.Sleep(500 * time.Millisecond)
		
		}	
	}

}


func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}


func saveSnapshot() {
	fmt.Println("All backups saved!")
}


func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}



func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}
