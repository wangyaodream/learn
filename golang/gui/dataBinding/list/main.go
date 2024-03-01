package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("List Data")

    data := binding.BindStringList(
        &[]string{"Item 1", "Item 2", "Item 3"},
    )
    
    lab := widget.NewLabel("Hello")

    list := widget.NewListWithData(
        data,
        func() fyne.CanvasObject {
            return widget.NewLabel("template")
        },
        func(i binding.DataItem, o fyne.CanvasObject) {
            o.(*widget.Label).Bind(i.(binding.String))
        })
    add := widget.NewButton("Append", func() {
        val := fmt.Sprintf("Item %d", data.Length() + 1)
        data.Append(val)
    })
    myWindow.SetContent(container.NewBorder(nil, nil, nil,add, list, lab))
    myWindow.ShowAndRun()
}
