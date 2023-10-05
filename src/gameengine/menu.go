package gameengine 

import ( "github.com/gen2brain/raylib-go/raylib" ) 

func (g *EngineStruct) menu() { 
	rl.BeginDrawing() 
	rl.ClearBackground(rl.White) 
	rl.DrawTexture(g.texture.wordle, 20, screenHeight/2 - g.texture.wordle.Height/2, rl.White) 
	rl.DrawTexture(g.texture.hangman, screenWidth/2 - g.texture.hangman.Width/2, screenHeight/2 - g.texture.hangman.Height/2, rl.White) 

	//motus
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (rl.GetMouseX() > 20 && rl.GetMouseX() < 20 + g.texture.wordle.Width && rl.GetMouseY() > screenHeight/2 - g.texture.wordle.Height/2 && rl.GetMouseY() < screenHeight/2 + g.texture.wordle.Height/2) { 
		g.numberGame = "Wordle" 
		g.showMenu = false 
	}
	//hangman
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) && (rl.GetMouseX() > screenWidth/2 - g.texture.hangman.Width/2 && rl.GetMouseX() < screenWidth/2 + g.texture.hangman.Width/2 && rl.GetMouseY() > screenHeight/2 - g.texture.hangman.Height/2 && rl.GetMouseY() < screenHeight/2 + g.texture.hangman.Height/2) { 
		g.numberGame = "Hangman"
		g.showMenu = false 
	} 
	rl.EndDrawing() 
}