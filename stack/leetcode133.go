package stack

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	deepClone(node, visited)
	return visited[node]
}

func deepClone(node *Node, visited map[*Node]*Node) {
	if _, has := visited[node]; has {
		return
	}
	visited[node] = &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
	for _, v := range node.Neighbors {
		deepClone(v, visited)
		visited[node].Neighbors = append(visited[node].Neighbors, visited[v])
	}
}
