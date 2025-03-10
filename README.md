#  Алгоритмы для прохождения лабиринта #
## Автор: Трофимов Владимир ##

## Прошу рассмотреть алгоритм класса DijkstraScouting ##
Алгоритм ParallelShortcutScouting сделал ради эксперимента. Его описание и результаты будут в конце README.

---
### Содержание ###
- [Интерфейс](#интерфейс)
- [Алгоритм Дейкстры](#алгоритм-дейкстры)
- [Алгоритм параллельного поиска (абсолютно неэффективный)](#алгоритм-параллельного-поиска-абсолютно-неэффективный)
---

### Интерфейс ###
Для возможности дальнейшего расширения множества алгоритмов поиска самого короткого пути был создан интерфейс ShortcutScouting.
```
type ShortcutScouting interface {
    Find(start, finish Point) (shortcut []Point, dist uint, err error)
}
```
Также был создан тип конструкторов всех структур-наследников интерфейса ShortcutScouting. Этот тип используется, например, в параметрах функции ShortcutScoutingTesting.
```
type CreateShortcutScouting func(labyrinth [][]uint) ShortcutScouting
```

### Алгоритм Дейкстры ###
В классе DijkstraScouting был реализован стандартный алгоритм Дейкстры на очереди с приоритетом. Запускается он функцией Find.

Для алгоритма были написаны обычные и стресс тесты.

В стресс-тестах оценивалось время выполнения алгоритма в зависимости от площади лабиринта. При этом рассматривался наихудший случай: стен нет, старт и финиш располагаются в противоположенных углах лабиринта (слева наверху и справа внизу). Получается, что при увеличении площади лабиринта, расстояние между стартом и финишом увеличивается.
#### Полученный график скорости выполнения алгоритма Дейкстры в зависимости от площади лабиринта: ####
![1](https://github.com/Vladimir220/labyrinth/blob/main/tests/Benchmark_DijkstraScouting.png)

### Алгоритм параллельного поиска (абсолютно неэффективный) ###
В классе ParallelShortcutScouting был реализован алгоритм параллельного поиска. Запускается он также функцией Find.

Суть алгоритма:
- В точке старта создаётся горутина-разведчик.
- Разведчик идёт по одному из 4 направлений.
- В остальные 3 направления создаются и отправляются ещё по горутине-разведчику и т.д.
- У каждого разведчика есть срез с историей своего маршрута, карта прошедших точек (для проверок доступности следующей точки), счётчик прошедшего расстояния.
- Перед перемещением в следующую точку проверяется:
    1) Не выходит ли следующий шаг за границы лабиринта?
    2) Не попадает ли следующий шаг на стену?
    3) Не будет ли следующий шаг граничить с траекторией текущего пути? (это будет означать, что точно есть более эффективный маршрут, проходящий через эту и граничащую точки)
    4) Не попадает ли следующий шаг на траекторию текущего пути?
- Разведчики, дошедшие до финиша, обновляют общую переменную, хранящую наикратчайшее расстояние от старта до финиша, если их счётчики показывают меньшее расстояние.
- Разведчики, что дошли до финиша, ждут завершения "гонки".
- Разведчики, которые не дошли до финиша, если их счётчик превышает опубликованное наикратчайшее расстояние, сдаются и самоуничтожаются.
- "Финалисты" сравнивают значения счётчиков, "победитель" отдаёт срез со своим маршрутом, а все остальные самоуничтожаются.

**Ожидание:** 

один из множества разведчиков пойдет идеальной самой короткой дорогой от старта до финиша, максимальная эффективность!

**Реальность:** 
- Алгоритм Дейкстры всё равно эффективнее: Log(S) против Log(S^2), судя по стресс-тестированиям.
- Количество одновременно работающих горутин-разведчиков сильно возрастает при увеличении площади лабиринта.
- Ужасная эффективность по памяти.
- Тратится много времени на выделение и очистку памяти.

**Предположение:** 

алгоритм может быть очень эффективным по времени, если придумать эффективное решение по хранению пути и карты для каждого разведчика, а также придумать, как сократить выделения и очистки памяти. Тогда может мы сможет получить временную эффективность Log($\sqrt{a^2+b^2}$) и при этом приемлемую эффективность по памяти.

Для алгоритма были написаны обычные и стресс тесты.

В стресс-тестах оценивалось время выполнения алгоритма в зависимости от площади лабиринта. При этом рассматривался наихудший случай: стен нет, старт и финиш располагаются в противоположенных углах лабиринта (слева наверху и справа внизу). Получается, что при увеличении площади лабиринта, расстояние между стартом и финишем увеличивается.
#### Полученный график скорости выполнения алгоритма параллельного поиска в зависимости от площади лабиринта: ####
![2](https://github.com/Vladimir220/labyrinth/blob/main/tests/Benchmark_ParallelShortcutScouting.png)

