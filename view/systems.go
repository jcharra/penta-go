package view

import (

	"image/color"
	"github.com/jcharra/penta-go/core"
	"github.com/engoengine/engo"
	"github.com/engoengine/engo/common"
)

type Checker struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	color color.Color
}

type FieldComponent struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	checker Checker
}

type BoardSystem struct {
	world      *ecs.World
	fields     [6][6]FieldComponent
	boardModel core.Board
}

// New is the initialisation of the System
func (bs *BoardSystem) New(w *ecs.World) {
	bs.world = w
	bs.boardModel = core.NewBoard()

	var renderSys *common.RenderSystem
	var mouseSys *common.MouseSystem
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			mouseSys = sys

		case *common.RenderSystem:
			renderSys = sys
		}
	}

	// size in pixels
	fieldSize := float32(150.0)

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			bs.fields[i][j] = FieldComponent{}
			bs.fields[i][j].RenderComponent = common.RenderComponent{
				Drawable: common.Rectangle{BorderWidth: 1, BorderColor: color.Black},
				Scale:    engo.Point{1.0, 1.0},
			}
			bs.fields[i][j].SpaceComponent = common.SpaceComponent{
				Position: engo.Point{fieldSize * float32(i), fieldSize * float32(j)},
				Width: fieldSize,
				Height: fieldSize,
			}
			mouseSys.Add(&bs.fields[i][j].BasicEntity, &bs.fields[i][j].MouseComponent, &bs.fields[i][j].SpaceComponent, nil)
			renderSys.Add(&bs.fields[i][j].BasicEntity, &bs.fields[i][j].RenderComponent, &bs.fields[i][j].SpaceComponent)
		}
	}
}

// Update is ran every frame, with `dt` being the time
// in seconds since the last frame
func (bs *BoardSystem) Update(dt float32) {
/*
	// Add to the system
	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&checker.BasicEntity, &checker.RenderComponent, &checker.SpaceComponent)
		}
	}
	*/
}

// Remove is called whenever an Entity is removed from the scene, and thus from this system
func (*BoardSystem) Remove(ecs.BasicEntity) {}

/*
func createChecker(bs *BoardSystem) (checker Checker) {
	checker = Checker{BasicEntity: ecs.NewBasic()}
	checker.RenderComponent = common.RenderComponent{
		Drawable: common.Circle{BorderWidth: 20, BorderColor: color.Black},
		Scale:    engo.Point{1.0, 1.0},
	}
	checker.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{bs.mouseTracker.MouseComponent.MouseX, bs.mouseTracker.MouseComponent.MouseY},
		Width:    100,
		Height:   100,
	}

	return checker
}

*/