package graph

// api
// hascycle: 是否有环
// cycle: 给出具体的环

type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	onStack []bool // 递归调用栈上的顶点
	cycle   []int  // 记录环
}

func NewDirectedCycle(g *Digraph) *DirectedCycle {
	d := DirectedCycle{
		marked:  make([]bool, g.V()),
		onStack: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
		cycle:   []int{},
	}

	for v := 0; v < g.V(); v++ {
		if !d.marked[v] && len(d.cycle) == 0 {
			d.dfs(g, v)
		}
	}

	return &d
}

func (d *DirectedCycle) dfs(g *Digraph, v int) {
	d.onStack[v] = true
	d.marked[v] = true
	for _, w := range g.Adj(v) { // 三分支
		if len(d.cycle) > 0 { // 已经有环
			return
		} else if !d.marked[w] { // 无环继续
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] { // 有环
			for x := v; x != w; x = d.edgeTo[x] {
				d.cycle = append([]int{x}, d.cycle...)
			}
			d.cycle = append([]int{w}, d.cycle...)
			d.cycle = append([]int{v}, d.cycle...)
		}
	}
	d.onStack[v] = false // 回溯
}

func (d *DirectedCycle) HasCycle() bool {
	return len(d.cycle) > 0
}

func (d *DirectedCycle) Cycle() []int {
	return d.cycle
}
