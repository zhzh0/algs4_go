package graph

// 负责search任务
type DepthFirstSearch struct {
	marked []bool
	count  int
}

// s : source
func NewDepthFirstSearch(g *Graph, s int) *DepthFirstSearch {
	dfs := &DepthFirstSearch{}
	dfs.marked = make([]bool, g.V())
	dfs.validVertex(s)
	dfs.dfs(g, s)
	return dfs
}

func (d *DepthFirstSearch) dfs(g *Graph, v int) {
	d.count++
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstSearch) validVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic("vertex invalid")
	}
}

// is there a path between the source vertex and vertex v
func (d *DepthFirstSearch) Marked(v int) bool {
	d.validVertex(v)
	return d.marked[v]
}

// returns the number of vertices connected to the source vertex
func (d *DepthFirstSearch) Count() int {
	return d.count
}
