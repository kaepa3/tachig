package tachig

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func (g *Game) DrawModeStart(screen *ebiten.Image) error {
	text.Draw(screen, "Start", arcadeFont, 300, 20, color.White)
	return nil
}

func (g *Game) UpdateModeStart() error {
	if isKeyJustPressed(ebiten.KeySpace) {
		g.mode = ModePlay
	}
	return nil
}
