package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay5_1(t *testing.T) {

	str := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	str, _ = loadInputFromFile("5")

	parts := strings.Split(str, "\n\n")

	_rules := strings.Split(parts[0], "\n")

	// convert rules to ints
	rules, _ := arrIntsFromString(_rules, "|")

	_order := strings.Split(parts[1], "\n")

	order, _ := arrIntsFromString(_order, ",")

	isCorrect := true

	result := 0

	for _, o := range order {

		for _, r := range rules {
			v1 := getIndexInt(o, r[0])
			v2 := getIndexInt(o, r[1])

			if v1 == -1 || v2 == -1 {
				continue
			}

			if v1 > v2 {
				fmt.Println("error")
				isCorrect = false
				break
			}
		}
		if isCorrect {
			m := getMidValue(o)
			result += m
		}
		isCorrect = true
	}

	fmt.Println(result)

	// assert.Equal(t, result, 143)
	assert.Equal(t, result, 7074)

}

func TestDay5_2(t *testing.T) {

	str := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	str, _ = loadInputFromFile("5")

	parts := strings.Split(str, "\n\n")

	_rules := strings.Split(parts[0], "\n")

	// convert rules to ints
	rules, _ := arrIntsFromString(_rules, "|")

	_order := strings.Split(parts[1], "\n")

	order, _ := arrIntsFromString(_order, ",")

	result := 0

	for _, o := range order {

		isCorrect := true

		for i := 0; i < len(rules); i++ {
			r := rules[i]
			idx1 := getIndexInt(o, r[0])
			idx2 := getIndexInt(o, r[1])

			if idx1 == -1 || idx2 == -1 {
				continue
			}

			if idx1 > idx2 {
				isCorrect = false
				swap(o, idx1, idx2)
				i = 0 // reset the loop
			}

		}

		if !isCorrect {
			result += getMidValue(o)
			isCorrect = true
		}
	}

	fmt.Println(result)

	// assert.Equal(t, result, 123)
	assert.Equal(t, result, 4828)

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func getIndexInt(arr []int, val int) int {

	for i, a := range arr {
		if a == val {
			return i
		}
	}
	return -1
}

func arrIntsFromString(input []string, delim string) ([][]int, error) {
	var arr [][]int

	for _, str := range input {
		parts := strings.Split(str, delim)
		ints := make([]int, len(parts)) // Preallocate slice
		for idx, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error converting string %q to int in %q: %w", part, str, err)
			}
			ints[idx] = num
		}
		arr = append(arr, ints)
	}

	return arr, nil
}

func getMidValue(arr []int) int {

	// gets the mid value of an array
	// if the array is even, return the first of the middle values
	// if the array is odd, return the mid value

	if len(arr)%2 == 0 {
		return arr[len(arr)/2]
	}

	return arr[len(arr)/2]
}
