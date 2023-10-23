package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) renderWordle() {
	rl.BeginDrawing()
	rl.ClearBackground(g.bgColor)

	// affiche le titre
	rl.DrawTextEx(g.font, "WORDLE", rl.NewVector2(float32(screenWidth/2-85), 10), 60, 0, rl.White)
	
	// affiche le mot en cours (le mot qu'il est en train d'essayer)
	if g.run {
		for i := 0; i < g.lenght; i++ {
			rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 75+(g.tryNum*(squareSize+spaceBetweenSquare)), squareSize, squareSize, rl.Black)
			rl.DrawTextEx(g.font, string(g.worldFind[i]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)+17), float32(75+(g.tryNum*(squareSize+spaceBetweenSquare))+8)), 50, 0, rl.White)
		}
	} else {
		if g.win {
			rl.DrawTextEx(g.font, "~~ GG T BO ~~", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+160, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)-75), 50, 0, rl.White)
		} else {
			rl.DrawTextEx(g.font, "~~ T NUL BOUFFON ~~", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+95, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)-75), 50, 0, rl.White)
		}
	}

	// affiche les mots déjà testés
	for j := 0; j < len(g.listeWorldFind); j++ {
		for i := 0; i < len(g.listeWorldFind[j][0]); i++ {
			rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 75+(int32(j)*(squareSize+spaceBetweenSquare)), squareSize, squareSize, g.color[g.listeWorldFind[j][i+1]])
			rl.DrawTextEx(g.font, string(g.listeWorldFind[j][0][i]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)+17), float32(75+(int32(j)*(squareSize+spaceBetweenSquare))+8)), 50, 0, rl.White)
		}
	}

	rl.EndDrawing()
}
