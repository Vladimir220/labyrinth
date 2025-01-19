package tests

import (
	"fmt"
	"main/labyrinth"
	"slices"
	"testing"
)

func ShortcutScoutingTesting(scoutConstructor labyrinth.CreateShortcutScouting, t *testing.T) {
	cases := GetTestCases()

	for id, tc := range cases {
		t.Run(fmt.Sprintf("Тест %d", id), func(t *testing.T) {
			lab := tc.Labyrinth
			start := tc.Start
			finish := tc.Finish

			scouting := scoutConstructor(lab)
			shortcut, dist, err := scouting.Find(start, finish)

			if err == nil {
				if slices.Equal(shortcut, tc.shortcut) && (dist == tc.dist) {
					t.Logf("Успех!")
				}
			} else if err.Error() == tc.err.Error() {
				t.Logf("Успех!")
			} else {
				t.Errorf("Ошибка: %v", err)
			}
		})
	}
}

func TestParallelShortcutScouting(t *testing.T) {
	ShortcutScoutingTesting(labyrinth.CreateParallelShortcutScouting, t)
}

func TestDijkstraScouting(t *testing.T) {
	ShortcutScoutingTesting(labyrinth.CreateDijkstraScouting, t)
}
