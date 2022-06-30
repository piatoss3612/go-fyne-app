package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
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
	// create main UIs
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	app.EditWidget = edit
	app.PreviewWidget = preview

	// event listener: listen to changes of text on edit widget
	// event handler: parsing text on edit widget to markdown and display on preview widget
	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(w fyne.Window) {
	// create menu items
	openMenuItem := fyne.NewMenuItem("Open...", app.openFunc(w))
	saveMenuItem := fyne.NewMenuItem("Save", app.saveFunc(w))

	// disactivate saveMenuItem
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("Save as", app.saveAsFunc(w))

	// aggregate menu items as File menu
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// create main menu
	menu := fyne.NewMainMenu(fileMenu)

	// set main menu to window
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

			// check entered file name's extension
			if !strings.HasSuffix(strings.ToLower(uc.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Please name your file with .md extension!", w)
			}

			// save file
			uc.Write([]byte(app.EditWidget.Text))

			// cache saved file path
			app.CurrentFile = uc.URI()

			defer uc.Close()

			// set current window title as Markdown - saved file name
			w.SetTitle("Markdown - " + uc.URI().Name())

			// activate save menu
			app.SaveMenuItem.Disabled = false
		}, w)

		// set default file name and filter
		saveDialog.SetFileName("untitled.md")
		saveDialog.SetFilter(filter)

		// show save as dialog
		saveDialog.Show()
	}
}

var filter = &storage.ExtensionFileFilter{
	Extensions: []string{".md", ".MD"},
}

func (app *config) openFunc(w fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			if uc == nil {
				return
			}

			defer uc.Close()

			// import markdown file
			data, err := ioutil.ReadAll(uc)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// set the content of EditWidget to imported markdown
			app.EditWidget.SetText(string(data))

			// cache imported file path
			app.CurrentFile = uc.URI()

			w.SetTitle("Markdown - " + uc.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, w)

		// set file extensions filter
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (app *config) saveFunc(w fyne.Window) func() {
	return func() {
		// check cached file path
		if app.CurrentFile != nil {
			// set up to write file to the URI reference
			write, err := storage.Writer(app.CurrentFile)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}

			// save file
			write.Write([]byte(app.EditWidget.Text))
			defer write.Close()
		}
	}
}
