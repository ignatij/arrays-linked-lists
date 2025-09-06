package lists

type Node struct {
	Value int
	Next  *Node
}

func BuildContiguousList(values []int) *Node {
	nodes := make([]Node, 0, len(values))
	for i := range values {
		curr := Node{
			Value: values[i],
		}
		nodes = append(nodes, curr)
		if i != 0 {
			nodes[i-1].Next = &nodes[i]
		}
	}
	return &nodes[0]
}
