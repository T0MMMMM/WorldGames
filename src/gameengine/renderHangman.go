package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) renderHangman() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)

	rl.DrawTextEx(g.font, "HANGMAN", rl.NewVector2(float32(screenWidth/2-135), 10), 60, 0, rl.Black)
	if g.run {
		for i := 0; i < g.lenght; i++ {
			rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 500, squareSize, 5, rl.Black)			
			if len(g.letterTest) > 0 {
				for j := 0; j < len(g.letterTest); j++ {	
					rl.DrawTextEx(g.font, string(g.letterTest[j]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(g.letterTestindice[j]*spaceBetweenSquare+g.letterTestindice[j]*squareSize)+17), float32(350+((squareSize+spaceBetweenSquare))+8)), 50, 0, rl.Black)
				}
			}
		}
		for a := 0; a < len(g.letters); a++ {
			rl.DrawTextEx(g.font, string(g.letters[a]), rl.NewVector2((screenWidth - 120 + float32(len(g.letters)*80)) / 2 +  - float32(a*70), 700), 50, 0, rl.Red)
		}

		if g.tryNum >= 1 {
			rl.DrawRectangle(700, 350, 100, 5, rl.Black)}
		if g.tryNum >= 2 {
			rl.DrawRectangle(750, 200, 5, 150, rl.Black) }
		if g.tryNum >= 3 {
			rl.DrawRectangle(750, 200, 150, 5, rl.Black) }
		
	} else {
		if g.win {
			for i := 0; i < g.lenght; i++ {
				rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 500, squareSize, 5, rl.Black)			
				if len(g.letterTest) > 0 {
					for j := 0; j < len(g.letterTest); j++ {	
						rl.DrawTextEx(g.font, string(g.letterTest[j]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(g.letterTestindice[j]*spaceBetweenSquare+g.letterTestindice[j]*squareSize)+17), float32(350+((squareSize+spaceBetweenSquare))+8)), 50, 0, rl.Green)
					}
				}
			}
			rl.DrawTextEx(g.font, "~~ GG T BO ~~", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+160, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)-75), 50, 0, rl.Black)
		} else {
			rl.DrawTextEx(g.font, "~~ T NUL BOUFFON ~~", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+95, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)-75), 50, 0, rl.Black)
		}
	}

	
	rl.EndDrawing()
}