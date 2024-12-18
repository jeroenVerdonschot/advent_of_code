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
		pos.walk(matrix, dir, "X")

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

	dir := Direction{-1, 0}

	startPos := findStrInMatrix(matrix, "^")

	// record path
	path := []Pos{startPos}

	for {
		p := startPos.walk(matrix, dir, "X")
		path = append(path, p...)
		if startPos.row == -1 {
			break
		}
		dir.rotateRight()
	}

	startDir := Direction{-1, 0}

	for i := 1; i < len(path); i++ {

		matrix[path[i].row][path[i].col] = "#"

		if makeTestWalk(startPos, matrix, startDir) {
			fmt.Println("Loop detected")
		}

		matrix[path[i].row][path[i].col] = "."

	}
	// assert.Equal(t, obstructions, 6)

}

func makeTestWalk(pos Pos, matrix [][]string, dir Direction) bool {

	escape := 0
	for {

		pos.walk(matrix, dir, "O")
		if pos.row == -1 {
			return false
		}

		// check for loop
		// if matrix[pos.row][pos.col] == "p" {
		// 	fmt.Println("Loop detected")
		// 	return true
		// }

		dir.rotateRight()

		escape++

		if escape > 10000 {
			break
		}
	}
	return true
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

func (pos *Pos) walk(matrix [][]string, dir Direction, marker string) []Pos {

	var str string
	path := []Pos{}

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

		matrix[pos.row][pos.col] = marker
		path = append(path, Pos{pos.row, pos.col})

	}
	return path
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
