package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

// test if the number of toolbar items is 4
func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}

// test addHoldings dialog
func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.addHoldingsDialog()

	test.Type(testApp.AddHoldingsPurchaseAmountEntry, "1")
	test.Type(testApp.AddHoldingsPurchasePriceEntry, "1000")
	test.Type(testApp.AddHoldingsPurchaseDateEntry, "2022-07-04")

	if testApp.AddHoldingsPurchaseDateEntry.Text != "2022-07-04" {
		t.Error("date not correct")
	}

	if testApp.AddHoldingsPurchaseAmountEntry.Text != "1" {
		t.Error("amount not correct")
	}

	if testApp.AddHoldingsPurchasePriceEntry.Text != "1000" {
		t.Error("price not correct")
	}
}
