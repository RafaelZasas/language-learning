package main

import "fmt"

func main(){
  a:= 42 // Initialize and assign to a single value
  b,c := true, 32.0 // Initialize and assign to multiple values
  d := "string"
  fmt.Println(a,b,c,d) // 42 true 32 string
}
