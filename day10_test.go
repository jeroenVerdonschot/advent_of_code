package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay10_1(t *testing.T) {

	str := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

	str, _ = loadInputFromFile("10")

	matrix := stringToMatrix(str)

	gt := newGridTree(len(matrix[0]), len(matrix))

	gt.buildGrid(matrix)

	gt.linkNodes()

	count, distinct := gt.findPaths()

	fmt.Println(count, distinct)

	assert.Equal(t, count, 659)
	assert.Equal(t, distinct, 1463)

}

func (gt *gridTree) linkNodes() {
	for i, row := range gt.grid {
		for j := range row {
			if isNumber(gt.grid[i][j].val) {
				gt.linkNeighbours(i, j)
			}
			if gt.grid[i][j].val == 0 {
				gt.roots = append(gt.roots, gt.grid[i][j])
			}

		}
	}
}

func (gt *gridTree) findPaths() (int, int) {

	count := 0
	distinct := 0

	for _, root := range gt.roots {

		st := stack{}
		st = append(st, root)

		for len(st) > 0 {

			current := st[len(st)-1]
			st = st[:len(st)-1]

			if len(current.highers) > 0 {
				st = append(st, current.highers...)
			}
			if current.val == 9 {
				gt.addUniqueEndNode(current)
				distinct++
			}
		}

		// fmt.Println("distinct", distinct)
		count += len(gt.endPoints)
		// fmt.Println("endPoints", len(gt.endPoints))
		gt.endPoints = []*pathNode{}

	}
	return count, distinct
}

func (gt *gridTree) buildGrid(matrix [][]string) {

	for i, row := range matrix {
		for j, col := range row {
			val := -1
			if col != "." {
				val, _ = strconv.Atoi(col)
			}
			pn := &pathNode{row: i, col: j, val: val}
			gt.addGridNode(pn)
		}
	}
}

func isNumber(s int) bool {
	return s >= 0 && s <= 9
}

func newGridTree(width, height int) *gridTree {

	gt := &gridTree{
		grid: make([][]*pathNode, height),
	}

	for i := range gt.grid {
		gt.grid[i] = make([]*pathNode, width)
	}

	return gt
}

func (g *gridTree) linkNeighbours(i, j int) {
	grid := g.grid
	node := grid[i][j]
	// top
	if i > 0 {
		g.linkIfNumber(node, grid[i-1][j])
	}

	// right
	if j < len(grid[0])-1 {
		g.linkIfNumber(node, grid[i][j+1])
	}

	// bottom
	if i < len(grid)-1 {
		g.linkIfNumber(node, grid[i+1][j])
	}

	// left
	if j > 0 {
		g.linkIfNumber(node, grid[i][j-1])
	}
}

func (g *gridTree) addUniqueEndNode(p *pathNode) {
	// check if it is already in the list
	for _, end := range g.endPoints {
		if end == p {
			return
		}
	}
	g.endPoints = append(g.endPoints, p)
}

func (g *gridTree) linkIfNumber(node, neighbor *pathNode) {
	if isNumber(neighbor.val) {
		// if node.val > neighbor.val {
		// 	node.lowers = append(node.lowers, neighbor)
		// }
		if node.val+1 == neighbor.val {
			node.highers = append(node.highers, neighbor)
		}
	}
}

type gridTree struct {
	roots     []*pathNode
	grid      [][]*pathNode
	endPoints []*pathNode
}

func (g *gridTree) addGridNode(p *pathNode) {
	g.grid[p.row][p.col] = p
}

type pathNode struct {
	row int
	col int
	val int
	// lowers  []*pathNode
	highers []*pathNode
}

type stack []*pathNode
