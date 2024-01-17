package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    _ "fyne.io/fyne/v2/theme"
)

func main() {
    myApp := app.New()
    w := myApp.NewWindow("Image")

    // image := canvas.NewImageFromResource(theme.FyneLogo())
    image := canvas.NewImageFromFile("/Users/wangyao/Downloads/target.jpg")
    image.FillMode = canvas.ImageFillOriginal
    w.SetContent(image)

    w.ShowAndRun()
}
