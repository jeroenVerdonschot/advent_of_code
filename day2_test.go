package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay2_1_info(t *testing.T) {

	str := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	matrix := makeMatrixOfInts(str)

	safe := scan(matrix)

	assert.Equal(t, safe, 2)

}

func TestDay2_2_info(t *testing.T) {

	str := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	matrix := makeMatrixOfInts(str)

	safe := scanPart2(matrix)
	assert.Equal(t, safe, 4)

}

func TestDay2_3_info(t *testing.T) {

	str := `44 37 34 31 30 27 25
	44 37 37 31 30 27 25
	1 2 3 4 5 6
	1 3 2 4 5
	1 3 6 7 9
	1 1 1 1 1 1 1
	68 74 71 73 76 79 82 87
	68 74 71 73 76 79 82 84
	`

	matrix := makeMatrixOfInts(str)
	matrix = matrix[:len(matrix)-1]

	safe := scanPart2(matrix)
	assert.Equal(t, safe, 5)
}

func scan(matrix Maxtrix) int {

	safe := len(matrix)

	for _, row := range matrix {
		if v, _ := validateDay2(row); !v {
			safe--
		}
	}
	return safe
}

// ugly solution but it works
func scanPart2(matrix Maxtrix) int {

	valid := 0

	for _, row := range matrix {

		v, _ := validateDay2(row)
		if !v {
			for idx := 0; idx < len(row); idx++ {
				r := row.remove(idx)
				v2, _ := validateDay2(r)
				if v2 {
					valid++
					break
				}
			}
		} else {
			valid++
		}
	}
	return valid
}

func validate(row []int, isInc bool, i int, j int) (bool, int) {
	if row[i] == row[j] {
		// fmt.Println("equal", row, i)
		return false, i
	}

	if row[i] > row[j] && isInc {
		// fmt.Println("not inc", row, i)
		return false, i
	}

	if row[i] < row[j] && !isInc {
		// fmt.Println("not dec", row, i)
		return false, i
	}

	if row[j]-row[i] > 3 || row[i]-row[j] > 3 {
		// fmt.Println(">3", row, i)
		return false, i
	}
	return true, 0
}

func validateDay2(row []int) (bool, int) {

	isInc := row[0] < row[1]
	// isDec := row[0] > row[1]

	for i := 0; i < len(row)-1; i++ {
		v, j := validate(row, isInc, i, i+1)
		if !v {
			return false, j
		}
	}
	return true, 0
}
func TestDay2_1(t *testing.T) {

	str, _ := loadInput("2")

	matrix := makeMatrixOfInts(str)
	matrix = matrix[:len(matrix)-1]

	safe := scan(matrix)

	fmt.Println("safe:", safe)
	assert.Equal(t, safe, 534)

}

func TestDay2_2(t *testing.T) {

	// str, _ := loadInput("2")
	str, _ := loadInputFromFile("2")

	matrix := makeMatrixOfInts(str)

	safe := scanPart2(matrix)

	fmt.Println("safe:", safe)
	assert.Equal(t, safe, 577)
}

func TestRemoveIntFromRow(t *testing.T) {

	row := Row{1, 2, 3, 4, 5}
	r := row.remove(2)
	assert.Equal(t, r, Row{1, 2, 4, 5})
}
