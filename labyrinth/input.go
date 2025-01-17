package labyrinth

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func checkBufioReadStringErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "не удалось считать данные: %v\n", err)
		os.Exit(1)
	}
}

func checkStrconvAtoiErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка преобразования значения в число: %v\n", err)
		os.Exit(1)
	}
}

func Input(rd io.Reader) (labyrinth [][]uint, start, finish Point) {
	reader := bufio.NewReader(rd)
	var (
		err   error
		input string
	)

	// Читаем размер лабиринта
	input, err = reader.ReadString('\n')
	checkBufioReadStringErr(err)
	labyrinthSizeSlice := strings.Split(strings.TrimSpace(input), " ")
	if len(labyrinthSizeSlice) != 2 {
		fmt.Fprintf(os.Stderr, "ожидалось 2 значения, получено - %d\n", len(labyrinthSizeSlice))
		os.Exit(1)
	}
	rows, err := strconv.Atoi(labyrinthSizeSlice[0])
	checkStrconvAtoiErr(err)
	columns, err := strconv.Atoi(labyrinthSizeSlice[1])
	checkStrconvAtoiErr(err)
	if rows < 0 || columns < 0 {
		fmt.Fprintf(os.Stderr, "получено отрицательное значение размера лабиринта: (%d, %d)\n", rows, columns)
		os.Exit(1)
	}

	// Выделяем память под лабиринт
	labyrinth = make([][]uint, rows)
	for i := 0; i < rows; i++ {
		labyrinth[i] = make([]uint, columns)
	}

	// Заполняем лабиринт
	for i := 0; i < rows; i++ {
		input, err = reader.ReadString('\n')
		checkBufioReadStringErr(err)
		cellSizesSlice := strings.Split(strings.TrimSpace(input), " ")
		if len(cellSizesSlice) != columns {
			fmt.Fprintf(os.Stderr, "ожидалось %d столбцов, получено: %d\n", columns, len(cellSizesSlice))
			os.Exit(1)
		}

		for j := 0; j < columns; j++ {
			cellSize, err := strconv.Atoi(cellSizesSlice[j])
			checkStrconvAtoiErr(err)
			if cellSize < 0 || cellSize > 9 {
				fmt.Fprintf(os.Stderr, "ожидался размер ячейки от 0 до 9, получен: %d\n", cellSize)
				os.Exit(1)
			} else {
				labyrinth[i][j] = uint(cellSize)
			}
		}
	}

	// Читаем координаты начала и конца маршрута
	input, err = reader.ReadString('\n')
	checkBufioReadStringErr(err)
	coordsSlice := strings.Split(strings.TrimSpace(input), " ")
	if len(coordsSlice) != 4 {
		fmt.Fprintf(os.Stderr, "ожидалось 4 значения, получено - %d\n", len(coordsSlice))
		os.Exit(1)
	}

	coords := make([]int, 0, 4)
	for _, c := range coordsSlice {
		ci, err := strconv.Atoi(c)
		checkStrconvAtoiErr(err)
		coords = append(coords, ci)
	}
	start = Point{Y: coords[0], X: coords[1]}
	if !start.InAria(Point{0, 0}, Point{rows - 1, columns - 1}) {
		fmt.Fprintf(os.Stderr, "указанная точка старта (%d, %d), выходит за границы лабиринта (0, 0)-(%d, %d)\n", start.Y, start.X, rows-1, columns-1)
		os.Exit(1)
	}
	finish = Point{Y: coords[2], X: coords[3]}
	if !start.InAria(Point{0, 0}, Point{Y: rows - 1, X: columns - 1}) {
		fmt.Fprintf(os.Stderr, "указанная точка финиша (%d, %d), выходит за границы лабиринта (0, 0)-(%d, %d)\n", finish.Y, finish.X, rows-1, columns-1)
		os.Exit(1)
	}

	return
}
