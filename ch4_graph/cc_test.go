package graph

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCC(t *testing.T) {
	g := NewGraphFromFile("/root/code/src/algs4_go/ch4_graph/data/tinyG.txt")
	cc := NewCC(g)
	M := cc.Count()
	fmt.Println(strconv.Itoa(M) + " components")
	// print
	ret := make([][]int, M)
	for v := 0; v < g.V(); v++ {
		ret[cc.Id(v)] = append([]int{v}, ret[cc.Id(v)]...)
	}

	for i := 0; i < M; i++ {
		for _, v := range ret[i] {
			fmt.Print(strconv.Itoa(v) + " ")
		}
		fmt.Println()
	}
}
