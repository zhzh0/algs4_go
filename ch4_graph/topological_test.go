package graph

import (
	"fmt"
	"testing"
)

func TestPaths(t *testing.T) {
	g := NewDigraphFromFile("/root/code/src/algs4_go/ch4_graph/data/tinyDG.txt")
	tt := NewTopological(g)
	fmt.Println(tt.Order())
}
