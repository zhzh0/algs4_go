package graph

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPaths(t *testing.T) {
	g := NewGraphFromFile("/root/code/src/algs4_go/ch4_graph/data/tinyCG.txt")
	source := 0
	pathSearch := NewBreadthFirstPaths(g, source)
	for v := 0; v < g.V(); v++ {
		fmt.Print("0 to " + strconv.Itoa(v) + ":")
		if pathSearch.HasPathTo(v) {
			for _, x := range pathSearch.PathTo(v) {
				if x == source {
					fmt.Print(x)
				} else {
					fmt.Print("-" + strconv.Itoa(x))
				}
			}
		}
		fmt.Println()
	}
}
