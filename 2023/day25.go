package main

import (
	"fmt"
	"strings"

	"lukechampine.com/advent/utils"
)

var input = utils.Input(2023, 25)

type node struct {
	name  string
	edges map[string]struct{}
	size  int
}

func (n *node) String() string {
	var edges []string
	for e := range n.edges {
		edges = append(edges, e)
	}
	return fmt.Sprintf("%v<->{%v}", n.name, strings.Join(edges, ", "))
}

func karger(edges [][2]string) int {
	// Karger's algorithm

	// create a new graph that tracks how many nodes were merged into each node
	g := make(map[string]*node)
	for _, e := range edges {
		if g[e[0]] == nil {
			g[e[0]] = &node{
				name:  e[0],
				edges: make(map[string]struct{}),
				size:  1,
			}
		}
		if g[e[1]] == nil {
			g[e[1]] = &node{
				name:  e[1],
				edges: make(map[string]struct{}),
				size:  1,
			}
		}
		a, b := g[e[0]], g[e[1]]
		a.edges[b.name] = struct{}{}
		b.edges[a.name] = struct{}{}
	}

	//	rand.Shuffle(len(edges), reflect.Swapper(edges))

	deleteRandom := func(g map[string]*node) {
		// pick a random edge

		allEdges := make(map[string]struct{})
		for a := range g {
			for b := range g[a].edges {
				allEdges[a+":"+b] = struct{}{}
				allEdges[b+":"+a] = struct{}{}
			}
		}
		var se string
		for se = range allEdges {
			break
		}
		ab := strings.Split(se, ":")
		a, b := g[ab[0]], g[ab[1]]

		// merge the two nodes
		for e := range b.edges {
			a.edges[e] = struct{}{}
			if g[e] != nil {
				delete(g[e].edges, b.name)
				g[e].edges[a.name] = struct{}{}
			}
		}
		delete(a.edges, a.name)
		delete(a.edges, b.name)
		a.size += b.size
		delete(g, b.name)
	}
	printGraph := func(g map[string]*node) {
		t := 0
		for _, n := range g {
			utils.Println(n.name, n.size, n.edges)
			t += n.size
		}
		utils.Println("total nodes:", t)
		utils.Println()
	}
	_ = printGraph

	// delete edges at random until only two nodes remain
	for len(g) > 2 {
		deleteRandom(g)
		//printGraph(g)
	}
	//printGraph(g)
	product := 1
	for _, n := range g {
		product *= n.size
	}
	return product
}

func main() {
	if false {
		input = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`
	}
	var edges [][2]string
	for _, line := range utils.Lines(input) {
		fs := strings.Fields(line)
		fs[0] = strings.Trim(fs[0], ":")
		for _, f := range fs[1:] {
			edges = append(edges, [2]string{fs[0], f})
		}
	}

	g := make(map[string]*node)
	for _, e := range edges {
		if g[e[0]] == nil {
			g[e[0]] = &node{
				name:  e[0],
				edges: make(map[string]struct{}),
				size:  1,
			}
		}
		if g[e[1]] == nil {
			g[e[1]] = &node{
				name:  e[1],
				edges: make(map[string]struct{}),
				size:  1,
			}
		}
		a, b := g[e[0]], g[e[1]]
		a.edges[b.name] = struct{}{}
		b.edges[a.name] = struct{}{}
	}

	// sfd->ljm
	// gst -> rph
	// jkn -> cfn
	delete(g["sfd"].edges, "ljm")
	delete(g["ljm"].edges, "sfd")
	delete(g["gst"].edges, "rph")
	delete(g["rph"].edges, "gst")
	delete(g["jkn"].edges, "cfn")
	delete(g["cfn"].edges, "jkn")

	count := func(g map[string]*node, start string) int {
		seen := make(map[string]struct{})
		var visit func(string)
		visit = func(n string) {
			seen[n] = struct{}{}
			for e := range g[n].edges {
				if _, ok := seen[e]; !ok {
					visit(e)
				}
			}
		}
		visit(start)
		return len(seen)
	}

	utils.Println(count(g, "sfd") * count(g, "ljm"))

	// sizes := make(map[int]int)
	// for i := 0; i < 10000; i++ {
	// 	if i%1 == 0 {
	// 		var bySize [][2]int
	// 		for k, v := range sizes {
	// 			bySize = append(bySize, [2]int{k, v})
	// 		}
	// 		sort.Slice(bySize, func(i, j int) bool {
	// 			return bySize[i][1] > bySize[j][1]
	// 		})
	// 		utils.Println(bySize)
	// 	}

	// 	sizes[karger(edges)]++
	// }
}
