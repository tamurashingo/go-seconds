package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tamurashingo/go-seconds/game"
)

func main() {
	fmt.Println("go-seconds")
	ebiten.SetWindowSize(480, 640)
	ebiten.SetWindowTitle("go-seconds")
	game := &game.Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
