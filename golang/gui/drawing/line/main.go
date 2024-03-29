package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
)

func main() {
    myApp := app.New()
    w := myApp.NewWindow("Line")

    line := canvas.NewLine(color.Black)
    line.StrokeWidth = 5
    w.SetContent(line)

    w.Resize(fyne.NewSize(100, 200))
    w.ShowAndRun()
}
