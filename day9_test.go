package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay9_1(t *testing.T) {

	str := "2333133121414131402"
	str, _ = loadInputFromFile("9")

	files := buildFiles(str)

	for i := 0; i < len(files); i++ {
		if files[i] == -1 {
			if last(files) == -1 {
				pop(&files, 1)
				i--
			} else {
				files[i] = last(files)
				pop(&files, 1)
			}
		}
	}

	sum := 0
	for i := 0; i < len(files); i++ {
		sum += files[i] * i
	}

	fmt.Println(sum)
	// assert.Equal(t, sum, 1928)
	assert.Equal(t, sum, 6446899523367)
}

// type block struct {
// 	// idx  int
// 	size int
// 	id   int
// }

func TestDay9_2(t *testing.T) {

	str := "2333133121414131402"
	// str, _ = loadInputFromFile("9")

	sum := 0

	b := buildFiles(str)

	fmt.Println(b)

	// files := make([]block, len(str))

	// for i := 0; i < len(str); i++ {

	// 	size := int(str[i] - 48)
	// 	files[i] = block{size: size, id: i / 2}

	// 	i++
	// 	if i == len(str) {
	// 		break
	// 	}
	// 	size = int(str[i] - 48)
	// 	files[i] = block{size: size, id: -1}
	// }

	// lastBlock := block{}

	// idx := len(files) - 1

	// for idx > 0 {
	// 	lastBlock = files[idx]
	// 	if lastBlock.id == -1 {
	// 		idx--
	// 		continue
	// 	}

	// 	fmt.Println(files)

	// 	for i := 0; i < len(files)-1; i++ {
	// 		// fmt.Println(files[i])
	// 		if files[i].id == -1 && files[i].size >= lastBlock.size {
	// 			fmt.Println(files)
	// 			remainingSize := files[i].size - lastBlock.size
	// 			files[i] = lastBlock

	// 			insert(&files, i, block{
	// 				size: remainingSize,
	// 				id:   -1,
	// 			})

	// 			popBlock(&files, 1)

	// 			fmt.Println(files)

	// 			break
	// 		}
	// 	}
	// 	idx--
	// }

	// fmt.Println(files)

	assert.Equal(t, sum, 2858)
}

// func insert(arr *[]block, afterIdx int, b block) {
// 	*arr = append((*arr)[:afterIdx+1], append([]block{b}, (*arr)[afterIdx+1:]...)...)
// }

func last(arr []int) int {
	return arr[len(arr)-1]
}

func pop(arr *[]int, n int) {
	*arr = (*arr)[:len(*arr)-n]
}

// func popBlock(arr *[]block, n int) { //TODO make files type with methd
// 	*arr = (*arr)[:len(*arr)-n]
// }

func buildFiles(str string) []int {
	idx := 0
	files := []int{}
	for i := 0; i < len(str); i += 2 {
		num := int(str[i] - 48)
		for j := 0; j < num; j++ {
			files = append(files, idx)
		}
		idx++

		if i+1 == len(str) {
			break
		}

		num = int(str[i+1] - 48)
		for j := 0; j < num; j++ {
			files = append(files, -1)
		}

	}

	return files
}
