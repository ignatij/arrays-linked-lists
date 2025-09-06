package main

import (
	"fmt"
	"unsafe"

	"github.com/ignatij/arrays-vs-linked-lists/lists"
)

func main() {
	node := lists.BuildContiguousList([]int{1, 2, 3, 4})
	for {
		if node == nil {
			break
		}
		fmt.Printf("Node with value: %d address: %p\n", node.Value, node)
		if node.Next != nil {
			diff := uintptr(unsafe.Pointer(node.Next)) - uintptr(unsafe.Pointer(node))
			fmt.Printf("Difference between adjacent nodes in memory = %d bytes\n", diff)
		}

		node = node.Next
	}

}
