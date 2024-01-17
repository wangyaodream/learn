package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    w := a.NewWindow("Form Layout")

    label1 := widget.NewLabel("lab1")
    value1 := widget.NewLabel("value1")
    label2 := widget.NewLabel("lab2")
    value2 := widget.NewLabel("Something")
    lab3 := widget.NewLabel("lab3")
    val3 := widget.NewLabel("valuvaluesljsdljflsjdfk")
    grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2, lab3, val3)
    w.SetContent(grid)

    w.ShowAndRun()
}
