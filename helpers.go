package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const SESSIONCOOKIE = "session=53616c7465645f5f567f717a5e9ed6261a183f8b086ff4e3b0ccca1f1cb33e858bf74adc50d9446ccd0d85f888d18709bfb3b80a34822ce2e35e594b18f636c2"

func getUrl(day string) string {
	return "https://adventofcode.com/2024/day/" + day + "/input"
}

func loadInput(day string) (string, error) {
	client := &http.Client{}

	url := getUrl(day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", SESSIONCOOKIE)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func deltas(row Row) (deltas []int) {
	for i := 0; i < len(row)-1; i++ {
		deltas = append(deltas, row[i+1]-row[i])
	}
	return
}

func allPositive(row Row) bool {
	for _, v := range row {
		if v < 0 {
			return false
		}
	}
	return true
}

func allSameSigned(row Row) bool {
	sign := row[0] > 0
	for _, v := range row {
		if (v > 0) != sign {
			return false
		}
	}
	return true
}

func maxInt(row Row) int {
	max := row[0]
	for _, v := range row {
		if v > max {
			max = v
		}
	}
	return max
}

func minInt(row Row) int {
	min := row[0]
	for _, v := range row {
		if v < min {
			min = v
		}
	}
	return min
}

func loadInputFromFile(day string) (string, error) {
	f, err := os.Open("inputs/day" + day + ".input")
	if err != nil {
		return "", err
	}

	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func getIntegers(buffer []byte) (arrOfint []int) {

	var intValue []byte
	var index int

	for index < len(buffer) {

		for isInt(buffer[index]) {
			intValue = append(intValue, buffer[index])
			index++
			if index >= len(buffer) {
				break
			}
		}

		if len(intValue) > 0 {
			num, err := strconv.Atoi(string(intValue))
			if err == nil {
				arrOfint = append(arrOfint, num)
			}
			intValue = []byte{}
		}
		index++
	}
	return
}

func isInt(b byte) bool {
	return string(b) >= "0" && string(b) <= "9"
}

type Row []int

func (r Row) remove(i int) Row {
	newRow := make(Row, len(r)-1)
	copy(newRow, r[:i])
	copy(newRow[i:], r[i+1:])
	return newRow
}

type Maxtrix []Row

func makeMatrixOfInts(str string) Maxtrix {

	rows := strings.Split(str, "\n")

	matrix := make(Maxtrix, len(rows))

	for i := 0; i < len(rows); i++ {

		col := strings.Split(rows[i], " ")
		for j := 0; j < len(col); j++ {
			col[j] = strings.TrimSpace(col[j])
			v, _ := strconv.Atoi(col[j])
			matrix[i] = append(matrix[i], v)
		}
	}
	return matrix
}
