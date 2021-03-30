package graph

type Cycle struct {
	marked []bool
	edgeTo []int // 倒序的路径记录
	cycle  []int // stack(array): 记录具体的环
}

// s : source
func NewCycle(g *Graph) *Cycle {
	if hasParallelEdges(g) {
		return nil
	}
	dfs := &Cycle{}
	dfs.marked = make([]bool, g.V())
	dfs.edgeTo = make([]int, g.V())
	dfs.cycle = make([]int, 0)
	for v := 0; v < g.V(); v++ {
		if !dfs.marked[v] {
			dfs.dfs(g, -1, v)
		}
	}
	return dfs
}

func (c *Cycle) hasParallelEdges(g *Graph) bool {
	marked := make([]bool, g.V())

	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if marked[w] {
				c.cycle = append(c.cycle, []int{v, w, v}...)
				return true
			}
			marked[w] = true
		}

		for v := 0; v < g.V(); v++ {
			marked[v] = false
		}
	}
	return false
}

func (d *Cycle) dfs(g *Graph, u, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if len(d.cycle) > 0 {
			return
		} else if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, v, w)
		} else if w != u {
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = append(d.cycle, x)
			}
			d.cycle = append(d.cycle, w)
			d.cycle = append(d.cycle, v)
		}
	}
}

func (d *Cycle) validVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic("vertex invalid")
	}
}

func (c *Cycle) HasCycle() bool {
	return len(c.cycle) > 0
}

func (c *Cycle) Cycle() []int {
	return c.cycle
}
