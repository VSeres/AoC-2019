package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type asteroid struct {
	x    int
	y    int
	hits map[float64]*asteroid
}

func (a asteroid) equals(b asteroid) bool {
	return a.x == b.x && a.y == b.y
}

func (a asteroid) distance(other asteroid) float64 {
	return math.Sqrt(math.Pow(float64(a.x-other.x), 2) + math.Pow(float64(a.y-other.y), 2))
}

func (a *asteroid) setHits(asteroids []*asteroid) {
	a.hits = make(map[float64]*asteroid, 0)
	for _, other := range asteroids {
		if a.equals(*other) {
			continue
		}
		x := other.x - a.x
		y := other.y - a.y
		angle := 90 - math.Atan2(float64(x), float64(y))*180/math.Pi
		if a.hits[angle] == nil {
			a.hits[angle] = other
		} else if a.hits[angle].distance(*a) > other.distance(*a) {
			a.hits[angle] = other
		}
	}
}

func main() {
	asteroids := parseFile("input.txt")
	station := partOne(asteroids)
	partTwo(station, asteroids)
}

func partOne(asteroids []*asteroid) *asteroid {
	var best *asteroid
	for _, asteroid := range asteroids {
		asteroid.setHits(asteroids)

		if best == nil || len(asteroid.hits) > len(best.hits) {
			best = asteroid
		}
	}
	fmt.Printf("Best:\nX: %d Y: %d visible: %d\n", best.x, best.y, len(best.hits))
	return best
}

func partTwo(station *asteroid, asteroids []*asteroid) {
	dCount := 0

	for len(asteroids) > 1 {
		keys := make([]float64, 0, len(station.hits))
		for k := range station.hits {
			keys = append(keys, k)
		}
		sort.Float64s(keys)

		for _, angle := range keys {
			dCount += 1
			removing := station.hits[angle]
			if dCount == 200 {
				fmt.Printf("200th: x: %d y: %d\n", removing.x, removing.y)
			}
			asteroids = remove(removing, asteroids)
		}
		station.setHits(asteroids)
	}

}

func remove(a *asteroid, from []*asteroid) []*asteroid {
	index := 0
	for i, v := range from {
		if v.equals(*a) {
			index = i
			break
		}
	}
	if index == 0 {
		return from[index:]
	}
	return append(from[:index-1], from[index:]...)
}

func parseFile(name string) []*asteroid {
	file, _ := os.Open(name)
	scanner := bufio.NewScanner(file)
	y := 0
	asteroids := make([]*asteroid, 0)
	for scanner.Scan() {
		line := scanner.Text()
		for x, loc := range line {
			if loc == '#' {
				asteroids = append(asteroids, &asteroid{
					x: x,
					y: y,
				})
			}
		}
		y += 1
	}

	return asteroids
}
