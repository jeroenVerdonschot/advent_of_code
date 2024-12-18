package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay16_1(t *testing.T) {

	str := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

	matrix := NewMatrixFromString(str)

	fmt.Println(matrix.String())

	posS := getString("S", matrix)
	posE := getString("E", matrix)

	fmt.Println(posS, posE)

}

// type walker struct {
// 	row, col int
// 	curVec   []int
// 	curPath  *path
// }

func TestFindPath(t *testing.T) {

	str := `E...
###.
#...
#S##`

	matrix := NewMatrixFromString(str)

	fmt.Println(matrix.String())

	posS := getString("S", matrix)
	posE := getString("E", matrix)

	fmt.Println(posS, posE)

	fmt.Println("_________________")

	queue := []*path{}

	id := 0

	curPath := newPath(posS, UP, id)

	queue = append(queue, curPath)

	for len(queue) > 0 {

		perps := curPath.getPerpendiculars(&matrix)

		if len(perps) > 0 {
			// create new path
			for _, vec := range perps {
				id++
				// fmt.Println("push id ", id, "vec", vec)
				newPath := newPath([]int{curPath.row, curPath.col}, vec, id)

				newPath.row = curPath.row + vec[0]
				newPath.col = curPath.col + vec[1]

				newPath.prevPaths = append(curPath.prevPaths, curPath)

				queue = append(queue, newPath)
			}
		}

		if !curPath.walk(&matrix) {
			fmt.Println("walked", curPath)
			// pop stack
			queue = queue[1:]
			curPath = queue[0]
			fmt.Println("new path", curPath)

		}

		if curPath.row == posE[0] && curPath.col == posE[1] {
			// finalize
			// curPath.endRow = .row + w.curVec[0] // check with E pos
			// curPath.endCol = w.col + w.curVec[1]

			break
		}

		fmt.Println("current pos", curPath.row, curPath.col)
		// fmt.Println(matrix.String())
	}

	fmt.Println("end", curPath)
	fmt.Println("end")

}

// type route []*path

func newPath(pos, vec []int, id int) *path {
	return &path{startRow: pos[0], startCol: pos[1], row: pos[0], col: pos[1], vec: vec, id: id}
}

type path struct {
	startRow, startCol int
	row, col           int
	vec                []int
	prevPaths          []*path
	id                 int
	// endPaths           []*path
}

func (p *path) getPerpendiculars(matrix *Matrix[string]) [][]int {

	vec := p.vec
	perps := [][]int{}

	if (vec[0] == UP[0] && vec[1] == UP[1]) || (vec[0] == DOWN[0] && vec[1] == DOWN[1]) {
		str, err := matrix.GetLeft(p.row, p.col)
		if err == nil && *str != "#" {
			perps = append(perps, LEFT)
		}
		str, err = matrix.GetRight(p.row, p.col)
		if err == nil && *str != "#" {
			perps = append(perps, RIGHT)
		}

	} else {
		str, err := matrix.GetUp(p.row, p.col)
		if err == nil && *str != "#" {
			perps = append(perps, UP)
		}
		str, err = matrix.GetDown(p.row, p.col)
		if err == nil && *str != "#" {
			perps = append(perps, DOWN)
		}
	}

	return perps

}

// func NewWalker(row, col int) walker {
// 	return walker{row: row, col: col, curVec: RIGHT}
// }

func (p *path) walk(matrix *Matrix[string]) bool {

	newRow := p.row + p.vec[0]
	newCol := p.col + p.vec[1]

	str, err := matrix.Get(newRow, newCol)
	if err != nil || *str == "#" {
		return false
	}

	// matrix.Set(p.row, p.col, ".")

	p.row = newRow
	p.col = newCol

	// matrix.Set(p.row, p.col, "S")

	return true
}

func TestWalk(t *testing.T) {

	str := `#.#
.S.
#.E`

	matrix := NewMatrixFromString(str)

	fmt.Println(matrix.String())

	posS := getString("S", matrix)
	posE := getString("E", matrix)

	fmt.Println(posS, posE)

	path := path{row: posS[0], col: posS[1], vec: UP}
	path.walk(&matrix)

	fmt.Println(matrix.String())
	assert.Equal(t, matrix.String(), `#S#
...
#.E`)

	path.vec = DOWN
	path.walk(&matrix)

	path.vec = DOWN
	path.walk(&matrix)

	fmt.Println(matrix.String())

	assert.Equal(t, matrix.String(), `#.#
...
#SE`)

	path.vec = UP
	path.walk(&matrix)

	path.vec = LEFT
	path.walk(&matrix)

	fmt.Println(matrix.String())

	assert.Equal(t, matrix.String(), `#.#
S..
#.E`)

	path.vec = RIGHT
	path.walk(&matrix)

	path.vec = RIGHT
	path.walk(&matrix)

	path.vec = RIGHT
	path.walk(&matrix)

	fmt.Println(matrix.String())

	assert.Equal(t, matrix.String(), `#.#
..S
#.E`)

}

// func (p *path) up() *path {
// 	p.vec = UP
// 	return p
// }

// func (p *path) down() *path {
// 	p.vec = DOWN
// 	return p
// }

// func (p *path) left() *path {
// 	p.vec = LEFT
// 	return p
// }

// func (p *path) right() *path {
// 	p.vec = RIGHT
// 	return p
// }
