package repository

import (
	"time"
)

// repository type for test
type TestRepository struct {
}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (repo *TestRepository) Migrate() error {
	return nil
}

func (repo *TestRepository) InsertHolding(h Holding) (*Holding, error) {
	return &h, nil
}

func (repo *TestRepository) AllHoldings() ([]Holding, error) {
	var all []Holding
	return all, nil
}

func (repo *TestRepository) GetHoldingByID(id int) (*Holding, error) {
	h := Holding{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}
	return &h, nil
}

func (repo *TestRepository) UpdateHolding(id int64, updated Holding) error {
	return nil
}

func (repo *TestRepository) DeleteHolding(id int64) error {
	return nil
}
