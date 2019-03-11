package component

import (
	"github.com/peterhellberg/gfx"
)

type TransformComponent struct {
	Position gfx.Vec
	Rotation float64
	Scale    gfx.Vec
}

type RigidBodyComponent struct {
	Mass       float64
	Drag       float64
	UseGravity bool
	Velocity   gfx.Vec
}

type KeyState int

const (
	KeyStateNone KeyState = iota
	KeyStatePressed
	KeyStateReleased
	KeyStateJustPressed
	KeyStateJustReleased
)

type PlayerControlComponent struct {
	Mapper    map[string]string
	KeyStates map[string]KeyState
}

func (c *PlayerControlComponent) GetPlayerControlComponent() *PlayerControlComponent {
	return c
}
