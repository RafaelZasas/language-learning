package main

import (
	"fmt"
	"os"

	"github.com/rafaelzasas/language-learning/golang/data_structures/bst"
	"github.com/rafaelzasas/language-learning/golang/data_structures/ll"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 1 {
		fmt.Printf("Expected one command.\nRecieved: %v\n", args)
		os.Exit(1)
	}

	switch args[0] {
	case "ll":
		fmt.Printf("running linked list\n\n")
		ll.Run()
	case "bst":
		fmt.Printf("running binary search tree\n\n")
		bst.Run()

	default:
		fmt.Printf("Command \"%s\" not part of module.\n", args[0])
	}
}
