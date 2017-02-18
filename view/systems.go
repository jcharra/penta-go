package view

import (
	"image/color"
	"github.com/jcharra/penta-go/core"
	"engo.io/ecs"
	"engo.io/engo/common"
	"engo.io/engo"
	"github.com/jcharra/penta-go/ai"
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

type StatusLabel struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type BoardSystem struct {
	world      *ecs.World
	entities   []Checker
	fields     [6][6]Field
	checker    [6][6]Checker
	gameState  int
	stateLabel StatusLabel
	boardModel core.Board
}

const fieldSize float32 = 100.0

// enumeration of ui game states
const (
	waitForChecker = iota
	waitForRotation = iota
	computerThinking = iota
	computerSettingChecker = iota
	computerRotating = iota
	gameDrawn = iota
	gameWonPlayer = iota
	gameWonAI = iota
	evaluatePosition = iota
)

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
	separator := fieldSize / 20.0

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			bs.fields[i][j] = Field{}
			bs.fields[i][j].RenderComponent = common.RenderComponent{
				Drawable: common.Rectangle{BorderWidth: 1, BorderColor: color.White},
				Scale:    engo.Point{1.0, 1.0},
				Color: color.Black,
			}
			bs.fields[i][j].SpaceComponent = common.SpaceComponent{
				Position: engo.Point{fieldSize * float32(i) + separator * float32(i / 3), fieldSize * float32(j) + separator * float32(j / 3)},
				Width: fieldSize,
				Height: fieldSize,
			}

			mouseSys.Add(&bs.fields[i][j].BasicEntity, &bs.fields[i][j].MouseComponent, &bs.fields[i][j].SpaceComponent, nil)
			renderSys.Add(&bs.fields[i][j].BasicEntity, &bs.fields[i][j].RenderComponent, &bs.fields[i][j].SpaceComponent)
		}
	}

	fntWhite := &common.Font{
		URL:  "UbuntuMono-R.ttf",
		FG:   color.White,
		Size: 64,
	}
	err := fntWhite.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	bs.stateLabel = StatusLabel{BasicEntity: ecs.NewBasic()}
	bs.stateLabel.RenderComponent.Drawable = common.Text{
		Font: fntWhite,
		Text: "Welcome to Pentago",
	}
	bs.stateLabel.SetShader(common.HUDShader)

	renderSys.Add(&bs.stateLabel.BasicEntity,
		&bs.stateLabel.RenderComponent,
		&common.SpaceComponent{Position: engo.Point{0, fieldSize * 6.0 + 10}})

	bs.gameState = waitForChecker
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

	// Synchronize board model and board UI
	for i, row := range bs.fields {
		for j, field := range row {
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

	// Handle game state changes
	if bs.gameState == waitForChecker || bs.gameState == waitForRotation {
		for i, row := range bs.fields {
			for j, field := range row {
				if field.MouseComponent.Clicked {

					if bs.gameState == waitForChecker {
						if bs.boardModel.Fields[i][j] != 0 {
							// Attempt to place checker on occupied field => ignore
							continue;
						}
						bs.boardModel = bs.boardModel.SetAt(i, j)
						bs.gameState = waitForRotation
					} else {
						quad := quadrantForIndexes(i, j)
						bs.boardModel = bs.boardModel.Rotate(quad, 0)
						bs.gameState = evaluatePosition
					}
				} else if field.MouseComponent.RightClicked && bs.gameState == waitForRotation {
					quad := quadrantForIndexes(i, j)
					bs.boardModel = bs.boardModel.Rotate(quad, 1)
					bs.gameState = evaluatePosition
				}
			}
		}
	} else if bs.gameState == computerThinking {
		bestMove := ai.FindBestMove(bs.boardModel, 2, 2)

		bs.gameState = computerSettingChecker
		bs.boardModel = bs.boardModel.SetAt(bestMove.Move.Row, bestMove.Move.Col)

		bs.gameState = computerRotating
		bs.boardModel = bs.boardModel.Rotate(bestMove.Move.Quadrant, bestMove.Move.Direction)

		bs.gameState = evaluatePosition
	} else if bs.gameState == evaluatePosition {
		winner := bs.boardModel.Winner()
		if winner == core.WHITE {
			bs.gameState = gameWonPlayer
		} else if winner == core.BLACK {
			bs.gameState = gameWonAI
		} else if winner == core.DRAW {
			bs.gameState = gameDrawn
		} else {
			if bs.boardModel.Turn == core.WHITE {
				bs.gameState = waitForChecker
			} else {
				bs.gameState = computerThinking
			}
		}
	}

	fntWhite := &common.Font{
		URL:  "UbuntuMono-R.ttf",
		FG:   color.White,
		Size: 16,
	}
	err := fntWhite.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	bs.stateLabel.RenderComponent.Drawable = common.Text{
		Font: fntWhite,
		Text: textForGameState(bs.gameState),
	}
}

func quadrantForIndexes(i, j int) int {
	quad := 0
	if i > 2 {
		quad += 2
	}
	if j > 2 {
		quad += 1
	}
	return quad
}

func textForGameState(state int) string {
	switch (state) {
	case waitForChecker:
		return "Place a checker"
	case waitForRotation:
		return "Rotate a quadrant (left-click rotates clockwise, right-click counterclockwise)"
	case computerThinking:
		return "Planning my next move ..."
	case computerSettingChecker:
		return ""
	case computerRotating:
		return "Rotating"
	case gameDrawn:
		return "The game has ended. It's a draw."
	case gameWonPlayer:
		return "You win - congratulations!"
	case gameWonAI:
		return "I win - better luck next time!"
	default:
		return ""
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