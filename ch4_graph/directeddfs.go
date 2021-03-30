package graph

type DirectedDFS struct {
	marked []bool
	count  int // 从s可到达的顶点个数
}

func NewDirectedDFS(g *Digraph, s int) *DirectedDFS {
	r := &DirectedDFS{
		marked: make([]bool, g.V()),
	}

	r.dfs(g, s)
	return r
}

func (d *DirectedDFS) dfs(g *Digraph, v int) {
	d.count++
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DirectedDFS) Marked(v int) bool {
	return d.marked[v]
}

func (d *DirectedDFS) Count() int {
	return d.count
}
