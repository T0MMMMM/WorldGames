package gameengine

import (
	"fmt"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

type EngineStruct struct {
	run bool
	bgColor rl.Color

	worldToFind []rune
	worldFind []rune
	listeWorldFind [][]string
	lenght int
	tryNum int32
	win bool
	letterValid []rune
	posLetterValid []int
	listWorlds []string
	listWorldsEnter []string

	font rl.Font

	color map[string]rl.Color

	texture TextureStruct
	letter []bool
	inputLetter []int32
}

func (g *EngineStruct) Play() {

	g.init()

	for g.run {
		g.update()
		g.input()
		g.render()
	}
	time.Sleep(3 * time.Second)

	for i := 0; i < g.lenght; i ++ {
		fmt.Printf(string(g.worldToFind[i]))
	}
	fmt.Printf("\n")
}
