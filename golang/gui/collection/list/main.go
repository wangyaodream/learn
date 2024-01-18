package main

import (
    "log"

    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("List widget")

    list := widget.NewList(
        func() int {
            // return len(data)
            return 3
        },
        func() fyne.CanvasObject {
            return widget.NewButton("template", func(){
                log.Println("Button clicked")
            })
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {
            o.(*widget.Button).SetText("hello")
        })
    
    myWindow.Resize(fyne.NewSize(100,300))
    myWindow.SetContent(list)
    myWindow.ShowAndRun()
}
