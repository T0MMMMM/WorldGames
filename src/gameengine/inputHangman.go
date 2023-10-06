package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) inputHangman() {
	
	for i := 0; i < 26; i++ {
		g.inputLetter = []int32{rl.KeyQ, rl.KeyB, rl.KeyC, rl.KeyD, rl.KeyE, rl.KeyF, rl.KeyG, rl.KeyH, rl.KeyI, rl.KeyJ, rl.KeyK, rl.KeyL, rl.KeySemicolon, rl.KeyN, rl.KeyO, rl.KeyP, rl.KeyA, rl.KeyR, rl.KeyS, rl.KeyT, rl.KeyU, rl.KeyV, rl.KeyZ, rl.KeyX, rl.KeyY, rl.KeyW}
		if rl.IsKeyPressed(g.inputLetter[i]) {
			for k := 0; k < len(g.usedLetters); k++ {
				if g.usedLetters[k] == g.inputLetter[i] {
					break
				} else if k == len(g.usedLetters)-1 {
					g.letter[i] = true
					g.usedLetters = append(g.usedLetters, g.inputLetter[i])
				}
			}
		}
	}
	if g.tryNum >= 10 {
		g.win = false
		g.run = false
	
	}
	if len(g.letterTest) == len(g.worldToFind) {
		g.win = true
		g.run = false
	}

	for i := 0; i < len(g.letter); i++ {
		if g.letter[i] {
			for j := 0; j < len(g.worldToFind); j++ {
				if rune(65+i) == rune(g.worldToFind[j]) {
					g.letterTest = append(g.letterTest, 65+i)
					g.letterTestindice = append(g.letterTestindice, j)
					g.lettersWithoutDoble = append(g.lettersWithoutDoble, 65+i)
				}
			}
			g.letters = append(g.letters, (65+i))
		}
		g.lettersWithoutDoble = removeDuplicate(g.lettersWithoutDoble)
		g.tryNum = int32(len(g.letters) - len(g.lettersWithoutDoble))
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