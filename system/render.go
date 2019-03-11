package system

import (
	"engo.io/ecs"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/kyeett/compo/component"
	"golang.org/x/image/colornames"
)

type renderEntity struct {
	*ecs.BasicEntity
	*component.TransformComponent
}

type RenderSystem struct {
	entities []renderEntity
}

func (s *RenderSystem) Add(basic *ecs.BasicEntity, t *component.TransformComponent) {
	s.entities = append(s.entities, renderEntity{basic, t})
}

func (s *RenderSystem) Remove(basic ecs.BasicEntity) {
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

func (s *RenderSystem) Update(dt float32) {}

func (s *RenderSystem) Render(screen *ebiten.Image) {
	for _, e := range s.entities {
		ebitenutil.DrawRect(screen, e.Position.X, e.Position.Y, 3, 3, colornames.Red)
	}
}
