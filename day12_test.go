package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay12_1_detect_islands(t *testing.T) {

	str := `AAAA
BBCD
BBCC
EEEC`

	str = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

	str = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	str, _ = loadInputFromFile("12")

	matrix := stringToMatrix(str)

	world := newField(matrix)

	isFinished := false

	for !isFinished {
		isFinished = world.detectIslands()
	}

	sum := 0
	for _, region := range world.regions {

		// fmt.Println("region: ", region.id, "value: ", region.value, "pixels: ", len(region.pixels))

		perimeter := len(region.pixels) * 4
		for _, pixel := range region.pixels {
			perimeter -= len(pixel.neighbours)
		}
		fmt.Println("total: ", perimeter*len(region.pixels), "perimeter: ", perimeter, "area: ", len(region.pixels))
		sum += perimeter * len(region.pixels)
	}

	fmt.Println("sum: ", sum)
	fmt.Println("ready")

	//assert.Equal(t, sum, 140)
	// assert.Equal(t, sum, 772)
	assert.Equal(t, sum, 1930)

}

func (w *field) detectIslands() bool {
	regionId := 0
	for i := 0; i < w.cols; i++ {
		for j := 0; j < w.rows; j++ {

			if l, err := w.Matrix.Get(i, j); err == nil && l.regionId == 0 {
				regionId++
				region := &region{id: regionId, value: l.value}

				w.regions = append(w.regions, region)

				fmt.Println("new region: ", regionId, l.value)

				w.floodRegion(l, region) // with delete
			}
		}
	}
	return true
}

func (f *field) floodRegion(startPixel *pixel, r *region) {

	stack := []*pixel{startPixel}

	// directions slice containing functions to get adjacent pixels
	directions := []func(row, col int) (*pixel, error){
		f.GetLeft,
		f.GetUp,
		f.GetRight,
		f.GetDown,
	}

	for len(stack) > 0 {
		curPixel := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// mark the pixel as visited

		if curPixel.regionId == 0 {
			r.pixels = append(r.pixels, curPixel)
		} else {
			continue
		}

		// fmt.Println("floodRegion: ", curPixel.row, curPixel.col, "region: ", curPixel.regionId)

		// remove the pixel from the matrix
		curPixel.regionId = r.id

		for _, getAdjacent := range directions {

			if adjacent, err := getAdjacent(curPixel.row, curPixel.col); err == nil &&
				adjacent.value == r.value {

				if !hasNeighbour(curPixel, adjacent) {

					// fmt.Println("linking: ", curPixel.row, curPixel.col, "to", adjacent.row, adjacent.col)
					// link the pixels
					curPixel.neighbours = append(curPixel.neighbours, adjacent)
					adjacent.neighbours = append(adjacent.neighbours, curPixel)

					// add to stack
					stack = append(stack, adjacent)
				}

			}
		}

	}
}

func hasNeighbour(p1, p2 *pixel) bool {
	for _, n := range p1.neighbours {
		if n == p2 {
			return true
		}
	}
	return false
}

func newField(matrix [][]string) field {

	_matrix := NewMatrix(len(matrix), len(matrix[0]), pixel{})

	for i, row := range matrix {
		for j, col := range row {
			_matrix.Set(i, j, pixel{value: col, row: i, col: j})
		}
	}

	return field{Matrix: _matrix}
}

type field struct {
	*Matrix[pixel]
	regions []*region
}

type region struct {
	pixels []*pixel
	id     int
	value  string
}

type pixel struct {
	neighbours []*pixel
	value      string
	row        int
	col        int
	regionId   int
}
