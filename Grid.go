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
	nnoErr = errors.New("No Neighbour in a 5% radius(hamilton) was occupied")
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

func (g *Grid) neighbourCheck(neighbours map[[3]int]Cell, neighbourUid [3]int, hasToBeFree bool) map[[3]int]Cell {
	if g.indexIsValid(neighbourUid) {
		occupied := g.Cells[neighbourUid]
		if hasToBeFree && !bool(occupied) {
			neighbours[neighbourUid] = occupied
		} else {
			neighbours[neighbourUid] = occupied
		}
	}
	return neighbours
}

func (g *Grid) GetNeighbors(x, y, z int, hasToBeFree bool) map[[3]int]Cell {
	neighbours := make(map[[3]int]Cell)

	for xi := -1; xi < 2; xi += 2 {
		neighbours = g.neighbourCheck(neighbours, [3]int{x + xi, y, z}, hasToBeFree)
	}
	for yi := -1; yi < 2; yi += 2 {
		neighbours = g.neighbourCheck(neighbours, [3]int{x, y + yi, z}, hasToBeFree)
	}
	for zi := -1; zi < 2; zi += 2 {
		neighbours = g.neighbourCheck(neighbours, [3]int{x, y, z + zi}, hasToBeFree)
	}

	return neighbours
}

func (g *Grid) NextOccupiedNeighbour(uid [3]int) ([3]int, error) {
	searchRadius := 1
	alreadySearched := make(map[[3]int]bool)
	alreadySearched[uid] = true
	found := false

	for !found {
		//if search radius grows larger than 5% of the grids total size then cancel
		if searchRadius > int(math.Pow(float64(len(g.Cells)), 1.0/3.0)*0.05) {
			return [3]int{0, 0, 0}, nnoErr
		}
		for xi := uid[0] - searchRadius; xi < uid[0]+searchRadius; xi++ {
			for yi := uid[1] - searchRadius; yi < uid[1]+searchRadius; yi++ {
				for zi := uid[2] - searchRadius; zi < uid[2]+searchRadius; zi++ {
					currentId := [3]int{xi, yi, zi}

					_, alreadyInList := alreadySearched[currentId]
					if alreadyInList {
						continue
					} else if !g.indexIsValid(currentId) {
						alreadySearched[currentId] = false
						continue
					} else if g.Cells[currentId] {
						found = true
						return currentId, nil
					}
				}
			}
		}
		searchRadius += 1
	}
	//obligatory return which can never be reached
	return [3]int{0, 0, 0}, nil
}
