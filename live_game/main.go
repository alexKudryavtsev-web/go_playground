package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width     = 80
	height    = 15
	iteration = 100
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)

	for i := range u {
		u[i] = make([]bool, width)
	}

	return u
}

func (p Universe) Set(x, y int, b bool) {
	p[y][x] = b
}

func (p Universe) Alive(x, y int) bool {
	x = (x + width) % width
	y = (y + height) % height

	return p[y][x]
}

func (p Universe) String() string {
	buffer := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var b byte = ' '

			if p[y][x] {
				b = '*'
			}

			buffer = append(buffer, b)
		}
		buffer = append(buffer, '\n')
	}

	return string(buffer)
}

func (u Universe) Neighbors(x, y int) int {
	n := 0
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			if !(v == 0 && h == 0) && u.Alive(x+h, y+v) {
				n++
			}
		}
	}
	return n

}

func (u Universe) Next(x, y int) bool {
	n := u.Neighbors(x, y)
	return n == 3 || n == 2 && u.Alive(x, y)
}

func Step(a, b Universe) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b.Set(x, y, a.Next(x, y))
		}
	}
}

func (p Universe) Show() {
	fmt.Print("\x0c", p.String())
}

func (p Universe) Seed() {
	for i := 0; i < (width*height)/4; i++ {
		p.Set(rand.Intn(width), rand.Intn(height), true)
	}
}

func main() {
	a, b := NewUniverse(), NewUniverse()

	a.Seed()
	a.Show()

	for i := 0; i < iteration; i++ {
		Step(a, b)

		a.Show()
		time.Sleep(time.Second / 30)

		a, b = b, a
	}
	a.Show()

}
