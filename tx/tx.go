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

func (tx Tx) ScanRows(query string, args ...interface{}) []map[string]interface{} {
	rows, err := tx.Query(query, args...)
	if err != nil {
		panic(err)
	}
	sr := scanRows(&Rows{rows})
	return sr
}

func (tx Tx) ScanRow(query string, args ...interface{}) map[string]interface{} {
	sr := tx.ScanRows(query, args...)
	if len(sr) == 0 {
		return nil

	} else if len(sr) > 1 {
		panic("should only return one row data.")
	} else {
		return sr[0]
	}
}
func scanRows(rows *Rows) []map[string]interface{} {
	result := []map[string]interface{}{}
	columns, err := rows.ColumnTypes()
	args := make([]interface{}, len(columns))
	for i, _ := range args {
		zero_str_p := new(string)
		args[i] = &zero_str_p
	}
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		rows.MustScan(args...)
		m := map[string]interface{}{}
		for i, _ := range args {
			val := *args[i].(**string)
			if val == nil {
				m[columns[i].Name()] = nil
			} else {
				m[columns[i].Name()] = *val
			}
		}
		result = append(result, m)
	}
	rows.Close()
	return result
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
