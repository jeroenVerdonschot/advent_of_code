package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

var UP = []int{-1, 0}
var DOWN = []int{1, 0}
var LEFT = []int{0, -1}
var RIGHT = []int{0, 1}

func TestDay15_1_countResult(t *testing.T) {

	str := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`

	matrix := NewMatrixFromString(str)

	sum := countResult(matrix)

	fmt.Println(matrix.String())

	assert.Equal(t, sum, 10092)

	str = `########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`
	matrix = NewMatrixFromString(str)

	sum = countResult(matrix)

	fmt.Println(matrix.String())

	assert.Equal(t, sum, 2028)

}

func TestDay15_1_large(t *testing.T) {

	str := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	split := strings.Split(str, "\n\n")
	matrix := NewMatrixFromString(split[0])

	// get the robot
	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	fmt.Println("robot: ", r)
	fmt.Println(matrix.String())

	cmds := readCmds(split[1])

	for _, cmd := range cmds {
		r.move(cmd, matrix)
	}

	sum := countResult(matrix)

	fmt.Println("sum: ", sum)
	assert.Equal(t, sum, 10092)

}

func TestDay15_1_small(t *testing.T) {

	str := `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

	split := strings.Split(str, "\n\n")
	matrix := NewMatrixFromString(split[0])

	// get the robot
	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	fmt.Println("robot: ", r)
	fmt.Println(matrix.String())

	cmds := readCmds(split[1])

	for _, cmd := range cmds {
		r.move(cmd, matrix)
	}

	sum := countResult(matrix)

	fmt.Println("sum: ", sum)
	assert.Equal(t, sum, 2028)

}

func TestDay15_1(t *testing.T) {

	str, _ := loadInputFromFile("15")

	split := strings.Split(str, "\n\n")
	matrix := NewMatrixFromString(split[0])

	// get the robot
	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	fmt.Println("robot: ", r)
	fmt.Println(matrix.String())

	cmds := readCmds(split[1])

	for _, cmd := range cmds {
		r.move(cmd, matrix)
	}

	sum := countResult(matrix)

	fmt.Println("sum: ", sum)
	assert.Equal(t, sum, 1509074)

}

// row * 100 + col
func countResult(matrix Matrix[string]) int {

	sum := 0

	for row := 0; row < matrix.rows; row++ {
		for col := 0; col < matrix.cols; col++ {
			if *matrix.data[row][col] == "O" {
				sum += row*100 + col
			}
		}
	}

	return sum
}

type robot struct {
	row, col int
}

// move robot
func (r *robot) move(vec []int, matrix Matrix[string]) {

	nextRow := r.row + vec[0]
	nextCol := r.col + vec[1]

	strPtr, err := matrix.Get(nextRow, nextCol)
	if err != nil {
		fmt.Println("can't move")
		return
	}

	if *strPtr == "#" {
		return
	}

	if *strPtr == "O" {
		canMove := pushBox(nextRow, nextCol, vec, matrix)
		if !canMove {
			return
		}
	}

	// move robot
	moveItem("@", r.row, r.col, vec, matrix)
	r.row += vec[0]
	r.col += vec[1]
}

func pushBox(row, col int, vec []int, matrix Matrix[string]) bool {

	nextRow := row + vec[0]
	nextCol := col + vec[1]

	strPtr, err := matrix.Get(nextRow, nextCol)
	if err != nil || *strPtr == "#" {
		return false
	}

	if *strPtr == "O" {
		val := pushBox(nextRow, nextCol, vec, matrix)
		if !val {
			return false
		}
	}

	// move box
	err = moveItem("O", row, col, vec, matrix)

	return err == nil
}

func moveItem(str string, row, col int, vec []int, matrix Matrix[string]) error {
	err := matrix.Set(row, col, ".")
	if err != nil {
		return err
	}
	err = matrix.Set(row+vec[0], col+vec[1], str)
	if err != nil {
		return err
	}
	return nil
}

func getString(str string, matrix Matrix[string]) []int {
	for row := 0; row < matrix.rows; row++ {
		for col := 0; col < matrix.cols; col++ {
			if *matrix.data[row][col] == str {
				return []int{row, col}
			}
		}
	}
	return []int{}
}

func readCmds(str string) [][]int {
	cmds := [][]int{}
	for _, c := range str {

		switch c {
		case '>':
			cmds = append(cmds, RIGHT)
		case '<':
			cmds = append(cmds, LEFT)
		case '^':
			cmds = append(cmds, UP)
		case 'v':
			cmds = append(cmds, DOWN)
		}

	}
	return cmds
}

func TestReadCmds(t *testing.T) {

	str := `<>^v`

	cmds := readCmds(str)

	fmt.Println(cmds)

}
func TestMoveRobot(t *testing.T) {

	str := `...
.@.
...`

	matrix := NewMatrixFromString(str)

	// get the robot
	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	assert.Equal(t, r.row, 1)
	assert.Equal(t, r.col, 1)

	r.move(UP, matrix)
	strRet := matrix.String()

	assert.Equal(t, strRet, `.@.
...
...`)

	r.move(RIGHT, matrix)
	strRet = matrix.String()
	assert.Equal(t, strRet, `..@
...
...`)

	r.move(DOWN, matrix)
	strRet = matrix.String()
	assert.Equal(t, strRet, `...
..@
...`)

	r.move(LEFT, matrix)
	strRet = matrix.String()
	assert.Equal(t, strRet, `...
.@.
...`)
}

func TestPushBox(t *testing.T) {

	str := `...
@.#
...`
	matrix := NewMatrixFromString(str)

	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	r.move(LEFT, matrix)
	strRet := matrix.String()

	fmt.Println(strRet)

	assert.Equal(t, strRet, `...
@.#
...`)
	r.move(RIGHT, matrix)
	r.move(RIGHT, matrix)
	r.move(RIGHT, matrix)
	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `...
.@#
...`)
}

func TestPushBox_move(t *testing.T) {

	str := `....
@O.#
....`
	matrix := NewMatrixFromString(str)

	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	r.move(RIGHT, matrix)
	strRet := matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	r.move(RIGHT, matrix)
	// r.move(RIGHT, matrix)
	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `....
.@O#
....`)
}
func TestPushBox_move2(t *testing.T) {

	str := `.....
@O..#
.....`
	matrix := NewMatrixFromString(str)

	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	r.move(RIGHT, matrix)
	strRet := matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	r.move(RIGHT, matrix)
	// r.move(RIGHT, matrix)
	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `.....
..@O#
.....`)
}

func TestPushBox_move3(t *testing.T) {

	str := `......
@OO..#
......`
	matrix := NewMatrixFromString(str)

	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]

	r.move(RIGHT, matrix)
	strRet := matrix.String()

	fmt.Println("")
	fmt.Println(strRet)
	assert.Equal(t, strRet, `......
.@OO.#
......`)

	r.move(RIGHT, matrix)
	// r.move(RIGHT, matrix)
	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `......
..@OO#
......`)
	r.move(RIGHT, matrix)
	// r.move(RIGHT, matrix)
	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `......
..@OO#
......`)
}

func TestPushBox_move4(t *testing.T) {

	str := `...
.@.
.O.
.O.
...
...
.#.`
	matrix := NewMatrixFromString(str)
	var r robot
	pos := getString("@", matrix)
	r.row, r.col = pos[0], pos[1]
	r.move(DOWN, matrix)

	strRet := matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `...
...
.@.
.O.
.O.
...
.#.`)
	r.move(DOWN, matrix)

	strRet = matrix.String()

	fmt.Println("")
	fmt.Println(strRet)

	assert.Equal(t, strRet, `...
...
...
.@.
.O.
.O.
.#.`)
}
