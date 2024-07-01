package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows  = 9
	cols  = 9
	empty = 0
)

var (
	ErrBounds     = errors.New("за пределами")
	ErrDigit      = errors.New("неправильная цифра")
	ErrInRow      = errors.New("цифра уже есть в ряду")
	ErrInCol      = errors.New("цифра уже есть в строчке")
	ErrInRegion   = errors.New("в данной части эта цифра уже есть")
	ErrFixedDigit = errors.New("начальные цифры нельзя переписать")
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][cols]Cell

func NewSudoku(defaultValue [cols][rows]int8) *Grid {
	var grid Grid

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			curValue := defaultValue[r][c]

			if curValue != empty {
				grid[r][c] = Cell{
					digit: curValue,
					fixed: true,
				}
			}
		}
	}

	return &grid
}

func (g *Grid) Set(row, col int, digit int8) (err error) {
	switch {
	case !inBounds(row, col):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(col, digit):
		return ErrInCol
	case g.inRegion(row, col, digit):
		return ErrInRegion
	case g.isFixed(row, col):
		return ErrFixedDigit
	}

	g[row][col].digit = digit
	return nil
}

func (g *Grid) Clear(row, col int) (err error) {
	switch {
	case !inBounds(row, col):
		return ErrBounds
	case g.isFixed(row, col):
		return ErrFixedDigit
	}

	g[row][col].digit = empty

	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= cols {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][column].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < cols; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3
	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func main() {
	s := NewSudoku([rows][cols]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(1, 1, 2)

	if err != nil {
		fmt.Println("Ошибка: значение", err)
		os.Exit(1)
	}

	fmt.Println(s)
}
