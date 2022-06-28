package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	// initialize application instance
	a := app.New()

	// create master window
	w := a.NewWindow("Hello, World!")

	// make UIs
	output, entry, btn := myApp.makeUI()

	// set the content of window
	w.SetContent(container.NewVBox(output, entry, btn))

	//
	w.Resize(fyne.NewSize(500, 500))

	// show the window and run application with infinite loop
	w.ShowAndRun()

}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Enter", func() {
		app.output.SetText(entry.Text)
	})
	btn.Importance = widget.HighImportance

	app.output = output
	return output, entry, btn
}
