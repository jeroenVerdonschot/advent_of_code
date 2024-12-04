package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func stringToMatrix(str string) [][]string {
	rows := strings.Split(str, "\n")
	matrix := make([][]string, len(rows))
	for i := 0; i < len(rows); i++ {
		col := strings.Split(rows[i], "")
		for j := 0; j < len(col); j++ {
			col[j] = strings.TrimSpace(col[j])
			if col[j] != "" {
				matrix[i] = append(matrix[i], col[j])
			}
		}
	}
	return matrix
}

func TestDay4_2_info(t *testing.T) {

	mask := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	str := `MMMSXXMASM
    MSAMXMSMSA
    AMXSXMAAMM
    MSAMASMSMX
    XMASAMXAMM
    XXAMMXXAMA
    SMSMSASXSS
    SAXAMASAAA
    MAMMMXMMMM
    MXMXAXMASX`

	matrix := stringToMatrix(str)

	iterMaskStringMatrix(matrix, mask)
	arr := iterMaskStringMatrix(matrix, mask)

	count := 0
	for _, v := range arr {
		if v == "MMASS" ||
			v == "SSAMM" ||
			v == "MSAMS" ||
			v == "SMASM" {

			{
				fmt.Println("Found")
				count++
			}
		}
	}
	fmt.Println(count)
}

func TestDay4_2(t *testing.T) {

	mask := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	str, _ := loadInputFromFile("4")

	matrix := stringToMatrix(str)

	iterMaskStringMatrix(matrix, mask)
	arr := iterMaskStringMatrix(matrix, mask)

	count := 0
	for _, v := range arr {
		if v == "MMASS" ||
			v == "SSAMM" ||
			v == "MSAMS" ||
			v == "SMASM" {

			{
				fmt.Println("Found")
				count++
			}
		}
	}
	fmt.Println(count)
}
func TestMaskStringMatrix(t *testing.T) {

	str := `m2m4
    5a78
    s0s2
    e4e6
    r8r0`

	matrix := stringToMatrix(str)

	fmt.Println(matrix)

	// Define a smaller mask matrix (still integers to indicate mask positions)
	mask := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1},
	}

	arr := iterMaskStringMatrix(matrix, mask)

	for _, v := range arr {
		if v == "mmass" ||
			v == "SSAMM" {
			{
				fmt.Println("Found")
			}
		}
	}
}

func TestDay4_1_mask(t *testing.T) {

	str := `MMMSXXMASM
    MSAMXMSMSA
    AMXSXMAAMM
    MSAMASMSMX
    XMASAMXAMM
    XXAMMXXAMA
    SMSMSASXSS
    SAXAMASAAA
    MAMMMXMMMM
    MXMXAXMASX`

	str, _ = loadInputFromFile("4")
	matrix := stringToMatrix(str)
	count := 0

	masks := [][][]int{
		{{1, 1, 1, 1}},
		{{1}, {1}, {1}, {1}},
		{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
		{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}},
	}

	for _, mask := range masks {
		arr := iterMaskStringMatrix(matrix, mask)
		for _, v := range arr {
			if v == "XMAS" ||
				v == "SAMX" {
				{
					count++
				}
			}
		}
	}

	// assert.Equal(t, count, 18)
	assert.Equal(t, count, 2500)
}

func iterMaskStringMatrix(matrix [][]string, mask [][]int) []string {
	maskRows := len(mask)
	maskCols := len(mask[0])
	maxI := len(matrix) - maskRows
	maxJ := len(matrix[0]) - maskCols

	values := []string{}

	for i := 0; i <= maxI; i++ {
		for j := 0; j <= maxJ; j++ {
			affectedValues := maskStringMatrix(matrix, mask, i, j)
			values = append(values, affectedValues)
		}
	}
	return values
}

func maskStringMatrix(matrix [][]string, mask [][]int, startX, startY int) string {
	maskRows := len(mask)
	maskCols := len(mask[0])
	affectedValues := ""

	for i := 0; i < maskRows; i++ {
		for j := 0; j < maskCols; j++ {
			matrixX := startX + i
			matrixY := startY + j
			if mask[i][j] != 0 {
				affectedValues += matrix[matrixX][matrixY]
			}
		}
	}
	return affectedValues
}
