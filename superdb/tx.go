package superdb

import "database/sql"

func ExecuteTransaction(db *sql.DB, tasks ...DbTask)  map[interface{}]interface{} {
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
	t:=&Tx{tx, map[interface{}]interface{}{}}
	for _, v := range tasks {
		v(t)
	}
	tx.Commit()
	return t.Data
}

type DbTask func(*Tx)

type Tx struct {
	*sql.Tx
	Data map[interface{}]interface{}
}

func (tx Tx) SetValue(key interface{}, value interface{}) {
	for k, _ := range tx.Data {
		if k == key {
			panic("the key has set value before")
		}
	}
	tx.Data[key] = value
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

func (tx Tx) Any(query string, args ...interface{}) bool {
	row := tx.QueryRow(query, args...)
	err:=row.Scan()
	return err!=sql.ErrNoRows
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
