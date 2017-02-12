package view

import (
	"image/color"
	"github.com/jcharra/penta-go/core"
	"engo.io/ecs"
	"engo.io/engo/common"
	"engo.io/engo"
	"fmt"
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
				Drawable: common.Rectangle{BorderWidth: 1, BorderColor: color.White},
				Scale:    engo.Point{1.0, 1.0},
				Color: color.Black,
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

// Update is run every frame, with `dt` being the time
// in seconds since the last frame
func (bs *BoardSystem) Update(dt float32) {
	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			synchronizeBoardView(bs, sys)
		}
	}

	for i, row := range bs.fields {
		for j, field := range row {
			if field.MouseComponent.Released {
				bs.boardModel = bs.boardModel.SetAt(i, j).Rotate(core.UPPERLEFT, core.CLOCKWISE)
			}
		}
	}
}

func synchronizeBoardView(bs *BoardSystem, sys *common.RenderSystem) {
	for i, row := range bs.boardModel.Fields {
		for j, val := range row {
			// fc is the corresponding field in the view model
			fc := bs.fields[i][j]

			actualColor := mappedColor(val)
			if fc.checker.color == nil {
				if actualColor != nil {
					c := createChecker(&fc.SpaceComponent, actualColor)
					sys.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
				}
			} else {
				if actualColor == nil {
					sys.Remove(fc.checker.BasicEntity)
				} else if actualColor != fc.checker.color {
					sys.Remove(fc.checker.BasicEntity)
					c := createChecker(&fc.SpaceComponent, actualColor)
					sys.Add(&c.BasicEntity, &c.RenderComponent, &c.SpaceComponent)
				}
			}
		}
	}
}

func mappedColor(col int) color.Color {
	if col == core.WHITE {
		return color.RGBA{R: 255, G:0, B:0, A:255}
	} else if col == core.BLACK {
		return color.RGBA{R: 0, G:255, B:0, A:255}
	}
	return nil
}

// Remove is called whenever an Entity is removed from the scene, and thus from this system
func (*BoardSystem) Remove(e ecs.BasicEntity) {
	fmt.Println("Should remove entity", e)
}

func createChecker(sc *common.SpaceComponent, col color.Color) (checker Checker) {
	checker = Checker{BasicEntity: ecs.NewBasic()}
	checker.RenderComponent = common.RenderComponent{
		Drawable: common.Circle{BorderWidth: 20, BorderColor: col},
		Scale:    engo.Point{1.0, 1.0},
		Color: color.Black,
	}

	size := sc.Width * 0.5
	offset := (sc.Width - size) / 2
	checker.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{sc.Position.X + offset, sc.Position.Y + offset},
		Width: size,
		Height: size,
	}
	return checker
}
