package graph

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 图的实现，邻接表
type Graph struct {
	v   int // 节点
	e   int // edge
	adj [][]int
}

func NewGraph(v int) *Graph {
	if v <= 0 {
		return nil
	}

	g := &Graph{}
	g.v = v
	g.adj = make([][]int, v)
	for i := 0; i < v; i++ {
		g.adj[i] = []int{}
	}
	return g
}

func NewGraphFromFile(fp string) *Graph {
	g := &Graph{}
	lines := readLines(fp)
	g.v, _ = strconv.Atoi(lines[0])
	g.adj = make([][]int, g.v)
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

func readLines(fp string) []string {
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func (g *Graph) validVertex(v int) {
	if v < 0 || v >= g.v {
		panic("invalid vertex")
	}
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

// v-w
func (g *Graph) AddEdge(v, w int) {
	g.validVertex(v)
	g.validVertex(w)
	g.adj[v] = append([]int{w}, g.adj[v]...)
	g.adj[w] = append([]int{v}, g.adj[w]...)
	g.e++
}

func (g *Graph) Adj(v int) []int {
	g.validVertex(v)
	return g.adj[v]
}

func (g *Graph) Degree(v int) int {
	g.validVertex(v)
	return len(g.adj[v])
}

func (g *Graph) ToString() string {
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
