package graph

// 双色问题,二分图判断
type Bipartite struct {
	isBipartite bool
	color       []bool // color[v] 标记颜色
	marked      []bool
	edgeTo      []int
	cycle       []int // slice充当stack, NOTE odd-length cycle
}

func NewBipartite(g *Graph) *Bipartite {
	b := &Bipartite{
		isBipartite: true,
		color:       make([]bool, g.V()),
		marked:      make([]bool, g.V()),
		edgeTo:      make([]int, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if !b.marked[v] {
			b.dfs(g, v)
		}
	}
	return b
}

func (b *Bipartite) dfs(g *Graph, v int) {
	b.marked[v] = true
	for _, w := range g.Adj(v) {
		if len(b.cycle) > 0 {
			return
		}
		if !b.marked[w] {
			b.edgeTo[w] = v
			b.color[w] = !color[v] // NOTE
			dfs(g, w)
		} else if b.color[w] == b.color[v] {
			b.isBipartite = false
			b.cycle = append(b.cycle, w)
			for x := v; x != w; x = b.edgeTo[x] {
				b.cycle = append(b.cycle, x)
			}
			b.cycle = append(b.cycle, w)
		}
	}
}

func (b *Bipartite) IsBipartite() {
	return b.isBipartite
}

func (b *Bipartite) Color(v int) bool {
	if !b.isBipartite {
		panic("graph is not Bipartite")
	}
	return b.color[v]
}

func (b *Bipartite) OddCycle() []int {
	return b.cycle
}
