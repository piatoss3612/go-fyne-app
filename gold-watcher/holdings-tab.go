package main

import (
	"gold-watcher/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/widget"
)

func (app *Config) holdingsTab() *fyne.Container {
	return nil
}

func (app *Config) getHoldingsTable() *widget.Table {
	return nil
}

func (app *Config) getHoldingSlice() [][]any {
	var slice [][]any
	return slice
}

func (app *Config) currentHoldings() ([]repository.Holding, error) {
	holdings, err := app.DB.AllHoldings()
	if err != nil {
		app.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
