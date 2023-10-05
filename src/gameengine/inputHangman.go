package gameengine

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) inputHangman() {
	
	for i := 0; i < 26; i++ {
		g.inputLetter = []int32{rl.KeyQ, rl.KeyB, rl.KeyC, rl.KeyD, rl.KeyE, rl.KeyF, rl.KeyG, rl.KeyH, rl.KeyI, rl.KeyJ, rl.KeyK, rl.KeyL, rl.KeySemicolon, rl.KeyN, rl.KeyO, rl.KeyP, rl.KeyA, rl.KeyR, rl.KeyS, rl.KeyT, rl.KeyU, rl.KeyV, rl.KeyZ, rl.KeyX, rl.KeyY, rl.KeyW}
		if rl.IsKeyPressed(g.inputLetter[i]) {
			g.letter[i] = true
			fmt.Println(g.letter[i], g.letter, g.tryNum)
		}
	}
	if g.tryNum >= 10 {
		g.win = false
		g.run = false
		fmt.Println("aaaaaaaaa")
	}
	if len(g.letterTest) == len(g.worldToFind) {
		g.win = true
		g.run = false
		fmt.Println("bbbbbbbb")
	}

	for i := 0; i < len(g.letter); i++ {
		if g.letter[i] {
			fmt.Println(g.letters)
			fmt.Println(g.letterTest)
			for j := 0; j < len(g.worldToFind); j++ {
				if rune(65+i) == rune(g.worldToFind[j]) {
					g.letterTest = append(g.letterTest, 65+i)
					g.letterTestindice = append(g.letterTestindice, j)
				}
			}
			g.letters = append(g.letters, (65+i))
		}
		g.letters = removeDuplicate(g.letters)
		g.letterTest = removeDuplicate(g.letterTest)
		g.letterTestindice = removeDuplicate(g.letterTestindice)
		g.tryNum = int32(len(g.letters) - len(g.letterTest))
	}

	if !g.run { 
		fmt.Println("abc")
	}
}


func removeDuplicate(arr []int) []int {
	map_var := map[int]bool{}
	result := []int{}
	for e := range arr {
	   if map_var[arr[e]] != true {
		  map_var[arr[e]] = true
		  result = append(result, arr[e])
	   }
	}
	return result
 }