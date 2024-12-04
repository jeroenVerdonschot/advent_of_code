package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay3_1_info(t *testing.T) {

	str := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	// str := `mul(11,8)mul(8,-5)`

	lexer := Lexer{input: str, pos: -1}
	result := 0

	for lexer.next() != "EOF" {
		if digitX, digitY, ok := lexer.mulExpression(); ok {
			fmt.Println(digitX, digitY)
			result += digitX * digitY
		}
	}

	assert.Equal(t, result, 161)
}

func TestDay3_2_info(t *testing.T) {

	// str := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
	str := `xmul(2,4)&mul[3,7]!^don't()mul(5,5)+mul(32,64](mul(11,8)undo()mul(8,5))`
	// str := `mul(11,8)mul(8,-5)`

	lexer := Lexer{input: str, pos: -1}
	result := 0

	do := true

	for lexer.next() != "EOF" {
		if digitX, digitY, ok := lexer.mulExpression(); ok {
			fmt.Println(digitX, digitY)
			if do {
				result += digitX * digitY
			}
		}
		if lexer.current() == "d" {
			if lexer.doExpression() {
				fmt.Println("do")
				do = true
			}
			if lexer.dontExpression() {
				fmt.Println("dont")
				do = false
			}
		}
	}

	assert.Equal(t, result, 48)
}

func TestDay3_2(t *testing.T) {

	str, _ := loadInputFromFile("3")

	lexer := Lexer{input: str, pos: -1}
	result := 0

	do := true

	for lexer.next() != "EOF" {
		if digitX, digitY, ok := lexer.mulExpression(); ok {

			if do {
				fmt.Println(digitX, digitY)
				result += digitX * digitY
			}
		}
		// if lexer.current() == "d" {
		if lexer.doExpression() {
			fmt.Println("do")
			do = true
		}
		if lexer.dontExpression() {
			fmt.Println("dont")
			do = false
		}
		// }
	}

	// larger 74236024
	fmt.Println(result)
}

func (l *Lexer) doExpression() bool {
	if l.peek(4) != "do()" {
		return false
	}
	l.pos += 3
	return true
}

func (l *Lexer) current() string {
	return string(l.input[l.pos])
}

func (l *Lexer) dontExpression() bool {
	if l.peek(7) != "don't()" {
		return false
	}
	l.pos += 6
	return true
}

func (l *Lexer) peek(n int) string {
	if l.pos+n >= len(l.input) {
		return ""
	}
	str := l.input[l.pos : l.pos+n]
	return str
}

func TestDay3_1(t *testing.T) {

	// str := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	// str := `mul(11,8)mul(8,-5)`

	str, _ := loadInputFromFile("3")

	lexer := Lexer{input: str, pos: -1}
	result := 0

	for lexer.next() != "EOF" {
		if digitX, digitY, ok := lexer.mulExpression(); ok {
			fmt.Println(digitX, digitY)
			result += digitX * digitY
		}
	}
	fmt.Println(result)
	// assert.Equal(t, result, 161)
}

type Lexer struct {
	input string
	pos   int
}

func (l *Lexer) next() string {
	l.pos++
	if l.pos >= len(l.input) {
		return "EOF"
	}
	str := string(l.input[l.pos])
	// fmt.Println(str)
	return str
}

func (l *Lexer) mulExpression() (int, int, bool) {
	if l.input[l.pos] != 'm' {
		return 0, 0, false
	}

	if l.next() != "u" {
		return 0, 0, false
	}

	if l.next() != "l" {
		return 0, 0, false
	}

	if l.next() != "(" {
		return 0, 0, false
	}

	if !isDigit(l.next()) {
		return 0, 0, false
	}
	x := l.readDigit(3)
	if l.next() != "," {
		return 0, 0, false
	}

	if !isDigit(l.next()) {
		return 0, 0, false
	}
	y := l.readDigit(3)
	if l.next() != ")" {
		return 0, 0, false
	}

	return x, y, true
}

func isDigit(l string) bool {
	return l >= "0" && l <= "9"
}

func (l *Lexer) readDigit(max int) int {
	start := l.pos

	for isDigit(l.next()) {
	}

	if l.pos-start > max || l.pos-start == 0 {
		return -1
	}

	d, err := strconv.Atoi(l.input[start:l.pos]) // convert string to int
	if err != nil {
		return -1
	}
	l.pos--

	return d
}
