package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    _ "fyne.io/fyne/v2/layout"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Border layout")
    myWindow.Resize(fyne.NewSize(250, 200))

    top := canvas.NewText("top bar", color.RGBA{100,100,100,255})
    left := canvas.NewText("left", color.RGBA{200, 200, 100, 255})
    middle := canvas.NewText("content", color.RGBA{155, 15, 100, 255})
    content := container.NewBorder(top, nil, left, nil, middle)
    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
