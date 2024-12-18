package main

import (
	"errors"
	"fmt"
	"strings"
)

// Matrix is a generic type for a 2D matrix with "filled" and "empty" states.

type data[T any] [][]*T

type Matrix[T any] struct {
	data       data[T]
	rows, cols int
}

// NewMatrix creates a new matrix with the specified number of rows and columns.
func NewMatrix[T any](rows, cols int, defaultValue T) *Matrix[T] {
	data := make([][]*T, rows)
	for i := range data {
		data[i] = make([]*T, cols)
	}
	return &Matrix[T]{data: data, rows: rows, cols: cols}
}

// Rows returns the number of rows in the matrix.
func (m *Matrix[T]) Rows() int {
	return m.rows
}

// Cols returns the number of columns in the matrix.
func (m *Matrix[T]) Cols() int {
	return m.cols
}

// Set sets the value at a specific row and column in the matrix.
func (m *Matrix[T]) Set(row, col int, value T) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return errors.New("index out of bounds")
	}
	m.data[row][col] = &value
	return nil
}

// Get retrieves the value at a specific row and column in the matrix.
func (m *Matrix[T]) Get(row, col int) (*T, error) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		var zero T
		return &zero, errors.New("index out of bounds")
	}
	v := m.data[row][col]

	return v, nil
}

// GetLeft retrieves the value to the left of the given position.
func (m *Matrix[T]) GetLeft(row, col int) (*T, error) {
	// fmt.Println("GetLeft: ", row, col-1)
	return m.Get(row, col-1)
}

// GetRight retrieves the value to the right of the given position.
func (m *Matrix[T]) GetRight(row, col int) (*T, error) {
	// fmt.Println("GetRight: ", row, col+1)
	return m.Get(row, col+1)
}

// GetUp retrieves the value above the given position.
func (m *Matrix[T]) GetUp(row, col int) (*T, error) {
	// fmt.Println("GetUp: ", row-1, col)
	return m.Get(row-1, col)
}

// GetDown retrieves the value below the given position.
func (m *Matrix[T]) GetDown(row, col int) (*T, error) {
	// fmt.Println("GetDown: ", row+1, col)
	return m.Get(row+1, col)
}

// String provides a string representation of the matrix (for debugging).
func (m *Matrix[T]) String() string {

	str := ""
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			str += fmt.Sprint(*m.data[i][j])
		}
		str += "\n"
	}
	// remove trailing newline
	str = str[:len(str)-1]
	return str
}

func (m *Matrix[T]) Clear(row, col int) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return errors.New("index out of bounds")
	}
	var zero T
	m.data[row][col] = &zero
	return nil
}

func NewMatrixFromString(str string) (matrix Matrix[string]) {

	rows := strings.Split(str, "\n")
	matrix = *NewMatrix(len(rows), len(rows[0]), "")

	for i := 0; i < len(rows); i++ {
		col := strings.Split(rows[i], "")
		for j := 0; j < len(col); j++ {
			col[j] = strings.TrimSpace(col[j])
			if col[j] != "" {
				matrix.Set(i, j, col[j])
			}
		}
	}
	return
}
