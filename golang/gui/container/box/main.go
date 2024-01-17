package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
    myApp := app.New()
    w := myApp.NewWindow("Box layout")

    text1 := canvas.NewText("Hello", color.RGBA{255,155,0,255})
    text2 := canvas.NewText("There", color.RGBA{155,200,100,255})
    text3 := canvas.NewText("(right)", color.RGBA{10,20,200,255})
    content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

    text4 := canvas.NewText("centered", color.RGBA{100,100,100,255})
    centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4)
    w.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
    w.ShowAndRun()


}
