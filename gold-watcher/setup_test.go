package main

import (
	"bytes"
	"gold-watcher/repository"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	// create dummy fyne application for testing
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.DB = repository.NewTestRepository()
	testApp.HTTPClient = client

	// run application test and pass exit code to os.Exit
	os.Exit(m.Run())
}

// actual data from https://data-asg.goldprice.org/dbXRates/USD as dummy data
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

// take http request then return dummy http response
type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// create dummy http client that only returns jsonToReturn dummy data as body of response
var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
