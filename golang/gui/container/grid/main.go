package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Grad Layout")

    text1 := canvas.NewText("1", color.RGBA{123,123,123,255})
    text2 := canvas.NewText("2", color.RGBA{123,123,123,255})
    text3 := canvas.NewText("3", color.RGBA{123,123,123,255})
    text4 := canvas.NewText("4", color.RGBA{123,123,123,255})
    // grid := container.New(layout.NewGridLayout(3), text1, text2, text3, text4)
    grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
            text1, text2, text3, text4)
    myWindow.SetContent(grid)
    myWindow.ShowAndRun()

}
