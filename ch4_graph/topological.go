package graph

// 拓扑排序
// api
// order : 拓扑顺序
type Topological struct {
	order []int // 拓扑顺序
	rank  []int // rank[v] 表示v在顺序中rank
}

func NewTopological(g *Digraph) *Topological {
	dc := NewDirectedCycle(g)
	if dc.HasCycle() {
		return nil
	}
	dfsOrder := NewDepthFirstOrder(g)
	//order := dfsOrder.ReversePost()

	t := &Topological{
		order: dfsOrder.ReversePost(),
		rank:  make([]int, g.V()),
	}
	t.rank = make([]int, g.V())
	for idx, v := range t.order {
		t.rank[v] = idx
	}
	return t
}

func (t *Topological) Order() []int {
	return t.order
}

func (t *Topological) HasOrder() bool {
	return len(t.order) > 0
}

func (t *Topological) IsDAG() bool {
	return t.HasOrder()
}
