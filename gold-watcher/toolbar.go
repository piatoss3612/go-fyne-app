package main

import (
	"gold-watcher/repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) getToolbar() *widget.Toolbar {
	// create toolbar with 4 items
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			app.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			app.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			w := app.showPreferences()
			w.Resize(fyne.NewSize(300, 200))
			w.Show()
		}),
	)

	return toolbar
}

// entry input validator
var (
	isIntValidator = func(s string) error {
		_, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		return nil
	}
	isFloatValidator = func(s string) error {
		_, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		return nil
	}
	dateValidator = func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}
		return nil
	}
)

func (app *Config) addHoldingsDialog() dialog.Dialog {
	// create entries and set references
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	app.AddHoldingsPurchaseAmountEntry = addAmountEntry
	app.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	app.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	// add validation
	addAmountEntry.Validator = isIntValidator
	purchasePriceEntry.Validator = isFloatValidator
	purchaseDateEntry.Validator = dateValidator

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	// create a dialog
	addForm := dialog.NewForm(
		"Add Gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount in toz", Widget: addAmountEntry},
			{Text: "Purchase Price", Widget: purchasePriceEntry},
			{Text: "Purchase Date", Widget: purchaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				purchaseDate, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
				purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)

				_, err := app.DB.InsertHolding(repository.Holding{
					Amount:        amount,
					PurchaseDate:  purchaseDate,
					PurchasePrice: int(purchasePrice * 100),
				})

				if err != nil {
					app.ErrorLog.Println(err)
				}

				app.refreshHoldingsTable()
			}
		},
		app.MainWindow,
	)

	// resize ans show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}

func (app *Config) showPreferences() fyne.Window {
	win := app.App.NewWindow("Preferences")

	label := widget.NewLabel("Preferred Currency")
	cur := widget.NewSelect([]string{"USD", "CAD", "KRW"}, func(value string) {
		currency = value
		app.App.Preferences().SetString("currency", value)
	})
	cur.Selected = currency

	btn := widget.NewButton("Save", func() {
		win.Close()
		app.refreshPriceContent()
	})
	btn.Importance = widget.HighImportance

	win.SetContent(container.NewVBox(label, cur, btn))

	return win
}
