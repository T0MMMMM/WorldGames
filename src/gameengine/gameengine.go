package gameengine

import (
	"fmt"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

type EngineStruct struct {
	run bool
	bgColor rl.Color

	showMenu bool
	numberGame string

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

	letterTest []int
	letterTestindice []int
	letters []int
	usedLetters []int32
	lettersWithoutDoble []int

	font rl.Font

	color map[string]rl.Color

	texture TextureStruct
	letter []bool
	inputLetter []int32
}

func (g *EngineStruct) Play() {

	g.init()

	for g.run && g.showMenu {
		g.menu()
	}

	for g.run {
		//Si le jeu est "Wordle", on lance ce Wordle, pareil pour Hangman
		if g.numberGame == "Wordle" {
			g.updateWordle()
			g.inputWordle()
			g.renderWordle()
		}
		if g.numberGame == "Hangman" {
			g.updateHangman()
			g.inputHangman()
			g.renderHangman()
		}
	}
	time.Sleep(2 * time.Second)

	for i := 0; i < g.lenght; i ++ {
		fmt.Printf(string(g.worldToFind[i]))
	}
	fmt.Printf("\n")
}
