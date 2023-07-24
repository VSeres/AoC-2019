package main

import (
	"aoc-2019/intcode"
	"fmt"
	"strings"
)

var roomMap = make(map[point]tile, 0)
var robotPos = point{}
var nodeStack = make([]*node, 0)

type tile int

const (
	unknown tile = iota
	empty
	goal
	wall
)

const (
	up    = 1
	down  = 2
	left  = 3
	right = 4
)

var dir = left
var program intcode.Program
var endNode *node

func main() {
	program = intcode.ParseFile("input.txt")
	roomMap[robotPos] = empty
	startNode := &node{
		point: robotPos,
	}

	nodeStack = append(nodeStack, startNode)
	stepC := 0
	distance := 0

	fmt.Print("\033[s")
	lookAround()
	addUnexploredNodes(startNode)
	nextNode := setNewDir(startNode)
	for !program.Stopped {
		newPos := getNewPos(dir)
		if nextNode.explored && nextNode.point.equals(robotPos) {
			nextNode = setNewDir(nextNode)
		} else if roomMap[newPos] == wall {
			exploreNode(nextNode, distance)
			nextNode = setNewDir(nextNode)
			distance = 0
			if len(nodeStack) == 0 {
				break
			}
		} else if isCrossing(robotPos) {
			exploreNode(nextNode, distance)
			nextNode = setNewDir(nextNode)
			distance = 0
		}

		move(dir)
		if !nextNode.explored {
			nextNode.point = newPos
			distance++
		}

		if roomMap[robotPos] == goal {
			endNode = nextNode
		}

		lookAround()
		stepC++
		display(stepC)
		// if stepC > 400 {
		// 	time.Sleep(150 * time.Millisecond)
		// }
	}
	fmt.Println(path(startNode, endNode, 0, point{}))
}

func isCrossing(p point) bool {
	if dir == up || dir == down {
		return roomMap[p.left()] == empty || roomMap[p.right()] == empty
	}
	return roomMap[p.up()] == empty || roomMap[p.down()] == empty
}

func setNewDir(n *node) *node {
	var nextNode *node
	if n.up != nil && !n.up.explored {
		dir = up
		nextNode = n.up
	} else if n.right != nil && !n.right.explored {
		dir = right
		nextNode = n.right
	} else if n.down != nil && !n.down.explored {
		dir = down
		nextNode = n.down
	} else if n.left != nil && !n.left.explored {
		dir = left
		nextNode = n.left
	}
	if nextNode == nil {
		prev := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		dir = n.point.dir(prev.point)
		return prev
	}
	nodeStack = append(nodeStack, n)
	return nextNode
}

func exploreNode(n *node, dist int) {
	if dir == up {
		n.downDistance = dist
		n.down.upDistance = dist
	} else if dir == down {
		n.upDistance = dist
		n.up.downDistance = dist
	} else if dir == right {
		n.leftDistance = dist
		n.left.rightDistance = dist
	} else if dir == left {
		n.rightDistance = dist
		n.right.leftDistance = dist
	}

	if addUnexploredNodes(n) == 0 {
		n.explored = true
	}

	n.explored = true
}

func addUnexploredNodes(n *node) int {
	nodeCount := 0
	if n.up == nil && roomMap[robotPos.up()] == empty {
		nodeCount += 1
		n.up = &node{point: n.point, down: n}
	}
	if n.left == nil && roomMap[robotPos.left()] == empty {
		nodeCount += 1
		n.left = &node{point: n.point, right: n}
	}
	if n.right == nil && roomMap[robotPos.right()] == empty {
		nodeCount += 1
		n.right = &node{point: n.point, left: n}
	}
	if n.down == nil && roomMap[robotPos.down()] == empty {
		nodeCount += 1
		n.down = &node{point: n.point, up: n}
	}
	return nodeCount
}

func lookAround() {
	if roomMap[robotPos.up()] == unknown && move(up) {
		move(down)
	}
	if roomMap[robotPos.right()] == unknown && move(right) {
		move(left)
	}
	if roomMap[robotPos.down()] == unknown && move(down) {
		move(up)
	}
	if roomMap[robotPos.left()] == unknown && move(left) {
		move(right)
	}
}

func display(steps int) {
	min := point{}
	max := point{}
	for pos := range roomMap {
		if min.x > pos.x {
			min.x = pos.x
		}
		if min.y > pos.y {
			min.y = pos.y
		}
		if max.x < pos.x {
			max.x = pos.x
		}
		if max.y < pos.y {
			max.y = pos.y
		}
	}
	var buff strings.Builder
	buff.WriteString(fmt.Sprintf("(% 3d,% 03d) % 04d\n", robotPos.x, robotPos.y, steps))
	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			p := roomMap[point{x, y}]
			if robotPos.x == x && robotPos.y == y && p != goal {
				buff.WriteByte('r')
				continue
			}
			switch p {
			case unknown:
				buff.WriteByte('?')
			case empty:
				buff.WriteByte(' ')
			case wall:
				buff.WriteByte('#')
			case goal:
				buff.WriteByte('x')
			}
		}
		buff.WriteByte(10)
	}
	fmt.Print("\033[u", buff.String())
}

func move(dir int) bool {
	program.SetInput(dir)
	newPos := getNewPos(dir)

	output := program.Execute()
	if len(output) == 0 {
		return false
	}
	updateMap(output[0], newPos)
	if output[0] != 0 {
		robotPos = newPos
		return true
	}
	return false
}

func updateMap(tile int, newPos point) {
	if tile == 0 {
		roomMap[newPos] = wall
	} else if tile == 1 {
		roomMap[newPos] = empty
	} else {
		roomMap[newPos] = goal
	}
}

func getNewPos(dir int) point {
	switch dir {
	case up:
		return robotPos.up()
	case right:
		return robotPos.right()
	case down:
		return robotPos.down()
	case left:
		return robotPos.left()
	}
	return point{}
}

var solution = make([]point, 32)

func path(start *node, end *node, dist int, parrent point) int {
	if start.point.equals(end.point) {
		solution = append(solution, start.point)
		return dist
	}
	distResult := -500
	if start.up != nil && !start.up.equals(parrent) && distResult <= 0 {
		distResult = path(start.up, end, dist+start.upDistance, start.point)
	}

	if start.down != nil && !start.down.equals(parrent) && distResult <= 0 {
		distResult = path(start.down, end, dist+start.downDistance, start.point)
	}

	if start.left != nil && !start.left.equals(parrent) && distResult <= 0 {

		distResult = path(start.left, end, dist+start.leftDistance, start.point)
	}

	if start.right != nil && !start.right.equals(parrent) && distResult <= 0 {

		distResult = path(start.right, end, dist+start.rightDistance, start.point)
	}
	return distResult
}
