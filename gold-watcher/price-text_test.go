package main

import "testing"

// test if expected result of getPriceText function is equal to formatted dummy data
func TestApp_getPriceText(t *testing.T) {
	open, current, change := testApp.getPriceText()
	if open.Text != "Open: $1818.5050 USD" {
		t.Error("wrong price returned", open.Text)
	}

	if current.Text != "Current: $1817.0750 USD" {
		t.Error("wrong price returned", current.Text)
	}

	if change.Text != "Change: $-1.4300 USD" {
		t.Error("wrong price returned", change.Text)
	}
}
