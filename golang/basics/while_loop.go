// There is no native implementation for while loops,
// however you can write for loops to act as a while loop
package main

import "fmt"

func main(){
  count := 1
  for count < 5 {
    count += count
  }

  fmt.Println(count) // 8
}
