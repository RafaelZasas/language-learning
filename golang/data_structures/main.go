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
		fmt.Println("running linked list")
		ll.Run()
	case "bst":
		fmt.Println("running binary search tree")
		bst.Run()

	default:
		fmt.Printf("Command \"%s\" not part of module.\n", args[0])
	}
}
