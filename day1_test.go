package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay1_1_info(t *testing.T) {

	str := `3   4
4   3
2   5
1   3
3   9
3   3`

	ba := []byte(str)
	arr := getIntegers(ba)
	assert.Equal(t, len(arr), 12)

	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)/2)

	for i := 0; i < len(left); i++ {
		left[i] = arr[i*2]
		right[i] = arr[(i*2)+1]
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {

		sum += right[i] - left[i]

	}

	fmt.Println("sum", sum)
	assert.Equal(t, sum, 11)

}

func TestDay1_2_info(t *testing.T) {

	str := `3   4
4   3
2   5
1   3
3   9
3   3`

	ba := []byte(str)
	arr := getIntegers(ba)
	assert.Equal(t, len(arr), 12)

	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)/2)

	for i := 0; i < len(left); i++ {
		left[i] = arr[i*2]
		right[i] = arr[(i*2)+1]
	}

	numOfOcc := make(map[int]int)
	sum := 0

	for i := 0; i < len(left); i++ {
		idx := right[i]
		numOfOcc[idx]++
	}

	for _, i := range left {

		value := numOfOcc[i] * i
		fmt.Println("value", value)
		sum += value
	}

	fmt.Println("sum", sum)
	assert.Equal(t, sum, 31)

}

func TestDay1_1(t *testing.T) {

	str, _ := loadInput("1")
	ba := []byte(str)
	arr := getIntegers(ba)
	assert.Equal(t, len(arr), 2000)

	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)/2)

	for i := 0; i < len(left); i++ {
		left[i] = arr[i*2]
		right[i] = arr[(i*2)+1]
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {

		if left[i] >= right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}

	}

	fmt.Println("sum", sum)

}

func TestDay1_2(t *testing.T) {

	str, _ := loadInput("1")
	ba := []byte(str)
	arr := getIntegers(ba)
	assert.Equal(t, len(arr), 2000)

	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)/2)

	for i := 0; i < len(left); i++ {
		left[i] = arr[i*2]
		right[i] = arr[(i*2)+1]
	}
	numOfOcc := make(map[int]int)
	sum := 0

	for i := 0; i < len(left); i++ {
		idx := right[i]
		numOfOcc[idx]++
	}

	for _, i := range left {

		value := numOfOcc[i] * i
		fmt.Println("value", value)
		sum += value
	}

	fmt.Println("sum", sum)

}
