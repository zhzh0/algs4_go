package graph

import (
	"math"
)

// api :
// hashPathTo
// pathTo
type BreadthFirstPaths struct {
	marked []bool // 标记哪些边已经扫描过了
	edgeTo []int  // edgeTo[v] = last edge on s-v path, NOTE
	distTo []int  // distTo[v] = s-v的最短路径的距离
	s      int
}

// s : source
func NewBreadthFirstPaths(g *Graph, s int) *BreadthFirstPaths {
	bfs := &BreadthFirstPaths{}
	bfs.marked = make([]bool, g.V())
	bfs.edgeTo = make([]int, g.V())
	bfs.distTo = make([]int, g.V())
	bfs.s = s
	for v := 0; v < g.V(); v++ {
		bfs.distTo[v] = math.MaxInt64
	}
	bfs.validVertex(s)
	bfs.bfs(g, s)
	return bfs
}

// 与dfs不同，这里不是递归，用queue实现
// 从s搜索
func (b *BreadthFirstPaths) bfs(g *Graph, s int) {
	queue := make([]int, 0)
	for v := 0; v < g.V(); v++ {
		b.distTo[v] = math.MaxInt64
	}
	b.distTo[s] = 0
	b.marked[s] = true
	queue = append(queue, s) // enqueue

	for len(queue) > 0 {
		v := queue[0] // dequeue
		queue = queue[1:]
		for _, w := range g.Adj(v) {
			if !b.marked[w] {
				b.edgeTo[w] = v
				b.distTo[w] = b.distTo[v] + 1
				b.marked[w] = true
				queue = append(queue, w)
			}
		}
	}
}

func (b *BreadthFirstPaths) HasPathTo(v int) bool {
	b.validVertex(v)
	return b.marked[v]
}

func (b *BreadthFirstPaths) PathTo(v int) []int {
	b.validVertex(v)
	if !b.HasPathTo(v) {
		return nil
	}
	path := make([]int, 0)
	for x := v; x != b.s; x = b.edgeTo[x] {
		path = append([]int{x}, path...)
	}
	path = append([]int{b.s}, path...)
	return path
}

func (b *BreadthFirstPaths) validVertex(v int) {
	V := len(b.marked)
	if v < 0 || v >= V {
		panic("vertex invalid")
	}
}

func (b *BreadthFirstPaths) DistTo(v int) int {
	b.validVertex(v)
	return b.distTo[v]
}
