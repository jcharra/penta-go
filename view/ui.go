package view

import (
	"image/color"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type pentagoScene struct{}

// Type uniquely defines your game type
func (*pentagoScene) Type() string {
	return "pentago"
}

// Preload is called before loading any assets from the disk, to allow you to register / queue them
func (*pentagoScene) Preload() {
	err := engo.Files.Load("UbuntuMono-R.ttf")
	if err != nil {
		panic(err)
	}
}

// Setup is called before the main loop starts. It allows you to add entities and systems to your Scene.
func (*pentagoScene) Setup(world engo.Updater) {
	common.SetBackground(color.Black)

	/* CONTINUE HERE
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.MouseSystem{})
	world.AddSystem(&BoardSystem{})
	*/
}

func RunUI() {
	opts := engo.RunOptions{
		Title:  "Pentago",
		Width:  1000,
		Height: 800,
	}

	engo.Run(opts, &pentagoScene{})
}
