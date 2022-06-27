package main

import "math/rand"

func (g *Game) EventDropToy() {
	if g.count != 0 && g.count%uint64((100+rand.Intn(100))) == 0 && g.toys.Spawn < maxToys {
		g.toys.Spawn += 1
	}
}
