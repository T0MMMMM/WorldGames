package gameengine

import (
	"bufio"
	"math/rand"
    "time"
	"fmt"
	"os"

	"github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth = 1600
	screenHeight = 900
	squareSize = 64
	spaceBetweenSquare = 10
)

type TextureStruct struct {
	keybord rl.Texture2D
	hangman rl.Texture2D
	wordle rl.Texture2D
	fondDegrade rl.Texture2D
	wordgame rl.Texture2D
}

func (g *EngineStruct) testEq(a, b []rune) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func (g *EngineStruct) stringInSlice(a rune, list []rune) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func (g *EngineStruct) count(a rune, list []rune) int {
	counter := 0
	for i := range list {
		if list[i] == a {
			counter++
		}
	}
	return counter
}

func (g *EngineStruct) inListeFind(k int, list [][]string) bool {
	if len(list[len(g.listeWorldFind)-1]) <= 1 {
		return true
	}
	for i := 0; i < len(list[len(g.listeWorldFind)-1])-1; i++ {
		if rune(list[len(g.listeWorldFind)-1][0][i]) == g.worldFind[k] && list[len(g.listeWorldFind)-1][i+1] == "yellow" {
			return false
		}
	}
	return true
}

func (g *EngineStruct) printWord() {
	mot := " "
	for i := 0; i < g.lenght; i ++ {
		mot += string(g.worldToFind[i])
	}
	rl.DrawTextEx(g.font, "LE MOT ETAIT : ", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2), float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)+20), 50, 0, rl.Black)	
	rl.DrawTextEx(g.font, mot, rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+295, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)+20), 50, 0, rl.Black)
}

func (g *EngineStruct) init() {
	rl.InitWindow(screenWidth, screenHeight, "|")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	g.run = true
	g.bgColor = rl.NewColor(21, 21, 21, 255)
	g.win = false

	g.numberGame = ""
	g.showMenu = true

	g.color = make(map[string]rl.Color)

	g.usedLetters = append(g.usedLetters, 0)

	g.color["green"] = rl.NewColor(0, 191, 99, 255)
	g.color["yellow"] = rl.NewColor(255, 222, 89, 255)
	g.color["grey"] = rl.Black
	g.font = rl.LoadFont("../texture/neue-helvetica-75-bold.otf")
	
	g.texture.keybord = rl.LoadTexture("../texture/clavierGris.png")
	g.texture.hangman = rl.LoadTexture("../texture/hangman.png")
	g.texture.wordle = rl.LoadTexture("../texture/wordle.png")
	g.texture.fondDegrade = rl.LoadTexture("../texture/fond.png")
	g.texture.wordgame = rl.LoadTexture("../texture/wordgame.png")
	
	g.inputLetter = []int32{rl.KeyA, rl.KeyB, rl.KeyC, rl.KeyD, rl.KeyE, rl.KeyF, rl.KeyG, rl.KeyH, rl.KeyI, rl.KeyJ, rl.KeyK, rl.KeyL, rl.KeyM, rl.KeyN, rl.KeyO, rl.KeyP, rl.KeyQ, rl.KeyR, rl.KeyS, rl.KeyT, rl.KeyU, rl.KeyV, rl.KeyW, rl.KeyX, rl.KeyY, rl.KeyZ}
	for i := 0; i < 26; i++ {
		g.letter = append(g.letter, false)
	}

	readFile, err := os.Open("../texture/pli07.txt")
	if err != nil {
			fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		g.listWorlds = append(g.listWorlds, fileScanner.Text())
	}
	readFile.Close()

	readFile2, err := os.Open("../texture/pli07.txt")
	if err != nil {
			fmt.Println(err)
	}
	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)
	for fileScanner2.Scan() {
		g.listWorldsEnter = append(g.listWorldsEnter, fileScanner2.Text())
	}
	readFile2.Close()

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(g.listWorlds))

	
	g.worldToFind = []rune(g.listWorlds[random])
	for len(g.worldToFind) >= 8 || len(g.worldToFind) <= 2 || g.stringInSlice('-', g.worldToFind) {
		random := rand.Intn(len(g.listWorlds))

		g.worldToFind = []rune(g.listWorlds[random])
	}
	g.lenght = len(g.worldToFind)
	for i := 0; i < g.lenght; i++ {
		g.worldFind = append(g.worldFind, ' ')
	}
}
