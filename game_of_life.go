package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenSize = 1024
	rl.InitWindow(screenSize, screenSize, "game of life")
	rl.SetTargetFPS(60)

	game := NewGame(128)
	blockSize := int32(screenSize / game.Size)

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
			game.Clear()
		}

		if rl.IsKeyPressed(rl.KeyF5) && setup {
			game.Save()
			displayCounter = 100
			displayMessage = "State Saved"
		} else if rl.IsKeyDown(rl.KeyF6) && setup {
			game.Load()
			displayCounter = 100
			displayMessage = "State Loaded"
		}

		if displayCounter > 0 {
			displayCounter--
			rl.DrawText(displayMessage, 0, 0, 24, rl.Red)
		}

		if rl.IsMouseButtonDown(rl.MouseButtonRight) && !paintStarted {
			blockX, blockY := MouseToBlockCoord(blockSize, game.Size)

			paintStarted = true
			paintMode = !game.cells[blockX][blockY]
		} else if rl.IsMouseButtonReleased(rl.MouseButtonRight) {
			paintStarted = false
		}

		if !setup || oneFrame {
			GameLogic(game)
		} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			blockX, blockY := MouseToBlockCoord(blockSize, game.Size)
			game.cells[blockX][blockY] = !game.cells[blockX][blockY]
		} else if rl.IsMouseButtonDown(rl.MouseButtonRight) {
			blockX, blockY := MouseToBlockCoord(blockSize, game.Size)
			game.cells[blockX][blockY] = paintMode
		}

		DrawGame(game, blockSize)

		oneFrame = false

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func NewGame(size int) *Game {
	c := make([][]bool, size)
	for i := range size {
		c[i] = make([]bool, size)
	}

	var buffer = make([][]bool, size)
	for i := range size {
		buffer[i] = make([]bool, size)
	}

	var save = make([][]bool, size)
	for i := range size {
		save[i] = make([]bool, size)
	}

	for i := range size {
		for j := range size {
			c[i][j] = false
			buffer[i][j] = false
			save[i][j] = false
		}
	}

	c[61][61] = true
	c[62][61] = true
	c[62][60] = true
	c[62][62] = true
	c[63][62] = true

	return &Game{cells: c, Generation: 0, Size: int32(size), buffer: buffer, save: save}
}

type Game struct {
	cells      [][]bool
	buffer     [][]bool
	save       [][]bool
	Size       int32
	Generation int
}

type Coord struct {
	X int32
	Y int32
}

func GameLogic(game *Game) {
	for i := int32(0); i < game.Size; i++ {
		for j := int32(0); j < game.Size; j++ {
			alive := game.cells[i][j]
			liveNeighbors := 0

			neighbors := Neighbors(i, j, game.Size)
			for _, n := range neighbors {
				// print("{", n.X, ", ", n.Y, "}, ")
				if game.cells[n.X][n.Y] {
					liveNeighbors++
				}
			}

			if alive {
				if liveNeighbors < 2 {
					alive = false
				} else if liveNeighbors > 3 {
					alive = false
				}
			} else {
				if liveNeighbors == 3 {
					alive = true
				}
			}

			game.buffer[i][j] = alive
		}
	}

	for i := int32(0); i < game.Size; i++ {
		for j := int32(0); j < game.Size; j++ {
			game.cells[i][j] = game.buffer[i][j]
		}
	}

	game.Generation = game.Generation + 1
}

var Neighbor_Offsets = [8][2]int32{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func Neighbors(i int32, j int32, size int32) []Coord {
	neighbors := make([]Coord, 8)
	for n := range 8 {
		offsets := Neighbor_Offsets[n]
		x := (i + size + offsets[0]) % size
		y := (j + size + offsets[1]) % size

		neighbors[n] = Coord{X: x, Y: y}
	}

	return neighbors
}

func DrawGame(game *Game, blockSize int32) {

	for i := int32(0); i < game.Size; i++ {
		for j := int32(0); j < game.Size; j++ {
			if game.cells[i][j] {
				rl.DrawRectangle(i*blockSize, j*blockSize, blockSize, blockSize, rl.RayWhite)
			}
		}
	}
}

func (g *Game) Clear() {
	for i := range g.Size {
		for k := range g.Size {
			g.cells[i][k] = false
		}
	}
}

func (g *Game) Save() {
	for i := range g.Size {
		for j := range g.Size {
			g.save[i][j] = g.cells[i][j]
		}
	}
}

func (g *Game) Load() {
	for i := range g.Size {
		for j := range g.Size {
			g.cells[i][j] = g.save[i][j]
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
