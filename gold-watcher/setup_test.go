package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// data from https://data-asg.goldprice.org/dbXRates/USD
var jsonToReturn = `
{
	"ts": 1656490029704,
	"tsj": 1656490025944,
	"date": "Jun 29th 2022, 04:07:05 am NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 1817.075,
		"xagPrice": 20.7963,
		"chgXau": -1.43,
		"chgXag": -0.0267,
		"pcXau": -0.0786,
		"pcXag": -0.1282,
		"xauClose": 1818.505,
		"xagClose": 20.823
	  }
	]
  }
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
