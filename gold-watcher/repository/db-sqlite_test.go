package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepo_Migrate(t *testing.T) {
	err := testRepo.Migrate()
	if err != nil {
		t.Error("migrate failed:", err)
	}
}

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

func TestSQLiteRepo_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed:", err)
	}

	if len(h) != 1 {
		t.Error("wrong number of rows returned; expected 1 but got", len(h))
	}
}

func TestSQLiteRepo_GetHondingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("get by id failed:", err)
	}

	if h.PurchasePrice != 1000 {
		t.Error("wrong purchase price returned; expected 1000 but got", h.PurchasePrice)
	}

	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("get one returned value for non-existent id")
	}
}

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

	h, err = testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	if h.PurchasePrice != 1001 {
		t.Error("wrong purchase price returned; expected 1001 but got", h.PurchasePrice)
	}
}

func TestSQLiteRepo_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding:", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete non-existent record")
	}
}
