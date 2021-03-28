package graph

// api :
// hashPathTo
// pathTo
type DepthFirstPaths struct {
	marked []bool
	edgeTo []int // edgeTo[v] = last edge on s-v path, NOTE
	s      int   // source
}

// s : source
func NewDepthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	dfs := &DepthFirstPaths{}
	dfs.marked = make([]bool, g.V())
	dfs.s = s
	dfs.edgeTo = make([]int, g.V())
	dfs.validVertex(s)
	dfs.dfs(g, s)
	return dfs
}

func (d *DepthFirstPaths) dfs(g *Graph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstPaths) HasPathTo(v int) bool {
	d.validVertex(v)
	return d.marked[v]
}

func (d *DepthFirstPaths) PathTo(v int) []int {
	d.validVertex(v)
	if !d.HasPathTo(v) {
		return nil
	}
	path := make([]int, 0)
	for x := v; x != d.s; x = d.edgeTo[x] {
		path = append([]int{x}, path...)
	}
	path = append([]int{d.s}, path...)
	return path
}

func (d *DepthFirstPaths) validVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic("vertex invalid")
	}
}
