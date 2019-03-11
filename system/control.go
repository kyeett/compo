package system

import (
	"engo.io/ecs"
	"github.com/kyeett/compo/component"
)

type controlEntity struct {
	*ecs.BasicEntity
	*component.PlayerControlComponent
	*component.RigidBodyComponent
}

type ControlSystem struct {
	entities []controlEntity
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, pc *component.PlayerControlComponent, rb *component.RigidBodyComponent) {
	c.entities = append(c.entities, controlEntity{basic, pc, rb})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	var delete = -1
	for index, entity := range c.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		if e.KeyStates["jump"] == component.KeyStateJustPressed || e.KeyStates["jump"] == component.KeyStatePressed {
			e.RigidBodyComponent.Velocity.Y = -10
		}
	}
}
