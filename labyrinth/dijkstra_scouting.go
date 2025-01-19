package labyrinth

import (
	"container/heap"
	"fmt"
	"slices"
)

type DijkstraScouting struct {
	labyrinth [][]uint
	maxPoint  Point
}

func (ds DijkstraScouting) checkNextPoint(p Point, visited map[Point]bool) bool {
	// Проверки на:
	// 1) Не выходит ли следующий шаг за границы лабиринта?
	// 2) Не попадает ли следующий шаг на стену?
	// 3) Не был ли уже посещён следующий шаг
	return p.InAria(Point{0, 0}, ds.maxPoint) && (ds.labyrinth[p.Y][p.X] != 0) && !visited[p]

}

// Запускает алгоритм Дейкстры поиска самого короткого пути от старта до финиша
func (ds DijkstraScouting) Find(start, finish Point) (shortcut []Point, dist uint, err error) {
	distMap := map[Point]uint{}
	prev := map[Point]Point{}
	visited := map[Point]bool{}
	queue := &PriorityQueue{}
	heap.Init(queue)

	for rowId, row := range ds.labyrinth {
		for columnId := range row {
			p := Point{X: columnId, Y: rowId}
			distMap[p] = ^uint(0)
		}
	}
	distMap[start] = 0
	heap.Push(queue, &Node{P: start, Dist: 0})

	for queue.Len() > 0 {
		node := heap.Pop(queue).(*Node)
		currPoint := node.P
		visited[currPoint] = true

		neighbours := [4]Point{currPoint.GoUp(1), currPoint.GoDown(1), currPoint.GoLeft(1), currPoint.GoRight(1)}

		var distToNeighbors uint = 0
		if currPoint != start {
			distToNeighbors = distMap[currPoint] + ds.labyrinth[currPoint.Y][currPoint.X]
		}

		for _, neigh := range neighbours {
			if ds.checkNextPoint(neigh, visited) && (distToNeighbors < distMap[neigh]) {
				distMap[neigh] = distToNeighbors
				prev[neigh] = currPoint
				heap.Push(queue, &Node{P: neigh, Dist: distToNeighbors})
			}
		}
	}

	if !visited[finish] {
		err = fmt.Errorf("маршрутов от старта до финиша не существует")
		return
	}

	dist = distMap[finish]
	pointBuf := finish
	for pointBuf != start {
		shortcut = append(shortcut, pointBuf)
		pointBuf = prev[pointBuf]
	}
	shortcut = append(shortcut, start)
	slices.Reverse(shortcut)

	return
}

// Создаёт объект типа DijkstraScouting и возвращает его под типом интерфейса ShortcutScouting
func CreateDijkstraScouting(labyrinth [][]uint) ShortcutScouting {
	return DijkstraScouting{labyrinth: labyrinth, maxPoint: Point{X: len(labyrinth[0]) - 1, Y: len(labyrinth) - 1}}
}
