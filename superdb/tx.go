package superdb

import "database/sql"

func ExecuteTransaction(db *sql.DB, tasks ...DbTask) {
	var tx *sql.Tx
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()
	for _, v := range tasks {
		v(&Tx{tx})
	}
	tx.Commit()
}

type DbTask func(*Tx)

type Tx struct {
	*sql.Tx
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
	return &Rows{*rows}
}

func (tx Tx) MustQueryRow(query string, args ...interface{}) *Row {
	row := tx.QueryRow(query, args...)
	return &Row{*row}
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
	sql.Rows
}

func (rows Rows) MustScan(args ...interface{}) {
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
