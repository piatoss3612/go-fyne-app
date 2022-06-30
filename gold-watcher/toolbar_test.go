package main

import "testing"

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of items in toolbar")
	}
}
