package GoHashGrid

import (
	"errors"
	"math"
)

type Cell bool

type Grid struct {
	CellSize float32
	Cells    map[[3]int]Cell
}

var (
	aigErr = errors.New("Point was already in grid")
)

func floatToIntId(num, cellSize float32) int {
	return int(math.Floor(float64(num / cellSize)))
}

func (g *Grid) IntIdToVector3(uid [3]int) Vector3 {
	return Vector3{
		X: float32(uid[0]) * g.CellSize,
		Y: float32(uid[1]) * g.CellSize,
		Z: float32(uid[2]) * g.CellSize,
	}
}

func NewGrid(totalSize, cellSize float32) Grid {
	cells := make(map[[3]int]Cell)
	for x := -totalSize / 2; x <= totalSize/2; x += cellSize {
		xi := floatToIntId(x, cellSize)

		for y := -totalSize / 2; y <= totalSize/2; y += cellSize {
			yi := floatToIntId(y, cellSize)

			for z := -totalSize / 2; z <= totalSize/2; z += cellSize {
				zi := floatToIntId(z, cellSize)

				cells[[3]int{xi, yi, zi}] = false
			}
		}
	}
	g := Grid{
		CellSize: cellSize,
		Cells:    cells,
	}
	return g
}

func (g *Grid) GetUidToPoint(point *Vector3) [3]int {
	x := int(math.Floor(float64(point.X / g.CellSize)))
	y := int(math.Floor(float64(point.Y / g.CellSize)))
	z := int(math.Floor(float64(point.Z / g.CellSize)))
	return [3]int{x, y, z}
}

func (g *Grid) Insert(point *Vector3) ([3]int, error) {
	xyz := g.GetUidToPoint(point)

	cell := g.Cells[xyz]
	if cell == true {
		return xyz, aigErr
	} else {
		g.Cells[xyz] = true
		return xyz, nil
	}

}

func (g *Grid) indexIsValid(index [3]int) bool {
	_, ok := g.Cells[index]
	return ok
}

func (g *Grid) GetNeighbors(x, y, z int, hasToBeFree bool) map[[3]int]Cell {
	neighbours := make(map[[3]int]Cell)
	for xi := -1; xi < 2; xi += 2 {
		for yi := -1; yi < 2; yi += 2 {
			for zi := -1; zi < 2; zi += 2 {
				neighborX := x + xi
				neighborY := y + yi
				neighborZ := z + zi

				if g.indexIsValid([3]int{neighborX, neighborY, neighborZ}) {
					occupied := g.Cells[[3]int{neighborX, neighborY, neighborZ}]
					if hasToBeFree && bool(!occupied) {
						neighbours[[3]int{neighborX, neighborY, neighborZ}] = occupied
					} else {
						neighbours[[3]int{neighborX, neighborY, neighborZ}] = occupied
					}
				}
			}
		}
	}
	return neighbours
}
