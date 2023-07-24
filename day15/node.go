package main

type node struct {
	point
	up            *node
	upDistance    int
	right         *node
	rightDistance int
	down          *node
	downDistance  int
	left          *node
	leftDistance  int
	explored      bool
}
