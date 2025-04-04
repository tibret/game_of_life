package main

import (
	"game_of_life/game"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenSize = 1024
	rl.InitWindow(screenSize, screenSize, "game of life")
	rl.SetTargetFPS(60)

	g := game.NewGame(128)
	blockSize := int32(screenSize / g.Size)

	setup := true
	oneFrame := false

	paintStarted := false
	paintMode := false

	displayCounter := 0
	displayMessage := ""

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		if rl.IsKeyPressed(rl.KeySpace) {
			setup = !setup
		}

		if rl.IsKeyPressed(rl.KeyN) {
			oneFrame = true
		}

		if rl.IsKeyPressed(rl.KeyC) {
			setup = true
			g.Clear()
		}

		if rl.IsKeyPressed(rl.KeyF5) && setup {
			g.Save()
			displayCounter = 100
			displayMessage = "State Saved"
		} else if rl.IsKeyDown(rl.KeyF6) && setup {
			g.Load()
			displayCounter = 100
			displayMessage = "State Loaded"
		}

		if displayCounter > 0 {
			displayCounter--
			rl.DrawText(displayMessage, 0, 0, 24, rl.Red)
		}

		if rl.IsMouseButtonDown(rl.MouseButtonRight) && !paintStarted {
			blockX, blockY := MouseToBlockCoord(blockSize, g.Size)

			paintStarted = true
			paintMode = !g.Cells[blockX][blockY]
		} else if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
			paintStarted = false
		}

		if !setup || oneFrame {
			g.Advance()
		} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			blockX, blockY := MouseToBlockCoord(blockSize, g.Size)
			g.Cells[blockX][blockY] = !g.Cells[blockX][blockY]
		} else if rl.IsMouseButtonDown(rl.MouseButtonRight) {
			blockX, blockY := MouseToBlockCoord(blockSize, g.Size)
			g.Cells[blockX][blockY] = paintMode
		}

		DrawGame(g, blockSize)

		oneFrame = false

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func DrawGame(g *game.Game, blockSize int32) {

	for i := int32(0); i < g.Size; i++ {
		for j := int32(0); j < g.Size; j++ {
			if g.Cells[i][j] {
				rl.DrawRectangle(i*blockSize, j*blockSize, blockSize, blockSize, rl.RayWhite)
			}
		}
	}
}

func MouseToBlockCoord(blockSize int32, gameSize int32) (int32, int32) {
	mousePosition := rl.GetMousePosition()
	blockX := int32(mousePosition.X / float32(blockSize))
	blockY := int32(mousePosition.Y / float32(blockSize))

	if blockX < 0 {
		blockX = 0
	}
	if blockY < 0 {
		blockY = 0
	}
	if blockX >= gameSize {
		blockX = gameSize - 1
	}
	if blockY >= gameSize {
		blockY = gameSize - 1
	}

	return blockX, blockY
}
