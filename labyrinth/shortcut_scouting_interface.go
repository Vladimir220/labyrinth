package labyrinth

// Интерфейс для алгоритмов нахождения самых коротких маршрутов
type ShortcutScouting interface {
	// Запускает алгоритм поиска самого короткого пути от старта до финиша
	Find(start, finish Point) (shortcut []Point, dist uint, err error)
}

// Тип всех конструкторов структур-наследников интерфейса ShortcutScouting
type CreateShortcutScouting func(labyrinth [][]uint) ShortcutScouting
