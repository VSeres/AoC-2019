package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	left    *node
	right   *node
	name    string
	parrent *node
}

var orbits = make(map[string][]string)

func main() {
	arr := parseFile("./input.txt")
	for _, v := range arr {
		orbits[v[0]] = append(orbits[v[0]], v[1])
	}

	graf := buildGraf("COM", nil)
	getOritbCount(graf, 0)
	fmt.Println(num)
	dist, found := partTwo(find(graf, "YOU"), 0)
	fmt.Println(dist-2, found)
}

func parseFile(filename string) [][2]string {
	file, _ := os.Open(filename)
	arr := make([][2]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, [2]string(strings.Split(line, ")")))
	}

	return arr
}

func buildGraf(name string, parrent *node) *node {
	currentNode := node{name: name}
	currentNode.parrent = parrent
	whatOribts := orbits[name]

	if len(whatOribts) > 0 {
		currentNode.left = buildGraf(whatOribts[0], &currentNode)
	}
	if len(whatOribts) == 2 {
		currentNode.right = buildGraf(whatOribts[1], &currentNode)
	}

	return &currentNode
}

var num int = 0

func getOritbCount(graf *node, indirect int) {
	if graf == nil {
		return
	}
	num += indirect
	getOritbCount(graf.left, indirect+1)
	getOritbCount(graf.right, indirect+1)
}

func find(graf *node, name string) *node {
	if graf == nil {
		return nil
	}
	if graf.name == name {
		return graf
	}
	var result *node

	result = find(graf.left, name)
	if result != nil {
		return result
	}

	result = find(graf.right, name)
	if result != nil {
		return result
	}

	return result
}

var visited = make(map[string]bool)

func partTwo(start *node, dist int) (int, bool) {
	if start == nil {
		return dist, false
	}
	if visited[start.name] {
		return 0, false
	}
	visited[start.name] = true

	if start.name == "SAN" {
		return dist, true
	}

	if start.name == "YOU" {
		return partTwo(start.parrent, dist+1)
	}

	if start.left == nil && start.right == nil {
		return 0, false
	}

	d, found := partTwo(start.left, dist+1)
	if found {
		return d, true
	}

	d, found = partTwo(start.right, dist+1)
	if found {
		return d, true
	}

	dist, found = partTwo(start.parrent, dist+1)

	return dist, found
}
