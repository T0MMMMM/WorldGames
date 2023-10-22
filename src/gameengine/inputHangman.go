package gameengine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *EngineStruct) inputHangman() {
	/* Cette boucle permet de tester les touches de clavier pressés, si la lettre a déja été pressé, cette dernière ne fonctionnera plus dans la suite du programme */
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
	/* Si le nombre d'essaie est supérieur ou égale à 10 : la boucle s'arrète et l'utilisateur a perdu la partie */
	if g.tryNum >= 10 {
		g.win = false
		g.run = false

	}
	/* Si la longueur du nombre de lettres trouvées est égale aux nombre de lettres du mot à trouver : la boucle s'arrète et l'utilisateur a gagné la partie */
	if len(g.letterTest) == len(g.worldToFind) {
		g.win = true
		g.run = false
	}
	/* Cette boucle permet plusieurs choses :
	Si la lettre pressé est dans le mot : on ajoute dans un tableau le(s) indice(s) de position de la lettre dans le mot
										  on ajoute dans un tableau la lettre (plusieurs fois si cette lettre a des doublés)
	*/
	for i := 0; i < len(g.letter); i++ {
		if g.letter[i] {
			for j := 0; j < len(g.worldToFind); j++ {
				if rune(65+i) == rune(g.worldToFind[j]) {
					g.letterTest = append(g.letterTest, 65+i)
					g.letterTestindice = append(g.letterTestindice, j)
					g.lettersWithoutDoble = append(g.lettersWithoutDoble, 65+i)
				}
			}
			g.letters = append(g.letters, (65 + i))
		}
		g.lettersWithoutDoble = removeDuplicate(g.lettersWithoutDoble)
		g.tryNum = int32(len(g.letters) - len(g.lettersWithoutDoble))
	}
}

func removeDuplicate(liste []int) []int {
	/* Cette fonction permet de supprimer tous les doublons d'un tableau */
	for i := 0; i < len(liste)-1; i++ {
		for j := i + 1; j < len(liste); j++ {
			if liste[i] == liste[j] {
				liste = append(liste[:j], liste[j+1:]...)
				j--
			}
		}
	}
	return liste
}
