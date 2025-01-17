package main

import (
	"fmt"
	lab "main/labyrinth"
	"os"
)

func main() {
	labyrinth, start, finish := lab.Input(os.Stdin)
	scouting := lab.CreateShortcutScouting(labyrinth)

	shortcut, dist, err := scouting.Find(start, finish)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	for _, s := range shortcut {
		fmt.Printf("%d %d\n", s.Y, s.X)
	}
	fmt.Println(".")
	fmt.Printf("Длинна самого короткого маршрута: %d\n", dist)
	fmt.Println("(размер клеток старт и финиш не учитываются, как если бы мы сначала стояли на самой границе старта, а затем пришли к самой границе финиша)")
}
