package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) updateWordle() {

	g.run = !rl.WindowShouldClose()
	// on place la lettre dans le jeu pour tester un mot
	for i := 0; i < 26; i++ {
		if g.letter[i] {
			for j := 0; j < g.lenght; j++ {
				if g.worldFind[j] == ' ' {
					g.worldFind[j] = rune(65+i)
					break
				}
			}
		}
	}

	for i := 0; i < 26; i++ {
		g.letter[i] = false
	}
	
}
