package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) update() {

	g.run = !rl.WindowShouldClose()

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
