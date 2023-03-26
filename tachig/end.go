package tachig

import (
	_ "embed"
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

//go:embed staff.txt
var staffText string
var (
	lines  []string
	posX   = 0
	posY   = 0
	width  = 0
	height = 0
)

const (
	step = 20
)

func init() {
	lines = strings.Split(staffText, "\n")
	for i := 0; i < 10; i++ {
		fmt.Println(lines[i])
	}
	posY = 0

}

func (g *Game) DrawModeEnd(screen *ebiten.Image) error {
	width, height = screen.Size()
	posX = width / 2

	for i, v := range lines {
		current := posY - (i * step)
		if current < height && 0 < current {
			text.Draw(screen, v, arcadeFont, posX, current, color.White)
		}
	}
	return nil
}

func (g *Game) UpdateModeEnd() error {
	if isKeyJustPressed(ebiten.KeySpace) {
		g.mode = ModeStart
	}
	posY += 1

	return nil
}
