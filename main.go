package main

import (
	"fmt"
	"math/rand"
)

func main() {
	g := NewGrid(1, 0.1)

	for i := 0; i < 10; i++ {
		RandomInsert(&g)
	}
}

func RandomFloatInRange(low float32, high float32) float32 {
	return low + (high-low)*rand.Float32()
}

func RandomInsert(g *Grid) {
	InRange := float32(1)
	randomPoint := Vector3{
		X: RandomFloatInRange(-InRange*0.5, InRange*0.5),
		Y: RandomFloatInRange(-InRange*0.5, InRange*0.5),
		Z: RandomFloatInRange(-InRange*0.5, InRange*0.5),
	}
	insert, err := g.Insert(&randomPoint)
	fmt.Println(randomPoint, insert, err)
}
