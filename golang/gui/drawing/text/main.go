package main

import (
    "image/color"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
)

func main() {
    a := app.New()
    w := a.NewWindow("Text")

    text := canvas.NewText("Text Object", color.Black)
    text.Alignment = fyne.TextAlignTrailing
    text.TextStyle = fyne.TextStyle{Italic: true}
    w.SetContent(text)

    w.ShowAndRun()
}
