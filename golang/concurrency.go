// Goroutiines are a sequential control flow in a program.
// Multiple threads can make use of multiple CPU cores
// They are initiated using the "go" keyword.
// Goroutines have built-in channels which are used to share data between them
// in general send and recieve operations across a channel blocks the execution until the other side is ready.

// Below we are considering 5 goroutines that run in parallel.

package main

import (
	"fmt"
)

func main() {
  c := make(chan int) // Create a channel to pass ints
  // This channel will be common to all gorouines.

  for i := 0; i < 5; i++ {
    go cookingGopher(i, c)  // start a goroutine
  }

  for i := 0; i < 5; i++ {
    gopherID := <-c  // Recieve a value from a channel (c)
    fmt.Println("gopher", gopherID, "finished the dish")
  } // All the goroutines are finished at this point

}

// Notice the channel as an argument
func cookingGopher(id int, c chan int) {
    fmt.Println("Gopher", id, "started cooking")
    c <- id  // send a vlaue back to main
  }
