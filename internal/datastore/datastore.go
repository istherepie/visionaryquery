package datastore

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Result struct {
	TimeCode  string
	Actor     string
	Character string
}

type VisionaryStore struct {
	DB     *sql.DB
	CTX    context.Context
	Schema string
	Table  string
}

func (v *VisionaryStore) Generate() string {
	return fmt.Sprintf("SELECT TimeCode,Actor,Character FROM %v.%v WHERE Dataset=? AND Studio=?", v.Schema, v.Table)
}

func (v *VisionaryStore) Query(dataset string, studio string) (Result, error) {

	var result Result

	query := v.Generate()

	rows, err := v.DB.QueryContext(v.CTX, query, dataset, studio)

	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&result.TimeCode, &result.Actor, &result.Character); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return result, err
		}

	}

	return result, nil
}

func New(connection string) (vs VisionaryStore, err error) {

	vs = VisionaryStore{
		CTX: context.Background(),
	}

	db, err := sql.Open("sqlserver", connection)

	if err != nil {
		return
	}

	vs.DB = db
	return
}
