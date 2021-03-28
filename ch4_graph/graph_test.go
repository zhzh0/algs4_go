package graph

import "testing"

func TestGraph(t *testing.T) {
	g := NewGraph(3)
	t.Log(g)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	t.Log(g.ToString())

	g = NewGraphFromFile("/root/code/src/algs4_go/ch4_graph/data/tinyG.txt")
	t.Log(g.ToString())
}
