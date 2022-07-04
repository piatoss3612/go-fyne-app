package repository

import (
	"errors"
	"time"
)

// custom database errors
var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

// database interface
type Repository interface {
	Migrate() error
	InsertHolding(h Holding) (*Holding, error)
	AllHoldings() ([]Holding, error)
	GetHoldingByID(id int) (*Holding, error)
	UpdateHolding(id int64, updated Holding) error
	DeleteHolding(id int64) error
}

// database record definition
type Holding struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
