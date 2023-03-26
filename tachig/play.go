package tachig

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) DrawModePlay(screen *ebiten.Image) error {

	text.Draw(screen, "Play", arcadeFont, 300, 20, color.White)
	return nil
}

func (g *Game) UpdateModePlay() error {
	if isKeyJustPressed(ebiten.KeySpace) {
		g.mode = ModeEnd
	}
	return nil
}
