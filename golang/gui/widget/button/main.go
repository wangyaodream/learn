package main

import (
    "log"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/theme"
)

func main() {
    a := app.New()
    w := a.NewWindow("Button Widget")

    content := widget.NewButton("click me", func() {
        log.Println("tapped")
    })

    content2 := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
        log.Println("tapped home")
    })


    w.SetContent(content)
    w.SetContent(content2)
    w.ShowAndRun()
}
