package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay11_1_info(t *testing.T) {

	str := `125 17`

	// str, _ = loadInputFromFile("11")

	c := processItems(str, 0) // 0 is blink 1
	assert.Equal(t, c, 4)

	c = processItems(str, 1)
	assert.Equal(t, c, 5)

	c = processItems(str, 2)
	assert.Equal(t, c, 9)

	c = processItems(str, 3)
	assert.Equal(t, c, 13)

	c = processItems(str, 4) // after blink 6
	assert.Equal(t, c, 22)

	c = processItems(str, 23) // after blink 25
	assert.Equal(t, c, 55312)

}

// WORKING
func TestDay11_1(t *testing.T) {

	str, _ := loadInputFromFile("11")

	maxLevel := 23 // after blink 25

	c := processItems(str, maxLevel)

	assert.Equal(t, c, 187738)

}

func TestDay11_1__(t *testing.T) {

	str := "0"

	for i := -2; i < 25; i++ {
		maxLevel := i // after blink 25
		c := processItems(str, maxLevel)
		fmt.Println(i+2, c)
		if i+2 == 20 {
			assert.Equal(t, c, 2377)
		}

	}
}

func TestDay11_1___(t *testing.T) {

	str := "0"

	c := processItems(str, 48)

	fmt.Println(c)

	// 396 0 at level 25
	// 19778 all at level 25
	// 396*396 0 at level 50
	// 396*19778 = 7832088 all at level 50

	// 14967989 0's at level50
	// measured: 663251546 all from top 0

	// 396*396*396 at level 75
	// 396*396*19778 at level 75

	// day input level 25 4291 0
	// 4291 * 4291 * 19778 = 36401870337 at level 75

	// 46.417.536 0 at level 75
	// 3.101.506.848 at level 75 from 1 top 0

}

func TestDay11_2(t *testing.T) {

	str := "77"

	maxLevel := 23 // after blink 75

	c := processItems(str, maxLevel)

	fmt.Println(c)
	// assert.Equal(t, c, 187738)

}

func BenchmarkDay11_2(b *testing.B) {

	str, _ := loadInputFromFile("11")

	maxLevel := 31 // after blink 75

	c := processItems(str, maxLevel)

	fmt.Println(c)
}

func processItems(str string, maxLevel int) int {
	items := strings.Split(str, " ")

	stack := []Item{}

	for _, i := range items {
		stack = append(stack, Item{val: i, level: 0})
	}

	count := 0
	cZero := 0

	for len(stack) > 0 {

		// for _, item := range items {
		item := stack[len(stack)-1]

		//pop stack
		stack = stack[:len(stack)-1]

		for _, s := range handleStr(item.val) {

			if item.level < maxLevel {
				i := Item{val: s, level: item.level + 1}
				stack = append(stack, i)
			} else {

				count += len(handleStr(s))
				if item.val == "0" {
					cZero++
				}
			}
		}

	}

	fmt.Println(cZero)
	fmt.Println(count * cZero)

	return count

}

func handleStr(str string) []string {

	if str == "0" {
		return []string{"1"}
	}
	if len(str)%2 == 0 {
		return splitStringInHalves(str)
	}

	num, _ := strconv.Atoi(str)
	return []string{fmt.Sprint(num * 2024)}

}

func splitStringInHalves(str string) []string {

	str1 := str[:len(str)/2]
	str2 := str[len(str)/2:]

	// num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)

	return []string{str1, fmt.Sprint(num2)}
}

type Item struct {
	val   string
	level int
}
