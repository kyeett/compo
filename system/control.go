package system

import (
	"fmt"

	"engo.io/ecs"
	"github.com/kyeett/compo/component"
)

type controlEntity struct {
	*ecs.BasicEntity
	*component.PlayerControlComponent
}

type InputSystem struct {
	entities []controlEntity
}

type AutoInputSystem struct {
	*InputSystem
	Events func() map[string]component.KeyState
}

// func NewInputSystem() InputSystem {
// 	return &InputSystem{}
// }

// Update updates all the entities in the InputSystem.
func (ac *AutoInputSystem) Update(dt float32) {

	for _, e := range ac.entities {
		for key := range e.PlayerControlComponent.KeyStates {
			delete(e.PlayerControlComponent.KeyStates, key)
		}

		// Loop through defined events, and mark as None if not triggered
		for eventName, mappedName := range e.PlayerControlComponent.Mapper {

			// Check if triggered "jump_1"
			state, ok := ac.Events()[eventName]
			switch ok {
			case true:
				e.PlayerControlComponent.KeyStates[mappedName] = state
			default:
				e.PlayerControlComponent.KeyStates[mappedName] = component.KeyStateNone
			}
		}
	}
}

func (c *InputSystem) Add(basic *ecs.BasicEntity, pc *component.PlayerControlComponent) {
	c.entities = append(c.entities, controlEntity{basic, pc})
}

func (c *InputSystem) Remove(basic ecs.BasicEntity) {
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

// Update updates all the entities in the InputSystem.
func (c *InputSystem) Update(dt float32) {
	// Translate Mouse.X and Mouse.Y into "game coordinates"
	// switch engo.CurrentBackEnd {
	// case engo.BackEndGLFW, engo.BackEndSDL, engo.BackEndVulkan:
	// 	m.mouseX = engo.Input.Mouse.X*m.camera.Z() + (m.camera.X()-(engo.GameWidth()/2)*m.camera.Z())/engo.GetGlobalScale().X
	// 	m.mouseY = engo.Input.Mouse.Y*m.camera.Z() + (m.camera.Y()-(engo.GameHeight()/2)*m.camera.Z())/engo.GetGlobalScale().Y
	// case engo.BackEndMobile:
	// 	m.mouseX = engo.Input.Mouse.X*m.camera.Z() + (m.camera.X()-(engo.GameWidth()/2)*m.camera.Z()+(engo.ResizeXOffset/2))/engo.GetGlobalScale().X
	// 	m.mouseY = engo.Input.Mouse.Y*m.camera.Z() + (m.camera.Y()-(engo.GameHeight()/2)*m.camera.Z()+(engo.ResizeYOffset/2))/engo.GetGlobalScale().Y
	// case engo.BackEndWeb:
	// 	m.mouseX = engo.Input.Mouse.X*m.camera.Z() + (m.camera.X()-(engo.GameWidth()/2)*m.camera.Z()+(engo.ResizeXOffset/2))/engo.GetGlobalScale().X
	// 	m.mouseY = engo.Input.Mouse.Y*m.camera.Z() + (m.camera.Y()-(engo.GameHeight()/2)*m.camera.Z()+(engo.ResizeYOffset/2))/engo.GetGlobalScale().X
	// }

	// // Rotate if needed
	// if m.camera.angle != 0 {
	// 	sin, cos := math.Sincos(m.camera.angle * math.Pi / 180)
	// 	m.mouseX, m.mouseY = m.mouseX*cos+m.mouseY*sin, m.mouseY*cos-m.mouseX*sin
	// }

	for _, e := range c.entities {
		fmt.Println(e)
		// Reset all values except these
		// *e.MouseComponent = MouseComponent{
		// 	Track:                e.MouseComponent.Track,
		// 	Hovered:              e.MouseComponent.Hovered,
		// 	startedDragging:      e.MouseComponent.startedDragging,
		// 	rightStartedDragging: e.MouseComponent.rightStartedDragging,
		// }

		// *e.PlayerControlComponent =

		// if e.MouseComponent.Track {
		// 	// track mouse position so that systems that need to stay on the mouse
		// 	// position can do it (think an RTS when placing a new building and
		// 	// you get a ghost building following your mouse until you click to
		// 	// place it somewhere in your world.
		// 	e.MouseComponent.MouseX = m.mouseX
		// 	e.MouseComponent.MouseY = m.mouseY
		// }

		// mx := m.mouseX
		// my := m.mouseY

		// if e.SpaceComponent == nil {
		// 	continue // with other entities
		// }

		// if e.RenderComponent != nil {
		// 	// Hardcoded special case for the HUD | TODO: make generic instead of hardcoding
		// 	if e.RenderComponent.shader == HUDShader || e.RenderComponent.shader == LegacyHUDShader {
		// 		mx = engo.Input.Mouse.X
		// 		my = engo.Input.Mouse.Y
		// 	}

		// 	if e.RenderComponent.Hidden {
		// 		continue // skip hidden components
		// 	}
		// }

		// // If the Mouse component is a tracker we always update it
		// // Check if the X-value is within range
		// // and if the Y-value is within range
		// if e.MouseComponent.Track || e.MouseComponent.startedDragging ||
		// 	e.SpaceComponent.Contains(engo.Point{X: mx, Y: my}) {

		// 	e.MouseComponent.Enter = !e.MouseComponent.Hovered
		// 	e.MouseComponent.Hovered = true
		// 	e.MouseComponent.Released = false

		// 	if !e.MouseComponent.Track {
		// 		// If we're tracking, we've already set these
		// 		e.MouseComponent.MouseX = mx
		// 		e.MouseComponent.MouseY = my
		// 	}

		// 	switch engo.Input.Mouse.Action {
		// 	case engo.Press:
		// 		switch engo.Input.Mouse.Button {
		// 		case engo.MouseButtonLeft:
		// 			e.MouseComponent.Clicked = true
		// 			e.MouseComponent.startedDragging = true
		// 		case engo.MouseButtonRight:
		// 			e.MouseComponent.RightClicked = true
		// 			e.MouseComponent.rightStartedDragging = true
		// 		}

		// 		m.mouseDown = true
		// 	case engo.Release:
		// 		switch engo.Input.Mouse.Button {
		// 		case engo.MouseButtonLeft:
		// 			e.MouseComponent.Released = true
		// 		case engo.MouseButtonRight:
		// 			e.MouseComponent.RightReleased = true
		// 		}
		// 	case engo.Move:
		// 		if m.mouseDown && e.MouseComponent.startedDragging {
		// 			e.MouseComponent.Dragged = true
		// 		}
		// 		if m.mouseDown && e.MouseComponent.rightStartedDragging {
		// 			e.MouseComponent.RightDragged = true
		// 		}
		// 	}
		// } else {
		// 	if e.MouseComponent.Hovered {
		// 		e.MouseComponent.Leave = true
		// 	}

		// 	e.MouseComponent.Hovered = false
		// }

		// if engo.Input.Mouse.Action == engo.Release {
		// 	// dragging stops as soon as one of the currently pressed buttons
		// 	// is released
		// 	e.MouseComponent.Dragged = false
		// 	e.MouseComponent.startedDragging = false
		// 	// TODO maybe separate out the release into left-button release and right-button release
		// 	e.MouseComponent.rightStartedDragging = false
		// 	// mouseDown goes false as soon as one of the pressed buttons is
		// 	// released. Effectively ending any dragging
		// 	m.mouseDown = false
		// }

		// // propagate the modifiers to the mouse component so that game
		// // implementers can take different decisions based on those
		// e.MouseComponent.Modifier = engo.Input.Mouse.Modifer
	}
}
