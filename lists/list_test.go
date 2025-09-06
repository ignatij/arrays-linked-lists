package lists_test

import (
	"testing"

	"github.com/ignatij/arrays-vs-linked-lists/lists"
)

func BenchmarkArrayListTraversal(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range array {
			sum += v
		}
	}
}

func BenchmarkLinkedListTraversal(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	linkedList := lists.BuildContiguousList(array)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		node := linkedList
		for node != nil {
			sum += node.Value
			node = node.Next
		}
	}
}

func BenchmarkArrayListInsertHead(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array = append([]int{i}, array...)
		array = array[1:]
	}
}

var sink *lists.Node

func BenchmarkLinkedListInsertHead(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	head := lists.BuildContiguousList(array)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		firstNode := &lists.Node{
			Value: i,
			Next:  head,
		}
		head = firstNode
		// cleanup
		head = firstNode.Next
		// preventing compiler optimization
		sink = head
	}
}

var sinkArray []int

func BenchmarkArrayListInsertTail(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array = append(array, i)
	}
	// prevent compiler optimization
	sinkArray = array

}

func BenchmarkLinkedListInsertTail(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	head := lists.BuildContiguousList(array)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curr := head
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = &lists.Node{
			Value: i,
			Next:  nil,
		}

		// cleanup
		head.Next = nil
	}

}

func BenchmarkLinkedListInsertTailWithNode(b *testing.B) {
	size := 100_000
	array := make([]int, size)

	for i := range size {
		array[i] = i
	}

	head := lists.BuildContiguousList(array)

	for head.Next != nil {
		head = head.Next
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curr := head
		curr.Next = &lists.Node{
			Value: i,
			Next:  nil,
		}
		// cleanup
		head.Next = nil

		// prevent compiler optimization
		sink = head
	}

}
