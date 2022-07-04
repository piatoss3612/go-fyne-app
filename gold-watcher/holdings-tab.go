package main

import (
	"fmt"
	"gold-watcher/repository"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	app.Holdings = app.getHoldingSlice()
	app.HoldingsTable = app.getHoldingsTable()
	return container.NewBorder(
		nil, nil, nil, nil, container.NewAdaptiveGrid(1, app.HoldingsTable),
	)
}

func (app *Config) getHoldingsTable() *widget.Table {
	// create table
	t := widget.NewTable(
		// return the number of rows and columns
		func() (int, int) {
			return len(app.Holdings), len(app.Holdings[0])
		},

		// return template objects
		func() fyne.CanvasObject {
			ctr := container.NewVBox(widget.NewLabel(""))
			return ctr
		},

		// apply data at specified location on the table
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// last cell of the row; expect header
			if i.Col == (len(app.Holdings[0])-1) && i.Row != 0 {
				// put in a delete button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					// show dialog when button is tabbed
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, err := strconv.ParseInt(app.Holdings[i.Row][0].(string), 10, 64)
							if err != nil {
								app.ErrorLog.Println(err)
								return
							}

							// delete current row
							err = app.DB.DeleteHolding(id)
							if err != nil {
								app.ErrorLog.Println(err)
							}
						}

						// refresh holdings table
						app.refreshHoldingsTable()
					}, app.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// put in textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(app.Holdings[i.Row][i.Col].(string)),
				}
			}
		})

	// set each column width on the table
	colWidths := []float32{45, 180, 180, 180, 180, 80}
	for i := 0; i < len(colWidths); i++ {
		t.SetColumnWidth(i, colWidths[i])
	}

	return t
}

// process records into slice to construct a table
func (app *Config) getHoldingSlice() [][]any {
	var slice [][]any

	holdings, err := app.currentHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// add header
	slice = append(slice, []any{"ID", "Amount", "Price", "Date", "Delete"})

	for _, h := range holdings {
		var currentRow []any

		currentRow = append(currentRow, strconv.FormatInt(h.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", h.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%.2f", float32(h.PurchasePrice)/100))
		currentRow = append(currentRow, h.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

	return slice
}

// get all records from DB and return
func (app *Config) currentHoldings() ([]repository.Holding, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
