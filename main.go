package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaepa3/tachinomi/tachig"
)

const (
	screenX = 640
	screenY = 360
)

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle("Tachinomi")
	if err := ebiten.RunGame(tachig.NewGame()); err != nil {
		log.Fatal(err)
	}

}
