package main

import (
	"fmt"

	"github.com/lukechampine/advent/utils"
)

const input = `42/37
28/28
29/25
45/8
35/23
49/20
44/4
15/33
14/19
31/44
39/14
25/17
34/34
38/42
8/42
15/28
0/7
49/12
18/36
45/45
28/7
30/43
23/41
0/35
18/9
3/31
20/31
10/40
0/22
1/23
20/47
38/36
15/8
34/32
30/30
30/44
19/28
46/15
34/50
40/20
27/39
3/14
43/45
50/42
1/33
6/39
46/44
22/35
15/20
43/31
23/23
19/27
47/15
43/43
25/36
26/38
1/10`

type comp struct {
	portA int
	portB int
	next  []*comp
}

func (c *comp) maxWeight(port int, chain []*comp) int {
	if len(c.next) == 0 {
		return c.portA + c.portB
	}
	var nextPort int
	if port == c.portA {
		nextPort = c.portB
	} else {
		nextPort = c.portA
	}
	maxChild := 0
outer:
	for _, child := range c.next {
		for _, cc := range chain {
			if cc == child {
				continue outer
			}
		}
		if child.portA == nextPort || child.portB == nextPort {
			maxChild = utils.Max(maxChild, child.maxWeight(nextPort, append(chain, c)))
		}
	}
	return c.portA + c.portB + maxChild
}

func (c *comp) maxLenChain(port int, chain []*comp) []*comp {
	var nextPort int
	if port == c.portA {
		nextPort = c.portB
	} else {
		nextPort = c.portA
	}
	maxChain := chain
outer:
	for _, child := range c.next {
		for _, cc := range chain {
			if cc == child {
				continue outer
			}
		}
		if child.portA == nextPort || child.portB == nextPort {
			cChain := child.maxLenChain(nextPort, append(chain, child))
			if len(cChain) > len(maxChain) ||
				(len(cChain) == len(maxChain) && chainWeight(cChain) > chainWeight(maxChain)) {
				maxChain = append([]*comp(nil), cChain...)
			}
		}
	}
	return maxChain
}

func chainWeight(chain []*comp) int {
	w := 0
	for _, c := range chain {
		w += c.portA + c.portB
	}
	return w
}

func parse(s string) []*comp {
	var cs []*comp
	for _, line := range utils.Lines(s) {
		var c comp
		utils.Sscanf(line, "%d/%d", &c.portA, &c.portB)
		cs = append(cs, &c)
	}
	return cs
}

func link(comps []*comp) {
	for i, c := range comps {
		for j, cc := range comps {
			if i == j {
				continue
			}
			if c.portA == cc.portA || c.portA == cc.portB ||
				c.portB == cc.portA || c.portB == cc.portB {
				comps[i].next = append(comps[i].next, cc)
			}
		}
	}
}

func main() {
	// part 1
	comps := parse(input)
	link(comps)
	maxWeight := 0
	for _, c := range comps {
		if c.portA == 0 || c.portB == 0 {
			maxWeight = utils.Max(maxWeight, c.maxWeight(0, nil))
		}
	}
	fmt.Println(maxWeight)

	var maxChain []*comp
	for _, c := range comps {
		if c.portA == 0 || c.portB == 0 {
			chain := c.maxLenChain(0, []*comp{c})
			if len(chain) > len(maxChain) ||
				(len(chain) == len(maxChain) && chainWeight(chain) > chainWeight(maxChain)) {
				maxChain = chain
			}
		}
	}
	fmt.Println(chainWeight(maxChain))

}
