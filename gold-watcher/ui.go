package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// get the current price of gold
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange)

	app.PriceContainer = priceContent

	// get toolbar
	toolbar := app.getToolbar()
	app.Toolbar = toolbar

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar)

	app.MainWindow.SetContent(finalContent)
}
