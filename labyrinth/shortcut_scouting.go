package labyrinth

import (
	"fmt"
	"maps"
	"slices"
	"sync"
	"time"
)

type ShortcutScouting struct {
	labyrinth    [][]uint
	maxPoint     Point
	start        Point
	finish       Point
	shortestDist uint
	sdMux        *sync.RWMutex
	wg           *sync.WaitGroup
}

func (s *ShortcutScouting) checkNextPoint(p Point, prevPoints map[Point]bool) bool {

	checkForNeighborhood :=
		func(p Point, points map[Point]bool) bool {
			if points[p.GoUp(1)] || points[p.GoDown(1)] || points[p.GoLeft(1)] || points[p.GoRight(1)] {
				return true
			}
			return false
		}

	// Проверки на:
	// 1) Не выходит ли следующий шаг за границы лабиринта?
	// 2) Не попадает ли следующий шаг на стену?
	// 3) Не будет ли следующий шаг граничить с траекторией текущего пути?
	// (это будет означать, что точно есть более эффективный маршрут, проходящий через эту и граничущую точки)
	// 4) Не попадает ли следуюший шаг на траекторию текущего пути?
	return p.InAria(Point{0, 0}, s.maxPoint) &&
		(s.labyrinth[p.Y][p.X] != 0) &&
		!checkForNeighborhood(p, prevPoints) &&
		!prevPoints[p]

}

func (s *ShortcutScouting) sendScout(route []Point, prevPoints map[Point]bool, loc Point, distCounter uint, res chan<- []Point, stopSignal chan interface{}) {
	for loc != s.finish {
		individualGoalSelected := false
		upPoint := loc.GoUp(1)
		isTopAvailable := s.checkNextPoint(upPoint, prevPoints)

		downPoint := loc.GoDown(1)
		isBottomAvailable := s.checkNextPoint(downPoint, prevPoints)

		leftPoint := loc.GoLeft(1)
		isLeftAvailable := s.checkNextPoint(leftPoint, prevPoints)

		rightPoint := loc.GoRight(1)
		isRightAvailable := s.checkNextPoint(rightPoint, prevPoints)

		// Проверяем, можем ли мы хоть куда-нибудь двигаться
		if !(isTopAvailable || isBottomAvailable || isLeftAvailable || isRightAvailable) {
			s.wg.Done()
			return
		}

		// Добавляем текущую точку в маршрут, поскольку мы её точно преодолеем
		route = append(route, loc)
		prevPoints[loc] = true

		// Обновляем счётчик расстояния (поскольку мы заканчиваем прохождение предыдущей клетки), если мы были не в точке старт
		if loc != s.start {
			distCounter += s.labyrinth[loc.Y][loc.X]
		}

		// Нет смысла идти дальше, если путь уже точно не эффективный
		s.sdMux.RLock()
		if distCounter > s.shortestDist {
			s.sdMux.RUnlock()
			s.wg.Done()
			return
		}
		s.sdMux.RUnlock()

		// Двигаемся / мобилизуем больше разведчиков на новые маршруты
		if isTopAvailable {
			loc = upPoint
			individualGoalSelected = true
		}
		if isBottomAvailable {
			if individualGoalSelected {
				s.wg.Add(1)
				buf1 := maps.Clone(prevPoints)
				buf2 := slices.Clone(route)
				go s.sendScout(buf2, buf1, downPoint, distCounter, res, stopSignal)
			} else {
				loc = downPoint
				individualGoalSelected = true
			}
		}
		if isLeftAvailable {
			if individualGoalSelected {
				s.wg.Add(1)
				buf1 := maps.Clone(prevPoints)
				buf2 := slices.Clone(route)
				go s.sendScout(buf2, buf1, leftPoint, distCounter, res, stopSignal)
			} else {
				loc = leftPoint
				individualGoalSelected = true
			}
		}
		if isRightAvailable {
			if individualGoalSelected {
				s.wg.Add(1)
				buf1 := maps.Clone(prevPoints)
				buf2 := slices.Clone(route)
				go s.sendScout(buf2, buf1, rightPoint, distCounter, res, stopSignal)
			} else {
				loc = rightPoint
				individualGoalSelected = true
			}
		}
	}

	// Разведчик дошел до финиша!
	route = append(route, s.finish)

	// Проверяем, лучший ли у него всё ещё результат
	s.sdMux.Lock()
	if distCounter < s.shortestDist {
		s.shortestDist = distCounter
	} else {
		s.sdMux.Unlock()
		s.wg.Done()
		return
	}
	s.sdMux.Unlock()

	// Теперь ждём остальных разведчиков, чтобы определить, у кого наилучший маршрут
	s.wg.Done()
	s.wg.Wait()

	// Сдаёмся, если после ожидания кто-то из финалистов обогнал нас
	s.sdMux.RLock()
	if distCounter > s.shortestDist {
		s.sdMux.RUnlock()
		return
	}
	s.sdMux.RUnlock()

	// Победитель может быть только один...
	s.sdMux.Lock()
	select {
	case <-stopSignal:
		return
	default:
	}
	close(stopSignal)
	s.sdMux.Unlock()

	// Победитель передаёт наилучший маршрут
	res <- route
}

// Этот метод публикуем
func (s *ShortcutScouting) Find(start, finish Point) (shortcut []Point, dist uint, err error) {
	s.start = start
	s.finish = finish
	res := make(chan []Point)
	stopSignal := make(chan interface{})

	s.wg.Add(1)
	go s.sendScout([]Point{}, map[Point]bool{}, start, 0, res, stopSignal)

	s.wg.Wait()

	s.sdMux.RLock()
	if s.shortestDist == ^uint(0) {
		s.sdMux.RUnlock()
		err = fmt.Errorf("маршрутов от старта до финиша не существует")
		return
	}
	s.sdMux.RUnlock()

	select {
	case shortcut = <-res:
		dist = s.shortestDist
	case <-time.After(5 * time.Second):
		err = fmt.Errorf("timeout: что-то пошло не так")
	}

	return
}

// Эту функцию публикуем
func CreateShortcutScouting(labyrinth [][]uint) *ShortcutScouting {
	return &ShortcutScouting{labyrinth: labyrinth, maxPoint: Point{Y: len(labyrinth) - 1, X: len(labyrinth[0]) - 1}, sdMux: &sync.RWMutex{}, wg: &sync.WaitGroup{}, shortestDist: ^uint(0)}
}
