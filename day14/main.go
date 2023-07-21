package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type chemical struct {
	name    string
	produce int
	cost    map[string]int
}

var edges map[string]int = make(map[string]int)
var chemicals = make(map[string]chemical, 16)

func main() {
	parsInput("input.txt")
	fmt.Println(top(1))
	target := 1000000000000
	l := 1
	r := 10000000000
	m := 0
	closest := 0
	for l <= r {
		m = (l + r) / 2
		result := top(m)
		if result < target {
			l = m + 1
		} else if result > target {
			r = m - 1
		} else {
			break
		}
		if result > closest && result < target {
			closest = m
		}
	}
	fmt.Println(closest)
}

func div(a int, b int) int {
	return (a + b + -1) / b
}

func top(nFuel int) int {
	edgeCopy := make(map[string]int, len(edges))
	for k, v := range edges {
		edgeCopy[k] = v
	}
	req := make(map[string]int, len(chemicals))
	req["FUEL"] = nFuel
	for {
		for name, edge := range edgeCopy {
			if edge == 0 {
				runs := div(req[name], chemicals[name].produce)
				if name == "ORE" {
					return req[name]
				}
				for n, c := range chemicals[name].cost {
					edgeCopy[n] -= 1
					req[n] += c * runs
				}
				delete(edgeCopy, name)
			}
		}
	}

}

func parsInput(name string) {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		reaction := strings.Split(line, "=>")
		produce := 0
		name := ""
		fmt.Sscanf(reaction[1], " %d %s ", &produce, &name)
		newChemical := chemical{
			name:    name,
			produce: produce,
			cost:    make(map[string]int),
		}

		for _, costString := range strings.Split(reaction[0], ",") {
			name := ""
			cost := 0
			fmt.Sscanf(strings.Trim(costString, " "), "%d %s", &cost, &name)
			newChemical.cost[name] = cost
			edges[name] += 1
		}
		chemicals[newChemical.name] = newChemical
		chemicals["ORE"] = chemical{name: "ORE", produce: 1}
		edges["FUEL"] = 0
	}
}
