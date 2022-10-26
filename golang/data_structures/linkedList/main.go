package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

func (n node) String() string {
	return fmt.Sprintf("%d", n.value)
}

type linkedList struct {
	head *node
}

func (l *linkedList) add(value int) {
	newNode := new(node)
	newNode.value = value
	// This is redundant but leaving for reference
	// newNode.next = nil

	// Ensures that there is at least one node as Head
	if l.head == nil {
		l.head = newNode
	} else {
		iterator := l.head
		// this will traverse the linked list until we reach the end node
		for ; iterator.next != nil; iterator = iterator.next {
		}

		iterator.next = newNode
	}
}

func (l *linkedList) remove(value int) {
	var prev *node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {

			if l.head == current {
				l.head = current.next
			} else {
				prev.next = current.next
				return
			}
		}

		prev = current
	}
}

func (l linkedList) get(value int) *node {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(iterator.String())
	}
	return sb.String()
}

func main() {

	l := linkedList{}
	l.add(1)
	l.add(2)
	l.add(3)
	fmt.Println(l)
	fmt.Println(*(l.get(3)))
	l.remove(1)
	fmt.Println(l)
	fmt.Println((l.head))
}
