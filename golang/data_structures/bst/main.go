package bst

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	size int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(root.String())
	b.inOrderTraversalByNode(sb, root.right)
}

func (b *bst) insert(value int) *node {
	b.root = b.insertByNode(b.root, value)
	return b.root
}

func (b *bst) insertByNode(root *node, value int) *node {
	if root == nil {
		b.size++
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.insertByNode(root.left, value)
	} else {
		root.right = b.insertByNode(root.right, value)
	}
	return root
}

func (b *bst) remove(value int) *node {
	return b.removeByNode(b.root, value)
}

func (b *bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		root.value = b.minValue(root.right)

		root.right = b.removeByNode(root.right, root.value)
	}
	return root
}

func (b bst) minValue(root *node) int {
	min := root.value
	for root.left != nil {
		min = root.left.value
		root = root.left
	}
	return min
}

func (b bst) maxValue(root *node) int {
	max := root.value
	for root.right != nil {
		max = root.right.value
		root = root.right
	}
	return max
}

func (b *bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		// at the bottom of the bst
		return nil, false
	}

	if value == root.value {
		// found the node
		return root, true
	} else if value < root.value {
		// recursively pass the left sub tree
		return b.searchByNode(root.left, value)
	} else {
		// recursively pass the right sub tree
		return b.searchByNode(root.right, value)
	}
}

func Run() {
	n := &node{7, nil, nil}
	n.left = &node{2, nil, nil}
	n.right = &node{8, nil, nil}

	b := bst{
		root: n,
		size: 3,
	}
	fmt.Printf("Created tree with 2, 7, 8: %v\n\n", b)

	b.insert(5)
	b.insert(4)
	b.insert(6)
	b.insert(3)

	fmt.Printf("Added nodes 5, 4, 6, 3 to bst: %v\n\n", b)

	searchRes, ok := b.search(5)
	fmt.Printf("Searching for node with value of 3:\nFound: %v\nValue: %v\n\n", ok, searchRes)

	searchRes, ok = b.search(27)
	fmt.Printf("Searching for node with value of 27:\nFound: %v\nValue: %v\n\n", ok, searchRes)

	b.remove(6)
	b.remove(27)

	fmt.Printf("Removing 6 and 27 from bst: %v\n\n", b)

	fmt.Printf("Min node value in bst: %v\n\n", b.minValue((b.root)))
	fmt.Printf("Max node value in bst: %v\n\n", b.maxValue((b.root)))
}
