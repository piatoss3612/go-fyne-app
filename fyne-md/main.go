package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// create a fyne app
	a := app.New()

	// create a window
	w := a.NewWindow("Markdown")

	// get the user interface
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(w)

	// set the content of the window
	w.SetContent(container.NewHSplit(edit, preview))

	// show window and run app
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	app.EditWidget = edit
	app.PreviewWidget = preview

	// add event listener
	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(w fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})

	// disactivate saveMenuItem
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save as", app.saveAsFunc(w))

	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	w.SetMainMenu(menu)
}

func (app *config) saveAsFunc(w fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if uc == nil {
				// user cancelled
				return
			}

			// save file
			uc.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = uc.URI()

			defer uc.Close()

			w.SetTitle("Markdown - " + uc.URI().Name())

			app.SaveMenuItem.Disabled = false
		}, w)

		saveDialog.Show()
	}
}
