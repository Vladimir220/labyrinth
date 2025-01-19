package tests

import (
	"fmt"
	"main/labyrinth"
)

type TestCase struct {
	Labyrinth     [][]uint
	Start, Finish labyrinth.Point
	shortcut      []labyrinth.Point
	dist          uint
	err           error
}

func GetTestCases() []TestCase {
	return []TestCase{
		{ // Тест 1
			Labyrinth: [][]uint{
				{1, 8, 1, 1, 1, 1},
				{1, 0, 2, 1, 1, 1},
				{8, 0, 1, 2, 1, 2},
				{1, 2, 0, 4, 2, 3},
				{0, 2, 1, 5, 1, 2},
				{1, 1, 1, 6, 7, 1},
				{2, 2, 2, 2, 2, 1},
			},
			Start:  labyrinth.Point{X: 0, Y: 0},
			Finish: labyrinth.Point{X: 4, Y: 5},
			dist:   16,
			shortcut: []labyrinth.Point{
				{X: 0, Y: 0},
				{X: 1, Y: 0},
				{X: 2, Y: 0},
				{X: 3, Y: 0},
				{X: 4, Y: 0},
				{X: 4, Y: 1},
				{X: 4, Y: 2},
				{X: 4, Y: 3},
				{X: 4, Y: 4},
				{X: 4, Y: 5},
			},
			err: nil,
		},
		{ // Тест 2
			Labyrinth: [][]uint{
				{1, 8, 0, 1, 1, 1},
				{1, 0, 2, 1, 1, 1},
				{8, 0, 1, 2, 1, 2},
				{1, 2, 0, 4, 2, 3},
				{0, 2, 1, 5, 1, 2},
				{1, 1, 1, 6, 7, 1},
				{2, 2, 2, 2, 2, 1},
			},
			Start:  labyrinth.Point{X: 0, Y: 0},
			Finish: labyrinth.Point{X: 4, Y: 5},
			dist:   21,
			shortcut: []labyrinth.Point{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 0, Y: 3},
				{X: 1, Y: 3},
				{X: 1, Y: 4},
				{X: 2, Y: 4},
				{X: 3, Y: 4},
				{X: 4, Y: 4},
				{X: 4, Y: 5},
			},
			err: nil,
		},
		{ // Тест 3
			Labyrinth: [][]uint{
				{1, 8, 0, 1, 1, 1},
				{1, 0, 2, 1, 1, 1},
				{8, 0, 1, 2, 1, 2},
				{1, 2, 0, 4, 2, 3},
				{0, 2, 0, 5, 1, 2},
				{1, 0, 0, 6, 7, 1},
				{2, 2, 2, 2, 2, 1},
			},
			Start:    labyrinth.Point{X: 0, Y: 0},
			Finish:   labyrinth.Point{X: 4, Y: 5},
			dist:     0,
			shortcut: nil,
			err:      fmt.Errorf("маршрутов от старта до финиша не существует"),
		},
		{ // Тест 4
			Labyrinth: [][]uint{
				{1, 2, 0},
				{2, 0, 1},
				{9, 1, 0},
			},
			Start:  labyrinth.Point{X: 0, Y: 0},
			Finish: labyrinth.Point{X: 1, Y: 2},
			dist:   11,
			shortcut: []labyrinth.Point{
				{X: 0, Y: 0},
				{X: 0, Y: 1},
				{X: 0, Y: 2},
				{X: 1, Y: 2},
			},
			err: nil,
		},
	}
}
