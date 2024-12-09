package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay8_1(t *testing.T) {

	str := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

	// find all locations

	// lines := strings.split(str, "\n")
	// lines := strings.Split(str, "\n"

	str, _ = loadInputFromFile("8")

	matrix := stringToMatrix(str)

	antennas := map[string][]antenna{}

	for i, l := range matrix {
		for j, c := range l {
			if c != "." {
				antennas[string(c)] = append(antennas[string(c)], antenna{typ: string(c), row: i, col: j})
			}
		}
	}

	for _, v := range antennas {
		for i, p := range v {
			as := iterOthers(v, i)
			for _, a := range as {

				delta := caldDelte(p, a)
				fmt.Println(p, a, delta)
				r := p.row + (delta[0] * 2)
				c := p.col + (delta[1] * 2)

				if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0]) {
					// if matrix[r][c] == "." {
					matrix[r][c] = "#"
					// }
				}
			}
			fmt.Println(p, i, as)
		}
	}
	fmt.Println(matrix)
	count := countOccurances(matrix, "#")
	fmt.Println(count)
	assert.Equal(t, count, 14)

}

type antenna struct {
	typ string
	row int
	col int
}

func caldDelte(a, b antenna) []int {
	return []int{b.row - a.row, b.col - a.col}
}

func iterOthers(arr []antenna, idx int) []antenna {
	res := make([]antenna, 0, len(arr)-1)
	res = append(res, arr[:idx]...)
	if idx < len(arr)-1 {
		res = append(res, arr[idx+1:]...)
	}
	return res
}
