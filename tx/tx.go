package tx

import "database/sql"

func NewDB(driverName, dataSourceName string) (*sql.DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	d.SetMaxIdleConns(0)
	return d, nil
}

type Tx struct {
	*sql.Tx
}

func NewTx(db *sql.DB) (*Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}
func (tx Tx) MustExec(query string, args ...interface{}) sql.Result {
	res, err := tx.Exec(query, args...)
	if err != nil {
		panic(err)
	}
	return res
}

func (tx Tx) MustQuery(query string, args ...interface{}) *Rows {
	rows, err := tx.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return &Rows{rows}
}

func (tx Tx) MustQueryRow(query string, args ...interface{}) *Row {
	row := tx.QueryRow(query, args...)
	return &Row{*row}
}

func (tx Tx) Any(query string, args ...interface{}) bool {
	row := tx.QueryRow(query, args...)
	err := row.Scan()
	return err != sql.ErrNoRows
}

type Row struct {
	sql.Row
}

func (row Row) MustScan(args ...interface{}) {
	err := row.Scan(args...)
	if err != nil {
		panic(err)
	}
}

type Rows struct {
	*sql.Rows
}

func (rows *Rows) MustScan(args ...interface{}) {
	err := rows.Scan(args...)
	if err != nil {
		panic(err)
	}
}

func (rows *Rows) Each(f func(*Rows)) {
	defer rows.Close()
	for rows.Next() {
		f(rows)
	}
	err := rows.Err()
	if err != nil {
		panic(err)
	}
}
