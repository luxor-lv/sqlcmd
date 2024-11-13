package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	connection *sqlx.DB
	showTypes  bool
}

func NewDatabaseConnection(driverType string, connectionString string, showTypes bool) (*Database, error) {
	db, error := sqlx.Open(driverType, connectionString)
	if error != nil {
		return nil, error
	}
	if error = db.Ping(); error != nil {
		return nil, error
	}
	return &Database{connection: db, showTypes: showTypes}, nil
}

func (db *Database) query(sqlCommand string) ([][]string, error) {
	rows, error := db.connection.Query(sqlCommand)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	columnTypes, error := rows.ColumnTypes()
	if error != nil {
		return nil, error
	}

	headers := make([]string, len(columnTypes))

	for i, column := range columnTypes {
		headers[i] = column.Name()
		if db.showTypes {
			headers[i] = headers[i] + " " + column.DatabaseTypeName()
		}
	}

	data := [][]string{}
	data = append(data, headers)

	for rows.Next() {
		rowData := make([]interface{}, len(columnTypes))
		for i := range rowData {
			rowData[i] = new(interface{})
		}

		rows.Scan(rowData...)

		rowStringData := make([]string, len(columnTypes))

		for i, col := range rowData {
			if *col.(*interface{}) == nil {
				rowStringData[i] = "(null)"
			} else {
				switch (*col.(*interface{})).(type) {
				case string, rune, []byte:
					rowStringData[i] = fmt.Sprintf("%s", *col.(*interface{}))
				default:
					rowStringData[i] = fmt.Sprintf("%v", *col.(*interface{}))
				}
			}
		}
		data = append(data, rowStringData)
	}

	return data, nil
}
