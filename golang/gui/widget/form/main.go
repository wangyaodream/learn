package main

import (
    "log"

    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Form Widget")

    entry := widget.NewEntry()
    textArea := widget.NewMultiLineEntry()

    form := &widget.Form{
        Items: []*widget.FormItem{
            {Text: "Entry", Widget: entry},
            {Text: "Text", Widget: textArea}},
        OnSubmit: func() {
            log.Println("Form submitted:", entry.Text)
            log.Println("multiline", textArea.Text)
            // 当按下submit按钮时会触发退出
            // myWindow.Close()
        },
    }

    // form.Append("Text", textArea)

    myWindow.SetContent(form)
    myWindow.ShowAndRun()
}
