package game

type Game struct {
	Cells      [][]bool
	Buffer     [][]bool
	save       [][]bool
	Size       int32
	Generation int
}

func (g *Game) Clear() {
	for i := range g.Size {
		for k := range g.Size {
			g.Cells[i][k] = false
		}
	}
}

func (g *Game) Save() {
	for i := range g.Size {
		for j := range g.Size {
			g.save[i][j] = g.Cells[i][j]
		}
	}
}

func (g *Game) Load() {
	for i := range g.Size {
		for j := range g.Size {
			g.Cells[i][j] = g.save[i][j]
		}
	}

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

	return &Game{Cells: c, Generation: 0, Size: int32(size), Buffer: buffer, save: save}
}
