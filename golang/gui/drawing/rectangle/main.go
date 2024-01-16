package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Rectangle")

    rect := canvas.NewRectangle(color.Black)
    myWindow.SetContent(rect)

    myWindow.Resize(fyne.NewSize(150, 100))
    myWindow.ShowAndRun()
}
