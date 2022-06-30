package main

import "testing"

// test if the number of toolbar items is 4
func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}
