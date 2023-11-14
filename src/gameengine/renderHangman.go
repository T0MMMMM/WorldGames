package gameengine

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) renderHangman() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawTexture(g.texture.fondDegrade, 0, 0, rl.White)
	rl.DrawTextEx(g.font, "HANGMAN", rl.NewVector2(float32(screenWidth/2-125), 10), 60, 0, rl.Black)
	if g.run {
		// affiche chaque étape du pendu
		drawHangman(g.tryNum)
		g.drawLetters()
		for a := 0; a < len(g.letters); a++ {
			rl.DrawTextEx(g.font, string(g.letters[a]), rl.NewVector2((screenWidth - 120 + float32(len(g.letters)*80)) / 2 +  - float32(a*70), 750), 50, 0, rl.Red)
		}
	} else {
		if g.win {
			drawHangman(g.tryNum)
			g.drawText()
			for a := 0; a < len(g.letters); a++ {
				rl.DrawTextEx(g.font, string(g.letters[a]), rl.NewVector2((screenWidth - 120 + float32(len(g.letters)*80)) / 2 +  - float32(a*70), 750), 50, 0, rl.Red)
			}
			for j := 0; j < len(g.letterTest); j++ {	
				rl.DrawTextEx(g.font, string(g.letterTest[j]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(g.letterTestindice[j]*spaceBetweenSquare+g.letterTestindice[j]*squareSize)+17), float32(450+((squareSize+spaceBetweenSquare))+8)), 50, 0, rl.Green)
			}
			rl.DrawTextEx(g.font, "Lets go !", rl.NewVector2(float32(screenWidth/2-g.texture.keybord.Width/2)+230, float32(screenHeight)/1.2-float32(g.texture.keybord.Height/2)+20), 50, 0, rl.Black)
		} else {
			//si on a perdu, le pendu est encore affiché
			drawHangman(g.tryNum)
			g.drawText()
			g.drawLetters()
			g.printWord()
		}
	}	
	rl.EndDrawing()
}




func drawHangman(try int32) {
	if try >= 1 {
		rl.DrawRectangle(660, 400, 180, 10, rl.Black)}
	if try >= 2 {
		rl.DrawRectangle(750, 150, 10, 250, rl.Black) }
	if try >= 3 {
		rl.DrawRectangle(750, 150, 160, 10, rl.Black) }
	if try >= 4 {
		rl.DrawLineEx((rl.NewVector2(755.0, 200.0)), (rl.NewVector2(800.0, 155.0)), 10.0, rl.Black)}
	if try >= 5 {
		rl.DrawRectangle(880, 150, 10, 50, rl.Black)}
	if try >= 6 {
		rl.DrawCircle(885, 225, 30, rl.Black)}
	if try >= 7 {
		rl.DrawRectangle(880, 210, 10, 120, rl.Black)}
	if try >= 8 {
		rl.DrawRectangle(845, 275, 80, 10, rl.Black)}
	if try >= 9 {
		rl.DrawLineEx((rl.NewVector2(886.0, 320.0)), (rl.NewVector2(852.0, 370.0)), 10.0, rl.Black)}
	if try >= 10 {
		rl.DrawLineEx((rl.NewVector2(884.0, 320.0)), (rl.NewVector2(918.0, 370.0)), 10.0, rl.Black)}
}


func(g *EngineStruct) drawText() {
	for i := 0; i < g.lenght; i++ {
		rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 600, squareSize, 5, rl.Black)			
	}
}

func(g *EngineStruct) drawLetters() {
	for i := 0; i < g.lenght; i++ {
		rl.DrawRectangle(int32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(i*spaceBetweenSquare+i*squareSize)), 600, squareSize, 5, rl.Black)			
		if len(g.letterTest) > 0 {
			for j := 0; j < len(g.letterTest); j++ {	
				rl.DrawTextEx(g.font, string(g.letterTest[j]), rl.NewVector2(float32(screenWidth/2-((squareSize*g.lenght+spaceBetweenSquare*g.lenght-1)/2)+(g.letterTestindice[j]*spaceBetweenSquare+g.letterTestindice[j]*squareSize)+17), float32(450+((squareSize+spaceBetweenSquare))+8)), 50, 0, rl.Black)
			}
		}
	}
}