package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) getPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g Gold
	var open, current, change *canvas.Text
	g.Client = app.HTTPClient

	// get prices from goldprice.org
	gold, err := g.GetPrices()
	if err != nil {
		// display error with grey colored text
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}

		open = canvas.NewText("Open: Unreachable", grey)
		current = canvas.NewText("Current: Unreachable", grey)
		change = canvas.NewText("Change: Unreachable", grey)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255} // green

		// if current gold price is loswer than previous close price
		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255} // red
		}

		// format texts and set contents
		openTxt := fmt.Sprintf("Open: $%.4f %s", gold.PreviousClose, currency)
		currentTxt := fmt.Sprintf("Current: $%.4f %s", gold.Price, currency)
		changeTxt := fmt.Sprintf("Change: $%.4f %s", gold.Change, currency)

		open = canvas.NewText(openTxt, nil)
		current = canvas.NewText(currentTxt, displayColor)
		change = canvas.NewText(changeTxt, displayColor)
	}

	open.Alignment = fyne.TextAlignLeading    // align left
	current.Alignment = fyne.TextAlignCenter  // align center
	change.Alignment = fyne.TextAlignTrailing // align right

	return open, current, change
}
