package main

import (
    "image/color"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/theme"
)

func main() {
    a := app.New()
    w := a.NewWindow("Center Layout")

    img := canvas.NewImageFromResource(theme.FyneLogo())
    img.FillMode = canvas.ImageFillOriginal

    text := canvas.NewText("Overlay", color.Black)
    text2 := canvas.NewText("Other text", color.Black)
    content := container.New(layout.NewCenterLayout(), img, text, text2)

    w.SetContent(content)
    w.ShowAndRun()
}

