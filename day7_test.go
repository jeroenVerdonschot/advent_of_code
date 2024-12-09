package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay7_1(t *testing.T) {

	str := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	str, _ = loadInputFromFile("7")

	rows := strings.Split(str, "\n")

	sum := 0

	for i, row := range rows {
		answerStr := strings.Split(row, ":")
		answer, _ := strconv.Atoi(answerStr[0])
		// answers = append(answers, int64(answer))

		inputStr := strings.Split(strings.TrimSpace(answerStr[1]), " ")
		part := []int{}
		for _, i := range inputStr {
			in, _ := strconv.Atoi(i)
			part = append(part, in)
		}

		v := handle(answer, part)

		fmt.Println(i)

		sum += v
	}
	fmt.Println(sum)
}

func TestDay7_2(t *testing.T) {

	str := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

	// str, _ = loadInputFromFile("7")

	rows := strings.Split(str, "\n")

	sum := 0

	for i, row := range rows {
		answerStr := strings.Split(row, ":")
		answer, _ := strconv.Atoi(answerStr[0])

		inputStr := strings.Split(strings.TrimSpace(answerStr[1]), " ")
		part := []int{}
		for _, i := range inputStr {
			in, _ := strconv.Atoi(i)
			part = append(part, in)
		}

		v := handle(answer, part)

		fmt.Println(i)

		sum += v
	}
	fmt.Println(sum)
}

func TestGetOperands(t *testing.T) {

	test := []int{10, 19}
	value := 190

	v := handle(value, test)
	_ = v
}

func handle3(value int, test []int) int {

	arrLen := len(test) - 1

	l := 1 << arrLen
	format := fmt.Sprintf("%%0%db", arrLen) // Create the format string once

	for i := 0; i < l; i++ {

		code := fmt.Sprintf(format, i)

		for j := 0; j < len(test)-1; j++ {

			code2 := fmt.Sprintf(format, j)
			ccArr := concat(test, code2)

			fmt.Println(ccArr, j)
			v := 0

			if len(ccArr) > 1 {
				v = calculator(ccArr, code)
			}

			fmt.Println(code, v)
			if value == v {
				return v
			}
		}
	}
	return 0
}

func handle(value int, test []int) int {

	arrLen := len(test) - 1

	l := 1 << arrLen
	format := fmt.Sprintf("%%0%db", arrLen) // Create the format string once

	for i := 0; i < l; i++ {

		code := fmt.Sprintf(format, i)

		v := calculator(test, code)

		if value == v {
			return v
		}
	}
	return 0
}

func concat(arr []int, ops string) []int {

	var values []int

	str := fmt.Sprint(arr[0])

	for i, o := range ops {

		if o == '0' {
			str += ","
			str += fmt.Sprint(arr[i+1])
		}
		if o == '1' { // CONCAT
			str += fmt.Sprint(arr[i+1])
		}
	}

	// values = strings.split(str, ",")
	a := strings.Split(str, ",")

	values = arrStrToArrInt(a)

	return values
}

func arrStrToArrInt(arr []string) (values []int) {
	for _, s := range arr {
		v, _ := strconv.Atoi(s)
		values = append(values, v)
	}
	return
}

func TestConcat(t *testing.T) {

	arr := []int{1, 2, 3, 4}
	ops := "000"

	res := concat(arr, ops)
	fmt.Println(res)
	//result 1,2,3,4

	arr = []int{1, 2, 3, 4}
	ops = "010"

	res = concat(arr, ops)
	fmt.Println(res)
	// 1,23,4

	arr = []int{1, 2, 3, 4}
	ops = "111"

	res = concat(arr, ops)
	fmt.Println(res)
	// 1234
}

func calculator(arr []int, ops string) int {

	value := arr[0]

	for i, o := range ops {

		if o == '0' {
			value += arr[i+1]
		} else {
			value *= arr[i+1]
		}

	}
	return value
}

func TestCalculator(t *testing.T) {

	arr := []int{1, 2, 3, 4}
	ops := "000"

	res := calculator(arr, ops)

	assert.Equal(t, res, 10)

}

func TestCalculator2(t *testing.T) {

	arr := []int{1, 2, 3, 4}
	ops := "111"

	res := calculator(arr, ops)

	assert.Equal(t, res, 24)

}
