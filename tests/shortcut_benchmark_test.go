package tests

import (
	"fmt"
	"main/labyrinth"
	"math/rand"
	"testing"
	"time"
)

func ShortcutScoutingBenchmark(scoutConstructor labyrinth.CreateShortcutScouting, b *testing.B, name string) {
	testCases := []struct {
		rows, columns int
	}{
		{rows: 6, columns: 2},
		{rows: 6, columns: 3},
		{rows: 6, columns: 4},
		{rows: 6, columns: 5},
		{rows: 6, columns: 6},
		{rows: 6, columns: 7},
		{rows: 6, columns: 8},
		{rows: 6, columns: 9},
		{rows: 6, columns: 10},
	}
	res := map[int]float64{}

	for _, tc := range testCases {
		start := labyrinth.Point{X: 0, Y: 0}
		finish := labyrinth.Point{X: tc.columns - 1, Y: tc.rows - 1}

		labyrinth := make([][]uint, tc.rows)
		for i := range labyrinth {
			labyrinth[i] = make([]uint, tc.columns)
		}

		b.Run(fmt.Sprintf("площадь между стартом и финишом:%d", tc.columns*tc.rows), func(b *testing.B) {
			//runtime.GC()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
				for i := range labyrinth {
					for j := 0; j < tc.columns; j++ {
						labyrinth[i][j] = uint(r.Intn(9)) + 1
					}
				}
				scouting := scoutConstructor(labyrinth)
				b.StartTimer()
				scouting.Find(start, finish)
			}
			res[tc.columns*tc.rows] = float64(b.Elapsed().Nanoseconds()) / float64(b.N)
		})
	}
	drawGraph(res, name)
}

func BenchmarkParallelShortcutScouting(b *testing.B) {
	ShortcutScoutingBenchmark(labyrinth.CreateParallelShortcutScouting, b, "Benchmark_ParallelShortcutScouting")
}

func BenchmarkDijkstraScouting(b *testing.B) {
	ShortcutScoutingBenchmark(labyrinth.CreateDijkstraScouting, b, "Benchmark_DijkstraScouting")
}
