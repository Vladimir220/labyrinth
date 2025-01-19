package tests

import (
	"fmt"
	"image/color"
	"slices"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func drawGraph(res map[int]float64, name string) {
	p := plot.New()

	p.Title.Text = "График времени исполнения функции"
	p.X.Label.Text = "Площадь"
	p.Y.Label.Text = "Время (мсек)"

	type kv struct {
		Key   int
		Value float64
	}
	var sortedData []kv
	for k, v := range res {
		sortedData = append(sortedData, kv{k, v})
	}
	slices.SortFunc(sortedData, func(a, b kv) int {
		return a.Key - b.Key
	})

	points := make(plotter.XYs, len(sortedData))
	for i, kv := range sortedData {
		points[i].X = float64(kv.Key)
		points[i].Y = kv.Value / 1000000
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}

	p.Add(line)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s.png", name)); err != nil {
		panic(err)
	}
}
