package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Raster")

	raster := canvas.NewRasterWithPixels(
		func(_, _, w, h int) color.Color {
			return color.RGBA{uint8(rand.Intn(50)),
				uint8(rand.Intn(50)),
				uint8(rand.Intn(50)), 0xF1}
		})
	// raster := canvas.NewRasterFromImage()
	w.SetContent(raster)
	w.Resize(fyne.NewSize(120, 100))
	w.ShowAndRun()
}
