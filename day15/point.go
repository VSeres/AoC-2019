package main

type point struct {
	x, y int
}

func (p point) dir(other point) int {
	if p.x > other.x {
		return left
	}
	if p.x < other.x {
		return right
	}
	if p.y > other.y {
		return down
	}
	if p.y < other.y {
		return up
	}
	return -1
}

func (p point) up() point {
	return point{
		x: p.x,
		y: p.y + 1,
	}
}

func (p point) down() point {
	return point{
		x: p.x,
		y: p.y - 1,
	}
}

func (p point) left() point {
	return point{
		x: p.x - 1,
		y: p.y,
	}
}

func (p point) right() point {
	return point{
		x: p.x + 1,
		y: p.y,
	}
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}
