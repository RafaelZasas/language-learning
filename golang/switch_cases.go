package main

import "fmt"

func average(x []float64) (avg float64){
  total := 0.0
  switch len(x) {
    // Go allows to use switch statements with variables instead of jsut const
    // Go will also only run the matching statement, no need for a break

    case 0:
      avg = 0
    default:
      for _, v := range x {
        total += v
      }
      avg = total /float(len(x))
  }
  return
}

func main() {
  x := []float64{2.15, 3.14, 42.0, 29.5}
  fmt.Println(average(x))  // 19.19749999999998
}
