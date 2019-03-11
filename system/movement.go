package system

import (
	"engo.io/ecs"
	"github.com/kyeett/compo/component"
)

type movementEntity struct {
	*ecs.BasicEntity
	*component.TransformComponent
	*component.RigidBodyComponent
}

type MovementSystem struct {
	entities []movementEntity
}

func (s *MovementSystem) Add(basic *ecs.BasicEntity, t *component.TransformComponent, rb *component.RigidBodyComponent) {
	s.entities = append(s.entities, movementEntity{basic, t, rb})
}

func (s *MovementSystem) Remove(basic ecs.BasicEntity) {
	var delete = -1
	for index, entity := range s.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *MovementSystem) Update(dt float32) {
	for _, e := range s.entities {
		v := e.RigidBodyComponent.Velocity.Scaled(float64(dt))
		e.TransformComponent.Position = e.TransformComponent.Position.Add(v)
	}
}
