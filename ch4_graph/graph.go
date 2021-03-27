package graph

import (
	f "algs4_go/ch1_fundamentals"
	"context"
	"strconv"
)

type Graph struct {
	v   int // 节点
	e   int // edge
	adj []*f.Bag
}

func NewGraph(v int) *Graph {
	if v <= 0 {
		return nil
	}

	g := &Graph{}
	g.v = v
	g.adj = make([]*f.Bag, v)
	for i := 0; i < v; i++ {
		g.adj[i] = f.NewBag()
	}
	return g
}

func (g *Graph) validVertex(v int) {
	if v < 0 || v >= g.v {
		panic("invalid vertex")
	}
}

// v-w
func (g *Graph) AddEdge(v, w int) {
	g.validVertex(v)
	g.validVertex(w)
	g.adj[v].Add(f.ItemType(w))
	g.adj[w].Add(f.ItemType(v))
	g.e++
}

func (g *Graph) Adj(v int) *f.Bag {
	g.validVertex(v)
	return g.adj[v]
}

func (g *Graph) Degree(v int) int {
	g.validVertex(v)
	return g.adj[v].Size()
}

func (g *Graph) ToString() string {
	var s string
	s += "v, e, \n"
	for v := 0; v < g.v; v++ {
		s += strconv.Itoa(v) + ":"
		for item := range g.adj[v].Chan(context.TODO()) {
			s += strconv.Itoa(int(item)) + " "
		}
		s += "\n"
	}
	return s
}
