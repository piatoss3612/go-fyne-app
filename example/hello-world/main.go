package main

import (
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// initialize application instance
	a := app.New()

	// create master window
	w := a.NewWindow("Hello, World!")

	// set the content of window
	w.SetContent(widget.NewLabel("Hello, world!"))

	// show the window and run application with infinite loop
	w.ShowAndRun()

	// excute when application is closed
	tidy()
}

func tidy() {
	log.Println("would tidy up")
}
