package main

import (
	"fmt"
	"gold-watcher/repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingsTable() *widget.Table {
	return nil
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
		currentRow = append(currentRow, fmt.Sprintf("$%2f", float32(h.PurchasePrice/100)))
		currentRow = append(currentRow, h.PurchaseDate.Format(time.RFC822))
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
