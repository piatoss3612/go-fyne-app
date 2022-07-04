package repository

import (
	"testing"
	"time"
)

// 1. test Migrate method
func TestSQLiteRepo_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

// 2. test InsertHolding method
func TestSQLiteRepo_InsertHolding(t *testing.T) {
	h := Holding{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	res, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if res.ID <= 0 {
		t.Error("invalid id sent back:", res.ID)
	}
}

// 3. test AllHoldings method
func TestSQLiteRepo_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed:", err)
	}

	// the number of records should be 1
	if len(h) != 1 {
		t.Error("wrong number of rows returned; expected 1 but got", len(h))
	}
}

// 4. test GetHoldingByID method
func TestSQLiteRepo_GetHondingByID(t *testing.T) {
	// test with valid id
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	// price in retrieved record should be 1000
	if h.PurchasePrice != 1000 {
		t.Error("wrong purchase price returned; expected 1000 but got", h.PurchasePrice)
	}

	// test with invalid id
	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("get one returned value for non-existent id")
	}
}

// 5. test UpdateHolding method
func TestSQLiteRepo_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	h.PurchasePrice = 1001

	err = testRepo.UpdateHolding(1, *h)
	if err != nil {
		t.Error("update failed:", err)
	}

	// retrieve updated record
	h, err = testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	// price in updated record should be 1001
	if h.PurchasePrice != 1001 {
		t.Error("wrong purchase price returned; expected 1001 but got", h.PurchasePrice)
	}
}

// 6. test DeleteHolding method
func TestSQLiteRepo_DeleteHolding(t *testing.T) {
	// test with valid id
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding:", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	// test with invalid id
	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete non-existent record")
	}
}
