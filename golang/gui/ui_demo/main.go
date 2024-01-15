package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("This is demo")

    hello := widget.NewLabel("Hello Fyne")
    w.SetContent(container.NewVBox(
        hello, 
        widget.NewButton("HI!", func() {
            hello.SetText("Welcome!")
        }),
    ))

    w.ShowAndRun()
}
