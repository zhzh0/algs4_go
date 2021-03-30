package graph

import (
	"strconv"
	"strings"
)

// 图的实现，邻接表
type Digraph struct {
	v        int     // 节点
	e        int     // edge
	adj      [][]int // 邻接表
	indegree []int   // 记录每个顶点的入度
}

func NewDigraph(v int) *Digraph {
	if v <= 0 {
		return nil
	}

	g := &Digraph{}
	g.v = v
	g.adj = make([][]int, v)
	for i := 0; i < v; i++ {
		g.adj[i] = []int{}
	}
	g.indegree = make([]int, v)
	return g
}

func NewDigraphFromFile(fp string) *Digraph {
	g := &Digraph{}
	lines := readLines(fp)
	g.v, _ = strconv.Atoi(lines[0])
	g.adj = make([][]int, g.v)
	g.indegree = make([]int, g.v)
	e, _ := strconv.Atoi(lines[1])
	lines = lines[2:]
	for i := 0; i < e; i++ {
		vw := strings.Split(lines[i], " ")
		v, _ := strconv.Atoi(vw[0])
		w, _ := strconv.Atoi(vw[1])
		g.AddEdge(v, w)
	}
	return g
}

func (g *Digraph) validVertex(v int) {
	if v < 0 || v >= g.v {
		panic("invalid vertex")
	}
}

func (g *Digraph) V() int {
	return g.v
}

func (g *Digraph) E() int {
	return g.e
}

// v-w
// NOTE 和无向图主要的区别点
func (g *Digraph) AddEdge(v, w int) {
	g.validVertex(v)
	g.validVertex(w)
	g.adj[v] = append([]int{w}, g.adj[v]...)
	g.indegree[w]++ // 更新入度
	g.e++
}

func (g *Digraph) Adj(v int) []int {
	g.validVertex(v)
	return g.adj[v]
}

func (g *Digraph) Reverse() *Digraph {
	newg := Digraph{
		v:        g.v,
		e:        g.e,
		indegree: make([]int, g.v),
		adj:      make([][]int, g.v),
	}
	for i := 0; i < g.v; i++ {
		for _, w := range g.adj[i] {
			newg.AddEdge(w, i)
		}
	}
	return &newg
}

func (g *Digraph) ToString() string {
	var s string
	for v := 0; v < g.v; v++ {
		s += strconv.Itoa(v) + ":"
		for _, item := range g.adj[v] {
			s += strconv.Itoa(item) + " "
		}
		s += "\n"
	}
	return s
}
