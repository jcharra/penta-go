package view

import (
	"image/color"
	"github.com/jcharra/penta-go/core"
	"engo.io/ecs"
	"engo.io/engo/common"
	"engo.io/engo"
)

type Checker struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	color color.Color
}

type Field struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
}

type BoardSystem struct {
	world      *ecs.World
	entities   []Checker
	fields     [6][6]Field
	checker    [6][6]Checker
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
	separator := fieldSize/20.0

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			bs.fields[i][j] = Field{}
			bs.fields[i][j].RenderComponent = common.RenderComponent{
				Drawable: common.Rectangle{BorderWidth: 1, BorderColor: color.White},
				Scale:    engo.Point{1.0, 1.0},
				Color: color.Black,
			}
			bs.fields[i][j].SpaceComponent = common.SpaceComponent{
				Position: engo.Point{fieldSize * float32(i) + separator * float32(i/3), fieldSize * float32(j) + separator * float32(j/3)},
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
	var renderSystem *common.RenderSystem
	for _, system := range bs.world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			renderSystem = sys
		}
	}

	for i, row := range bs.fields {
		for j, field := range row {
			if field.MouseComponent.Clicked {
				if bs.boardModel.Fields[i][j] != 0 {
					continue;
				}
				bs.boardModel = bs.boardModel.SetAt(i, j)
			}

			if field.MouseComponent.RightClicked {
				quad := 0
				if i > 2 {
					quad += 2
				}
				if j > 2 {
					quad += 1
				}
				bs.boardModel = bs.boardModel.Rotate(quad, 0)
			}

			if bs.boardModel.Fields[i][j] == 0 && bs.checker[i][j].color != nil {
				bs.checker[i][j].color = nil
				renderSystem.Remove(bs.checker[i][j].BasicEntity)
			} else if bs.boardModel.Fields[i][j] != 0 && bs.checker[i][j].color != mappedColor(bs.boardModel.Fields[i][j]) {
				if bs.checker[i][j].color != nil {
					renderSystem.Remove(bs.checker[i][j].BasicEntity)
				}
				bs.checker[i][j] = createChecker(&field.SpaceComponent, mappedColor(bs.boardModel.Fields[i][j]))
				renderSystem.Add(&bs.checker[i][j].BasicEntity, &bs.checker[i][j].RenderComponent, &bs.checker[i][j].SpaceComponent)
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

func (b *BoardSystem) Add(c *Checker) {
	b.entities = append(b.entities, *c)
}

func (b *BoardSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range b.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		b.entities = append(b.entities[:delete], b.entities[delete + 1:]...)
	}
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
	checker.color = col
	return checker
}