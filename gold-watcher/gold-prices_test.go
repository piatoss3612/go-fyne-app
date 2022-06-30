package main

import (
	"testing"
)

// test if expected result of GetPrices function is equal to dummy data
func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 1817.075 {
		t.Error("wrong price returned:", p.Price)
	}
}
