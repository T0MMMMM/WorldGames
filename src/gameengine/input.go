package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) input() {
	for i := 0; i < 26; i++ {
		g.inputLetter = []int32{rl.KeyQ, rl.KeyB, rl.KeyC, rl.KeyD, rl.KeyE, rl.KeyF, rl.KeyG, rl.KeyH, rl.KeyI, rl.KeyJ, rl.KeyK, rl.KeyL, rl.KeySemicolon, rl.KeyN, rl.KeyO, rl.KeyP, rl.KeyA, rl.KeyR, rl.KeyS, rl.KeyT, rl.KeyU, rl.KeyV, rl.KeyZ, rl.KeyX, rl.KeyY, rl.KeyW}
		if rl.IsKeyPressed(g.inputLetter[i]) {
			g.letter[i] = true
		}
	}
	if rl.IsKeyPressed(rl.KeyEnter) && !g.stringInSlice(' ', g.worldFind) {

		for i := 0; i < len(g.listWorlds); i++ {
			if g.testEq(g.worldFind, []rune(g.listWorlds[i])) {
				if g.testEq(g.worldToFind, g.worldFind) {
					g.win = true
					g.run = false
				}
				if g.tryNum > 4 {
					g.run = false
				}
		
				g.tryNum += 1
				g.listeWorldFind = append(g.listeWorldFind, []string{string(g.worldFind)})// ajouter au début
		
				for i := 0; i < g.lenght; i++ {
					if g.worldFind[i] == g.worldToFind[i] {
						g.listeWorldFind[len(g.listeWorldFind)-1] = append(g.listeWorldFind[len(g.listeWorldFind)-1], "green")
						if g.count(g.worldFind[i], g.letterValid) != g.count(g.worldFind[i], g.worldToFind) {
							add := true
							for j := 0; j < len(g.posLetterValid); j++ {
								if i == g.posLetterValid[j] {
									add = false
								}
							}
							if add {
								g.letterValid = append(g.letterValid, g.worldFind[i])
								g.posLetterValid = append(g.posLetterValid, i)
							}
						}
					} else if g.stringInSlice(g.worldFind[i], g.worldToFind) && g.count(g.worldFind[i], g.letterValid) != g.count(g.worldFind[i], g.worldToFind) && g.inListeFind(i, g.listeWorldFind) {
						g.listeWorldFind[len(g.listeWorldFind)-1] = append(g.listeWorldFind[len(g.listeWorldFind)-1], "yellow")
					} else {
						g.listeWorldFind[len(g.listeWorldFind)-1] = append(g.listeWorldFind[len(g.listeWorldFind)-1], "grey")
					}
				}
		
				for i := 0; i < len(g.listeWorldFind[len(g.listeWorldFind)-1])-1; i++ {
					if g.count(rune(g.listeWorldFind[len(g.listeWorldFind)-1][0][i]), g.letterValid) == g.count(rune(g.listeWorldFind[len(g.listeWorldFind)-1][0][i]), g.worldToFind) {
						for j := 0; j < len(g.listeWorldFind[len(g.listeWorldFind)-1])-1; j++ {
							if g.listeWorldFind[len(g.listeWorldFind)-1][j+1] == "yellow" && g.listeWorldFind[len(g.listeWorldFind)-1][0][j] == g.listeWorldFind[len(g.listeWorldFind)-1][0][i] {
								g.listeWorldFind[len(g.listeWorldFind)-1][j+1] = "grey"
							}
						}
					}
				}
		
				g.worldFind = []rune{}
				for i := 0; i < g.lenght; i++ {
					g.worldFind = append(g.worldFind, ' ')
				}
				break
			}
		}
		
	}

	if rl.IsKeyPressed(rl.KeyBackspace) {
		for i := g.lenght-1; i >= 0; i-- {
			if g.worldFind[i] != ' ' {
				g.worldFind[i] = ' '
				break
			}
		}
	}

}
