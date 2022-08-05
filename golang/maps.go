// Notice how the outputs of the map are not in the order that they were programmed in
package main

import "fmt"

func main(){
  // define a map containing the release year of several languages

  firstReleases := map[string]int{
	  "C": 1972,
	  "C++": 1985,
	  "Java": 1996,
	  "Python": 1991,
	  "Javascript": 1996,
	  "Go": 2012,
  }

  // loop through each entry and output the name and release year

  for k,v := range firstReleases {
    fmt.Printf("%s was first released in %d\n", k, v)
  }
}
