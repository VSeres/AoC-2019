package main

type point struct {
	x, y int
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
