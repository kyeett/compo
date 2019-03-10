package world

import (
	"github.com/kyeett/compo/system"
)

type World struct {
	systems []system.System
}

var timeStep = 0.01

func (w *World) Update(dt float32) {
	for _, s := range w.systems {
		s.Update(dt)
	}
}

func New() *World {
	w := World{}
	w.AddSystem(&system.ControlSystem{})
	return &w
}

// AddSystem adds the given System to the World, sorted by priority.
func (w *World) AddSystem(s system.System) {
	w.systems = append(w.systems, s)
}
