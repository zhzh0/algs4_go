package graph

// 连通分量api
// connected(int v, int w): 是否连通
// count() : 连通分量个数
// id(int v):  v 所在的连通分量id
type CC struct {
	marked []bool
	count  int   // 连通分量的个数
	id     []int // id[v] : 表示v所在连通分量的id
	size   []int // size[id] : 连通分量id的大小
}

// s : source
func NewCC(g *Graph) *CC {
	dfs := &CC{}
	dfs.marked = make([]bool, g.V())
	dfs.id = make([]int, g.V())
	dfs.size = make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		if !dfs.marked[v] {
			dfs.dfs(g, v)
			dfs.count++
		}
	}
	return dfs
}

func (d *CC) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.id[v] = d.count
	d.size[d.count]++
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (c *CC) Id(v int) int {
	c.validVertex(v)
	return c.id[v]
}

func (c *CC) Size(v int) int {
	c.validVertex(v)
	return c.size[v]
}

func (c *CC) Count() int {
	return c.count
}

func (c *CC) Connected(v, w int) bool {
	c.validVertex(v)
	c.validVertex(w)
	return c.id[v] == c.id[w]
}

func (d *CC) validVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic("vertex invalid")
	}
}
