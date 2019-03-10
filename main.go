package main

import (
	"engo.io/ecs"
	"github.com/kyeett/compo/component"
	"github.com/kyeett/compo/system"
)

type Player struct {
	ecs.BasicEntity
	component.PlayerControlComponent
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
	// w.AddSystem(&system.ControlSystem{})

	player := Player{ecs.NewBasic(), component.PlayerControlComponent{
		Mapper: map[string]string{
			"jump_1": "jump",
		},
		KeyStates: map[string]component.KeyState{},
	}}

	player2 := Player{ecs.NewBasic(), component.PlayerControlComponent{
		Mapper: map[string]string{
			"jump_2": "jump",
		},
		KeyStates: map[string]component.KeyState{},
	}}

	// Add our entity to the appropriate systems
	for _, sys := range w.Systems() {
		switch typedSystem := sys.(type) {
		case *system.AutoInputSystem:
			typedSystem.Add(&player.BasicEntity, &player.PlayerControlComponent)
			typedSystem.Add(&player2.BasicEntity, &player2.PlayerControlComponent)
		}
	}

	dt := float32(1.0)
	for i := 0; i < 10; i++ {
		w.Update(dt)
	}
}

// type InputManager struct {
// }

// func NewInputManager() *InputManager {
// 	return &InputManager{
// 		Touches: make(map[int]Point),
// 		axes:    make(map[string]Axis),
// 		buttons: make(map[string]Button),
// 		keys:    NewKeyManager(),
// 	}
// }
