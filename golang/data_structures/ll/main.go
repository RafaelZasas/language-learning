package ll

import (
	"fmt"
	"strings"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

func (n node[T]) String() string {
	return fmt.Sprintf("%v ", n.value)
}

type linkedList[T comparable] struct {
	head *node[T]
}

func (l *linkedList[T]) add(value T) {
	newNode := new(node[T])
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

func (l *linkedList[T]) remove(value T) {
	var prev *node[T]

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

func (l linkedList[T]) get(value T) *node[T] {
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList[T]) String() string {
	sb := strings.Builder{}

	for iterator := l.head; iterator != nil; iterator = iterator.next {
		sb.WriteString(iterator.String())
	}
	return sb.String()
}

func Run() {
	l := linkedList[string]{}
	l.add("B")
	l.add("A")
	l.add("C")
	fmt.Printf("Creating string linked list: %v\n", l)

	fmt.Printf("Getting \"A\" from linked list: %v\n", *l.get("A"))
	fmt.Println("Removing \"A\" from linked list")
	l.remove("A")
	fmt.Println(l)
	fmt.Printf("Head of linked list: %v\n\n", l.head)

	t := linkedList[int]{}
	t.add(2)
	t.add(4)
	t.add(1)
	fmt.Printf("Creating an integer linked list: %v\n", t)
	fmt.Printf("Getting 4 from linked list: %v\n", *t.get(4))
	fmt.Println("Removing \"4\" from linked list")
	t.remove(4)
	fmt.Println(t)
	fmt.Printf("Head of linkedlist: %v", l.head)
}
