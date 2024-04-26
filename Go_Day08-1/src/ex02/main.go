package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

const (
	WIDTH  int = 300
	HEIGHT int = 200
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(app.Size(unit.Dp(WIDTH), unit.Dp(HEIGHT)), app.Title("School 21"))
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			title := material.H6(theme, "<3")
			title.Alignment = text.Middle
			title.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}
