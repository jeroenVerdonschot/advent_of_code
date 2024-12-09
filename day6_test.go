package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay6_1(t *testing.T) {

	str := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	str, _ = loadInputFromFile("6")
	matrix := stringToMatrix(str)

	dir := Direction{-1, 0}

	pos := findStrInMatrix(matrix, "^")
	matrix[pos.row][pos.col] = "X"

	totalSteps := 0
	fmt.Println(pos)
	for {
		pos.walk(matrix, dir)

		if pos.row == -1 {
			break
		}
		dir.rotateRight()
	}

	totalSteps = countOccurances(matrix, "X")

	fmt.Println(totalSteps)

	assert.Equal(t, totalSteps, 5331)
	// fmt.Println(matrix)
}

func TestDay6_2(t *testing.T) {

	str := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	// str, _ = loadInputFromFile("6")
	matrix := stringToMatrix(str)

	initDir := Direction{-1, 0}
	dir := initDir

	var p path

	initPos := findStrInMatrix(matrix, "^")
	p.curPos = initPos
	p.add(p.curPos)

	obstructions := 0

	fmt.Println(p.curPos)

mainLoop1:
	for {
		for {
			p.walk(matrix, dir) // TO DO one step et a time
			p.add(p.curPos)
			if p.isOutOfBound {
				break mainLoop1 // breaks out of outer loop
			}
			if p.hasHitMarker {
				break // breaks inner loop only
			}
		}
		dir.rotateRight()
	}

	for i := 2; i < len(p.positions); i++ {

		matrixAlt := matrix
		p.curPos = initPos
		dir = initDir

		matrixAlt[p.positions[i].row][p.positions[i].col] = "#"

	mainloop2:
		for {

			for {
				p.walk(matrix, dir)

				// if pos & dir are init then loop
				if p.curPos == initPos && dir == initDir {
					fmt.Println("found loop", p.positions[i])
					obstructions++
				}

				if p.isOutOfBound {
					fmt.Println("out of bounds")
					break mainloop2
				}
				if p.hasHitMarker {
					break // breaks inner loop only
				}
			}

			dir.rotateRight()
		}
	}

	fmt.Println(obstructions)

	assert.Equal(t, obstructions, 6)
	// fmt.Println(matrix)

	// record Path
	// brute force every point on path

}

type path struct {
	isOutOfBound bool
	hasHitMarker bool
	curPos       Pos
	positions    []Pos
}

func (p *path) walk(matrix [][]string, dir Direction) {

	var str string
	p.isOutOfBound = false
	p.hasHitMarker = false

	if p.curPos.row < 1 || p.curPos.col < 1 || p.curPos.row >= len(matrix)-1 || p.curPos.col >= len(matrix[0])-1 {
		p.isOutOfBound = true
		return
	}

	str = matrix[p.curPos.row+dir.row][p.curPos.col+dir.col]

	if str == "#" {
		p.hasHitMarker = true
		return
	}

	p.curPos.row += dir.row
	p.curPos.col += dir.col

}

func (p *path) add(pos Pos) {
	if len(p.positions) != 0 {
		if p.curPos != p.positions[len(p.positions)-1] {
			p.positions = append(p.positions, pos)
		}

	} else {
		p.positions = append(p.positions, pos)
	}
}

func countOccurances(matrix [][]string, str string) int {
	count := 0
	for _, row := range matrix {
		for _, col := range row {
			if col == str {
				count++
			}
		}
	}
	return count
}

type Pos struct {
	row int
	col int
}

type Direction struct {
	row int
	col int
}

func (dir *Direction) rotateRight() {
	switch {
	case dir.row == -1 && dir.col == 0:
		dir.row = 0
		dir.col = 1
	case dir.row == 0 && dir.col == 1:
		dir.row = 1
		dir.col = 0
	case dir.row == 1 && dir.col == 0:
		dir.row = 0
		dir.col = -1
	case dir.row == 0 && dir.col == -1:
		dir.row = -1
		dir.col = 0
	}
}

func (pos *Pos) walk(matrix [][]string, dir Direction) {

	var str string

	for {

		if pos.row < 1 || pos.col < 1 || pos.row >= len(matrix)-1 || pos.col >= len(matrix[0])-1 {
			pos.row = -1
			break
		}

		str = matrix[pos.row+dir.row][pos.col+dir.col]

		if str == "#" {
			break
		}

		pos.row += dir.row
		pos.col += dir.col

		matrix[pos.row][pos.col] = "X"

	}

}

func (pos *Pos) peek(matrix [][]string, dir Direction) {

	var str string

	for {

		if pos.row < 1 || pos.col < 1 || pos.row >= len(matrix)-1 || pos.col >= len(matrix[0])-1 {
			pos.row = -1
			break
		}

		str = matrix[pos.row+dir.row][pos.col+dir.col]

		if str == "#" {
			break
		}

		pos.row += dir.row
		pos.col += dir.col

		matrix[pos.row][pos.col] = "p"

	}

}

func findStrInMatrix(matrix [][]string, str string) Pos {
	for r, row := range matrix {
		for c, col := range row {
			if col == str {
				return Pos{r, c}
			}
		}
	}
	return Pos{}
}
