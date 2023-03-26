package tachig

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Mode int

const (
	ModeStart = iota
	ModePlay
	ModeEnd
)

const (
	fontSize = 10
)

var (
	arcadeFont font.Face
)

type Game struct {
	mode Mode
}

func NewGame() *Game {
	g := Game{}
	g.init()
	return &g
}

//go:embed fonts
var fontsDir embed.FS

func (g *Game) init() {
	g.mode = ModeStart
	bytes, err := fontsDir.ReadFile("fonts/Cica-Regular.ttf")
	tt, err := opentype.Parse(bytes)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 96
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func (g *Game) Draw(screen *ebiten.Image) {
	var err error
	switch g.mode {
	case ModeStart:
		err = g.DrawModeStart(screen)
	case ModePlay:
		err = g.DrawModePlay(screen)
	case ModeEnd:
		err = g.DrawModeEnd(screen)
	}
	if err != nil {
		log.Println(err)
	}
}

func (g *Game) Update() error {
	var err error
	switch g.mode {
	case ModeStart:
		err = g.UpdateModeStart()
	case ModePlay:
		err = g.UpdateModePlay()
	case ModeEnd:
		err = g.UpdateModeEnd()
	}
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func isKeyJustPressed(key ebiten.Key) bool {
	if inpututil.IsKeyJustPressed(key) {
		return true
	}
	return false
}
