package graph

import (
	"fmt"
	"testing"
)

func TestNewDepthFirstSearch(t *testing.T) {
	g := NewGraphFromFile("/root/code/src/algs4_go/ch4_graph/data/tinyG.txt")

	d := NewDepthFirstSearch(g, 9)

	for v := 0; v < g.V(); v++ {
		if d.Marked(v) {
			fmt.Print(v)
		}
	}

	if d.Count() != g.V() {
		fmt.Println("not")
	}
}
