package main

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestLoadInput(t *testing.T) {

	str, err := loadInput("1")
	assert.NoError(t, err)

	fmt.Println(str)

}

func TestDeltas(t *testing.T) {

	row := []int{1, 2, 10, 4, 5}

	deltas := deltas(row)

	fmt.Println(deltas)
	assert.Equal(t, deltas, []int{1, 8, -6, 1})

}

func TestAllPositive(t *testing.T) {

	row := []int{1, 2, 10, 4, 5}

	result := allPositive(row)

	fmt.Println(result)
	assert.Equal(t, result, true)

	row = []int{1, 2, -10, 4, 5}

	result = allPositive(row)

	fmt.Println(result)
	assert.Equal(t, result, false)

}

func TestAllSameSigned(t *testing.T) {

	row := []int{1, 2, 10, 4, 5}

	result := allSameSigned(row)

	fmt.Println(result)
	assert.Equal(t, result, true)

	row = []int{1, 2, -10, 4, 5}

	result = allSameSigned(row)

	fmt.Println(result)
	assert.Equal(t, result, false)

}

func TestMaxInt(t *testing.T) {

	row := []int{1, 2, 10, 4, 5}

	result := maxInt(row)

	fmt.Println(result)
	assert.Equal(t, result, 10)

}

func TestMinInt(t *testing.T) {

	row := []int{1, 2, 10, 4, 5}

	result := minInt(row)

	fmt.Println(result)
	assert.Equal(t, result, 1)

}
