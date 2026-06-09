package ds

type queue struct {
	nodes []Node
}

func newQueue() *queue {
	return &queue{nodes: []Node{}}
}

func (q *queue) enqueue(n Node) {
	q.nodes = append(q.nodes, n)
}

func (q *queue) dequeue() Node {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *queue) isEmpty() bool {
	return len(q.nodes) == 0
}

func BFS(graph *Graph, start Node, visit func(Node)) {

	if _, ok := graph.Vertices[start]; !ok {
		return
	}

	visited := map[Node]bool{}
	queue := newQueue()
	visited[start] = true
	queue.enqueue(start)
	for !queue.isEmpty() {

		currentNode := queue.dequeue()
		visit(currentNode)
		for neighbor := range graph.Edges[currentNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.enqueue(neighbor)
			}
		}
	}
}
