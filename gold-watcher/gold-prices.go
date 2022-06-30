package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var currency = "USD"

type Gold struct {
	Prices []Price `json:"items"`
	Client *http.Client
}

type Price struct {
	Currency      string    `json:"currency"`
	Price         float64   `json:"xauPrice"`
	Change        float64   `json:"chgXau"`
	PreviousClose float64   `json:"xauClose"`
	Time          time.Time `json:"-"`
}

func (g *Gold) GetPrices() (*Price, error) {
	if g.Client == nil {
		g.Client = &http.Client{}
	}

	client := g.Client
	url := fmt.Sprintf("https://data-asg.goldprice.org/dbXRates/%s", currency)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error while creating request to goldprice.org")
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error while connecting goldprice.org")
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while reading response from goldprice.org")
		return nil, err
	}

	gold := Gold{}

	var previous, current, change float64
	err = json.Unmarshal(body, &gold)
	if err != nil {
		log.Println("error while decoding response from goldprice.org")
		return nil, err
	}

	previous, current, change = gold.Prices[0].PreviousClose, gold.Prices[0].Price, gold.Prices[0].Change

	var currentInfo = Price{
		Currency:      currency,
		Price:         current,
		Change:        change,
		PreviousClose: previous,
		Time:          time.Now(),
	}

	return &currentInfo, nil
}
