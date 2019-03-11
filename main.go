package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/peterhellberg/gfx"

	"engo.io/ecs"
	"github.com/kyeett/compo/component"
	"github.com/kyeett/compo/system"
)

type Player struct {
	ecs.BasicEntity
	component.TransformComponent
	component.PlayerControlComponent
	component.RigidBodyComponent
}

func main() {
	w := ecs.World{}
	// w := world.New()
	// w.AddSystem(&system.Input{})

	// Input = NewInputManager()

	// startTime := time.Now()

	f := func() map[string]component.KeyState {
		cmds := map[string]component.KeyState{}
		// since := time.Since(startTime).Nanoseconds()
		// fmt.Println(since / int64(time.Millisecond))

		cmds["jump_1"] = component.KeyStateJustPressed
		cmds["jump_2"] = component.KeyStateJustPressed
		return cmds
	}

	w.AddSystem(&system.AutoInputSystem{&system.InputSystem{}, f})
	w.AddSystem(&system.ControlSystem{})
	w.AddSystem(&system.MovementSystem{})
	r := system.RenderSystem{}
	w.AddSystem(&r)
	ww := WorldWrap{world: &w, RenderSystem: &r}

	player := Player{
		ecs.NewBasic(),
		component.TransformComponent{
			Position: gfx.V(100, 100),
		},
		component.PlayerControlComponent{
			Mapper: map[string]string{
				"jump_1": "jump",
			},
			KeyStates: map[string]component.KeyState{},
		},
		component.RigidBodyComponent{},
	}

	player2 := Player{
		ecs.NewBasic(),
		component.TransformComponent{
			Position: gfx.V(200, 100),
		},
		component.PlayerControlComponent{
			Mapper: map[string]string{
				// "jump_2": "jump",
			},
			KeyStates: map[string]component.KeyState{},
		},
		component.RigidBodyComponent{},
	}

	// Add our entity to the appropriate systems
	for _, sys := range w.Systems() {
		switch typedSystem := sys.(type) {
		case *system.AutoInputSystem:
			typedSystem.Add(&player.BasicEntity, &player.PlayerControlComponent)
			typedSystem.Add(&player2.BasicEntity, &player2.PlayerControlComponent)
		case *system.ControlSystem:
			typedSystem.Add(&player.BasicEntity, &player.PlayerControlComponent, &player.RigidBodyComponent)
			typedSystem.Add(&player2.BasicEntity, &player2.PlayerControlComponent, &player2.RigidBodyComponent)
		case *system.MovementSystem:
			typedSystem.Add(&player.BasicEntity, &player.TransformComponent, &player.RigidBodyComponent)
			typedSystem.Add(&player2.BasicEntity, &player2.TransformComponent, &player2.RigidBodyComponent)
		case *system.RenderSystem:
			typedSystem.Add(&player.BasicEntity, &player.TransformComponent)
			typedSystem.Add(&player2.BasicEntity, &player2.TransformComponent)
		}
	}

	if err := ebiten.Run(ww.Update, 300, 300, 1, "ECS test"); err != nil {
		log.Fatal("exited")
	}
	fmt.Println(player)
	fmt.Println(player2)
}

type WorldWrap struct {
	i     int
	world *ecs.World
	*system.RenderSystem
}

var dt = float32(0.10)

func (ww *WorldWrap) Update(screen *ebiten.Image) error {
	ww.i++
	ww.world.Update(dt)

	ww.RenderSystem.Render(screen)
	return nil
}
