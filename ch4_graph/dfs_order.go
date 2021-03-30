package graph

// 先序、后序、逆后序
type DepthFirstOrder struct {
	marked []bool
	pre    []int
	post   []int

	preorder    []int
	postorder   []int
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	d := &DepthFirstOrder{
		pre:       make([]int, g.V()),
		post:      make([]int, g.V()),
		preorder:  []int{},
		postorder: []int{},
		marked:    make([]bool, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}

	return d
}

func (d *DepthFirstOrder) dfs(g *Digraph, v int) {
	d.marked[v] = true
	d.preCounter++
	d.pre[v] = d.preCounter
	d.preorder = append(d.preorder, v)
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}

	// 后序： 节点v访问完成
	d.postCounter++
	d.post[v] = d.postCounter
	d.postorder = append(d.postorder, v)
}

func (d *DepthFirstOrder) PreOrder() []int {
	return d.preorder
}

func (d *DepthFirstOrder) PostOrder() []int {
	return d.postorder
}

func (d *DepthFirstOrder) ReversePost() []int {
	stack := []int{}
	for _, v := range d.postorder {
		stack = append(stack, v)
	}
	return stack
}
