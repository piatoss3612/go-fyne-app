package main

import "testing"

// test getting records from DB
func TestConfig_getHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from database:", err)
	}

	// test repository should have 2 dummy records
	if len(all) != 2 {
		t.Error("wrong number of rows returned")
	}
}

// test conversion of records to slice
func TestConfig_getHoldingSlice(t *testing.T) {
	slice := testApp.getHoldingSlice()

	// slice length should be 3 (2 dummy records + 1 header)
	if len(slice) != 3 {
		t.Error("wrong number of rows returnd; expected 3 but got", len(slice))
	}
}
