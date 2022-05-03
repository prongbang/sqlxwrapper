package pqwrapper

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Count(conn *sqlx.DB, query string, args ...any) int64 {
	var id int64 = 0
	err := conn.Get(&id, query, args...)
	if err == nil {
		return id
	}
	return 0
}

func SelectOne[R any](conn *sqlx.DB, query string, args ...any) R {
	var row R
	var rows []R
	if err := conn.Select(&rows, query, args...); err != nil {
		return row
	}
	if len(rows) > 0 {
		return rows[0]
	}
	return row
}

func SelectList[R any](conn *sqlx.DB, query string, args ...any) []R {
	var rows []R
	if err := conn.Select(&rows, query, args...); err != nil {
		return []R{}
	}
	if len(rows) > 0 {
		return rows
	}
	return []R{}
}

func Create(conn *sqlx.DB, query string, dest []any, args ...any) (*sql.Tx, error) {
	tx, err := conn.Begin()
	if err != nil {
		return tx, err
	}

	if err = tx.QueryRow(query, args...).Scan(dest...); err != nil {
		return tx, err
	}

	return tx, nil
}

func Update(conn *sqlx.DB, query string, set string, params map[string]any) (*sql.Tx, error) {
	tx, err := conn.Begin()

	if set == "" {
		return tx, errors.New("Is not data to update")
	}
	query = fmt.Sprintf(query, set[1:])

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return tx, err
	}
	query = conn.Rebind(query)

	if err != nil {
		return tx, err
	}

	rs, err := tx.Exec(query, args...)
	if err != nil {
		return tx, err
	}

	row, err := rs.RowsAffected()
	if err != nil {
		return tx, err
	}

	if row <= 0 {
		return tx, errors.New("Cannot update a child row")
	}

	return tx, nil
}

func Delete(conn *sqlx.DB, query string, args ...any) (*sql.Tx, error) {
	tx, err := conn.Begin()
	if err != nil {
		return tx, err
	}

	rs, err := tx.Exec(query, args...)
	if err != nil {
		return tx, err
	}

	row, err := rs.RowsAffected()
	if err != nil {
		return tx, err
	}

	if row <= 0 {
		return tx, errors.New("Cannot delete a child row")
	}

	return tx, nil
}
