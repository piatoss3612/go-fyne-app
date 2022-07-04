package repository

import (
	"database/sql"
	"errors"
	"time"
)

type SQLiteRepo struct {
	Conn *sql.DB
}

// factory function for SQLiteRepo struct
func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{
		Conn: db,
	}
}

// create holdings table if not exists else ignore
func (repo *SQLiteRepo) Migrate() error {
	query := `
	create table if not exists holdings(
		id integer primary key autoincrement,
		amount real not null,
		purchase_date integer not null,
		purchase_price integer not null);
	`

	_, err := repo.Conn.Exec(query)
	return err
}

// insert one record to SQLite DB
func (repo *SQLiteRepo) InsertHolding(h Holding) (*Holding, error) {
	stmt := `
	insert into holdings (amount, purchase_date, purchase_price)
	values (?,?,?)
	`

	res, err := repo.Conn.Exec(stmt, h.Amount, h.PurchaseDate.Unix(), h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	h.ID = id
	return &h, nil
}

// get all records from SQLite DB
func (repo *SQLiteRepo) AllHoldings() ([]Holding, error) {
	query := `
	select id, amount, purchase_date, purchase_price
	from holdings
	order by purchase_date
	`

	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var all []Holding

	for rows.Next() {
		var h Holding
		var unixTime int64

		err = rows.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
		if err != nil {
			return nil, err
		}

		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}

	return all, nil
}

// get one record by id
func (repo *SQLiteRepo) GetHoldingByID(id int) (*Holding, error) {
	query := `
	select id, amount, purchase_date, purchase_price
	from holdings where id = ?
	`
	row := repo.Conn.QueryRow(query, id)

	var h Holding
	var unixTime int64

	err := row.Scan(&h.ID, &h.Amount, &unixTime, &h.PurchasePrice)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)
	return &h, nil
}

// update one record
func (repo *SQLiteRepo) UpdateHolding(id int64, updated Holding) error {
	// validate id
	if id == 0 {
		return errors.New("invalid updated id")
	}

	stmt := `
	update holdings set amount=?, purchase_date=?, purchase_price=? 
	where id=?`
	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	// check the number of affected rows by executing query
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

// delete one record
func (repo *SQLiteRepo) DeleteHolding(id int64) error {
	res, err := repo.Conn.Exec(`delete from holdings where id=?`, id)
	if err != nil {
		return err
	}

	// check the number of affected rows by executing query
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errDeleteFailed
	}

	return nil
}
