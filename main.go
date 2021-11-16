package main

import (
	"math/rand"
)

func main() {
	g := NewGrid(30, 0.1)

	for i := 0; i < 10000; i++ {
		RandomInsert(&g)
	}


}

func RandomFloatInRange(low float32, high float32) float32 {
	return low + (high-low)*rand.Float32()
}

func RandomInsert(g *Grid) {
	randomPoint := Vector3{
		X: RandomFloatInRange(-15, 15),
		Y: RandomFloatInRange(-15, 15),
		Z: RandomFloatInRange(-15, 15),
	}
	g.Insert(&randomPoint)
}
