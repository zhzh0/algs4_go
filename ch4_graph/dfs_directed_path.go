package graph

type DepthFirstDirectedPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDepthFirstDirectedPaths(g *Digraph, s int) *DepthFirstDirectedPaths {
	d := &DepthFirstDirectedPaths{
		marked: make([]bool, g.V()),
		edgeTo: make([]int, g.V()),
		s:      s,
	}

	d.dfs(g, s)
	return d
}

func (d *DepthFirstDirectedPaths) dfs(g *Digraph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstDirectedPaths) HasPathTo(v int) {
	return d.marked[v]
}

func (d *DepthFirstDirectedPaths) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return nil
	}

	var ret []int
	for x := v; x != d.s; x = d.edgeTo[x] {
		ret = append([]int{x}, ret...)
	}
	ret = append([]int{d.s}, ret...)
	return ret
}
