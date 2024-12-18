package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay14_01_2(t *testing.T) {

	// str = `p=2,4 v=2,-3`
	// str = `p=4,4 v=3,-3`

	str, _ := loadInputFromFile("14")

	times := 100

	width := 101
	height := 103

	quads := [4]int{}

	for _, line := range strings.Split(str, "\n") {
		var cmd cmd
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &cmd.Px, &cmd.Py, &cmd.Vx, &cmd.Vy)
		assert.NoError(t, err)
		pxNew, pyNew := cmd.calcNewPos(times, width, height)

		q := getQuadrant(width, height, pxNew, pyNew)
		if q != -1 {
			quads[q]++
		}
	}

	fmt.Println(quads)

	sum := 1
	for i := 0; i < 4; i++ {
		sum = sum * quads[i]
	}

	fmt.Println("Sum:", sum)

	// wrong 705600
}
func TestDay14_01(t *testing.T) {

	str := `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

	times := 100

	width := 11
	height := 7

	quads := [4]int{}

	for _, line := range strings.Split(str, "\n") {
		var cmd cmd
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &cmd.Px, &cmd.Py, &cmd.Vx, &cmd.Vy)
		assert.NoError(t, err)
		pxNew, pyNew := cmd.calcNewPos(times, width, height)

		q := getQuadrant(width, height, pxNew, pyNew)
		if q != -1 {
			quads[q]++
		}

	}

	fmt.Println(quads)

	sum := 1
	for i := 0; i < 4; i++ {
		sum = sum * quads[i]
	}

	fmt.Println("Sum:", sum)
	assert.Equal(t, sum, 12)

	// wrong 705600
}

func getQuadrant(w, h, x, y int) int {

	if x == w/2 {
		return -1 // Exclude column 5
	}

	if y == h/2 {
		return -1 // Exclude row 3
	}

	if y > h/2 {
		if x > w/2 {
			return 3 // Bottom-right
		}
		return 2 // Bottom-left
	}
	if x > w/2 {
		return 1 // Top-right
	}
	return 0 // Top-left
}

func (c cmd) calcNewPos(t, w, h int) (int, int) {
	pxNew := move(c.Px, c.Vx, w, t)
	pyNew := move(c.Py, c.Vy, h, t)
	return pxNew, pyNew
}

type cmds []cmd

type cmd struct {
	Px, Py, Vx, Vy int
}

func TestMove(t *testing.T) {

	w := 11
	p := 0 // 0 to 10
	v := 3
	ti := 5

	p = move(p, v, w, ti)

	assert.Equal(t, p, 4)

}

func move(index, vector, size, ti int) int {
	return ((index+(vector*ti))%size + size) % size
}

func TestMove2(t *testing.T) {

	w := 11
	p := 2 // 0 to 10
	v := 2
	ti := 5

	p = move(p, v, w, ti)

	assert.Equal(t, p, 1)
}

func TestMove3(t *testing.T) {

	w := 11
	p := 4 // 0 to 10
	v := -3
	ti := 2

	// 0000x00000
	// 000x000000
	// 00x0000000
	// 0x00000000 -3
	// x000000000
	// 000000000x
	// 00000000x0 -3

	p = move(p, v, w, ti)

	assert.Equal(t, p, 9)
}
