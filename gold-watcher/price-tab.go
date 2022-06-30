package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	// get price chart
	chart := app.getChart()

	// create price chart container
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	// get price chart image from apiURL and save as gold.png
	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		// use bundled image to issue error
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.NewSize(770, 410))
	//img.FillMode = canvas.ImageFillOriginal // doesn't display image
	img.FillMode = canvas.ImageFillContain

	return img
}

func (app *Config) downloadFile(URL, filename string) error {
	// get the response bytes from calling the URL
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	// check response status code
	if response.StatusCode != 200 {
		return errors.New("received wrong response code while downloading image")
	}

	// read bytes from response body
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// decode bytes to image
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	// create named file
	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}

	// encode decoded image to named file
	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
