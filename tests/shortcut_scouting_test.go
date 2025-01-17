package main

import (
	"fmt"
	"image/color"
	lab "main/labyrinth"
	"math/rand"
	"sort"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func BenchmarkFindWays(b *testing.B) {
	testCases := []struct {
		rows, columns int
	}{
		{rows: 7, columns: 3},
		{rows: 7, columns: 7},
		{rows: 7, columns: 9},
		{rows: 7, columns: 10},
		{rows: 7, columns: 11},
		{rows: 7, columns: 12},
		{rows: 7, columns: 13},
	}
	res := map[int]float64{}

	for _, tc := range testCases {
		start := lab.Point{X: 0, Y: 0}
		finish := lab.Point{X: tc.columns - 1, Y: 0}

		labyrinth := make([][]uint, tc.rows)
		for i := range labyrinth {
			labyrinth[i] = make([]uint, tc.columns)
		}

		b.Run(fmt.Sprintf("площадь между стартом и финишом:%d", tc.columns*tc.rows), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
				for i := range labyrinth {
					for j := 0; j < tc.columns; j++ {
						labyrinth[i][j] = uint(r.Intn(9)) + 1
					}
				}
				scouting := lab.CreateShortcutScouting(labyrinth)
				b.StartTimer()
				scouting.Find(start, finish)
			}
			res[tc.columns*tc.rows] = b.Elapsed().Seconds()
		})
	}

	// Создаем новый график
	p := plot.New()

	// Устанавливаем заголовки
	p.Title.Text = "График времени исполнения функции"
	p.X.Label.Text = "Площадь"
	p.Y.Label.Text = "Время (сек)"

	// Подготовка данных для графика
	type kv struct {
		Key   int
		Value float64
	}

	var sortedData []kv
	for k, v := range res {
		sortedData = append(sortedData, kv{k, v})
	}

	// Сортируем по ключу (площадь)
	sort.Slice(sortedData, func(i, j int) bool {
		return sortedData[i].Key < sortedData[j].Key
	})

	// Создаем точки для графика
	points := make(plotter.XYs, len(sortedData))
	for i, kv := range sortedData {
		points[i].X = float64(kv.Key)
		points[i].Y = kv.Value
	}

	// Создаем линию графика
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Красный цвет линии

	// Добавляем линию на график
	p.Add(line)

	// Сохраняем график в файл
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "graphh.png"); err != nil {
		panic(err)
	}
}
