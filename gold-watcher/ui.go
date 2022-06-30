package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

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

	// get price tab content
	priceTabContent := app.pricesTab()

	// get application tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("Holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceContent, toolbar, tabs)
	app.MainWindow.SetContent(finalContent)

	// refresh price contents every 10 seconds
	go app.refreshPriceContentAtInterval(10)
}

func (app *Config) refreshPriceContent() {
	app.InfoLog.Println("refreshing prices")

	// renew and refresh container
	open, current, change := app.getPriceText()
	app.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	app.PriceContainer.Refresh()

	chart := app.getChart()
	app.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	app.PriceChartContainer.Refresh()
}

func (app *Config) refreshPriceContentAtInterval(n int64) {
	for range time.Tick(time.Second * time.Duration(n)) {
		app.refreshPriceContent()
	}
}
