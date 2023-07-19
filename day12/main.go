package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type vector struct {
	x, y, z int
}

func (v vector) equals(other vector) bool {
	return v.x == other.x && v.y == other.y && v.z == other.z
}

func (v vector) toString() string {
	return fmt.Sprintf("<x=% 4v, y=% 4v, z=% 4v>", v.x, v.y, v.z)
}

func (v *vector) add(other vector) {
	v.x += other.x
	v.y += other.y
	v.z += other.z
}

type naturalSatellite struct {
	position vector
	velocity vector
}

func (m naturalSatellite) equals(other naturalSatellite) bool {
	return m.position.equals(other.position) && m.velocity.equals(other.velocity)
}

func (m *naturalSatellite) updatePos() {
	m.position.add(m.velocity)
}

func (m *naturalSatellite) energy() int {
	pot := abs(m.position.x) + abs(m.position.y) + abs(m.position.z)
	kin := abs(m.velocity.x) + abs(m.velocity.y) + abs(m.velocity.z)
	return kin * pot
}
func (m *naturalSatellite) print() {
	fmt.Printf("pos=%s\tvel=%s\n", m.position.toString(), m.velocity.toString())
}
func (m *naturalSatellite) updateVelocity(moons []*naturalSatellite) {
	for _, otherMoon := range moons {
		if m.equals(*otherMoon) {
			continue
		}
		if m.position.x > otherMoon.position.x {
			m.velocity.x -= 1
		} else if m.position.x < otherMoon.position.x {
			m.velocity.x += 1
		}

		if m.position.z > otherMoon.position.z {
			m.velocity.z -= 1
		} else if m.position.z < otherMoon.position.z {
			m.velocity.z += 1
		}

		if m.position.y > otherMoon.position.y {
			m.velocity.y -= 1
		} else if m.position.y < otherMoon.position.y {
			m.velocity.y += 1
		}
	}
}

func abs(n int) int {
	return ((n ^ (n >> 31)) - (n >> 31))
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile(`=(-?\d+)`)
	moons := make([]*naturalSatellite, 4)
	i := 0
	for scanner.Scan() {
		numbers := reg.FindAllStringSubmatch(scanner.Text(), -1)
		x, _ := strconv.Atoi(numbers[0][1])
		y, _ := strconv.Atoi(numbers[1][1])
		z, _ := strconv.Atoi(numbers[2][1])
		pos := vector{x: x, y: y, z: z}
		moon := naturalSatellite{
			position: pos,
		}
		moons[i] = &moon
		i++
	}
	origMoons := make([]naturalSatellite, 4)
	for i, v := range moons {
		origMoons[i] = *v
	}
	step := 0
	xP := 0
	yP := 0
	zP := 0
	for xP == 0 || yP == 0 || zP == 0 {
		for _, moon := range moons {
			moon.updateVelocity(moons)
		}
		for _, moon := range moons {
			moon.updatePos()
		}
		xp := true
		yp := true
		zp := true
		for i, moon := range moons {
			if moon.position.x != origMoons[i].position.x || moon.velocity.x != origMoons[i].velocity.x {
				xp = false
			}
			if moon.position.y != origMoons[i].position.y || moon.velocity.y != origMoons[i].velocity.y {
				yp = false
			}
			if moon.position.z != origMoons[i].position.z || moon.velocity.z != origMoons[i].velocity.z {
				zp = false
			}
		}

		step++
		if xp && xP == 0 {
			xP = step
		}
		if yp && yP == 0 {
			yP = step
		}
		if zp && zP == 0 {
			zP = step
		}
		if period(moons, origMoons) && step != 0 {
			fmt.Println(step)
			break
		}
		if step == 1000 {
			totalEnergy := 0
			for _, moon := range moons {
				totalEnergy += moon.energy()
			}
			// fmt.Println(moons)
			fmt.Println(totalEnergy)
		}
	}
	fmt.Println(LCM(xP, yP, zP))
}

func period(moons []*naturalSatellite, origMoons []naturalSatellite) bool {
	for i, moon := range moons {
		if !moon.position.equals(origMoons[i].position) || !moon.velocity.equals(origMoons[i].velocity) {
			return false
		}
	}
	return true
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
