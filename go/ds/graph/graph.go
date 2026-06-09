package ds

type Node string

type Graph struct {
	Vertices map[Node]struct{}
	Edges    map[Node]map[Node]float64
	directed bool
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		Vertices: map[Node]struct{}{},
		Edges:    map[Node]map[Node]float64{},
		directed: directed,
	}
}

func (g *Graph) AddVertice(node Node) {
	g.Vertices[node] = struct{}{}
}

func (g *Graph) AddEdge(from, to Node, weight float64) {
	g.AddVertice(from)
	g.AddVertice(to)

	if g.Edges[from] == nil {
		g.Edges[from] = map[Node]float64{}
	}
	g.Edges[from][to] = weight

	if !g.directed {
		if g.Edges[to] == nil {
			g.Edges[to] = map[Node]float64{}
		}
		g.Edges[to][from] = weight
	}
}

func (g *Graph) String() string {
	if g == nil {
		return "<empty>"
	}
	return ""
}
