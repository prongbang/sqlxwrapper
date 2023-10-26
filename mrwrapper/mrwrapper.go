package mrwrapper

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	errExeFail := errors.New("cannot add a child row")

	tx, err := conn.Begin()
	if err != nil {
		return tx, errExeFail
	}

	_, err = tx.Exec(query, args...)
	_ = tx.QueryRow("SELECT LAST_INSERT_ID()").Scan(dest...)

	if err == nil {
		return tx, nil
	}
	return tx, errExeFail
}

func Update(conn *sqlx.DB, query string, set string, params map[string]any) (*sql.Tx, error) {
	errExeFail := errors.New("cannot update a child row")

	tx, err := conn.Begin()
	if err != nil {
		return tx, errExeFail
	}

	if set == "" {
		return tx, errors.New("is not data to update")
	}
	query = fmt.Sprintf(query, set[1:])

	query, args, err := sqlx.Named(query, params)
	if err != nil {
		return tx, errExeFail
	}
	query = conn.Rebind(query)

	rs, err := tx.Exec(query, args...)
	if err != nil {
		return tx, errExeFail
	}

	_, err = rs.RowsAffected()
	if err != nil {
		return tx, errExeFail
	}

	return tx, nil
}

func Delete(conn *sqlx.DB, query string, args ...any) (*sql.Tx, error) {
	errExeFail := errors.New("cannot delete a child row")
	tx, err := conn.Begin()
	if err != nil {
		return tx, errExeFail
	}

	rs, err := tx.Exec(query, args...)
	if err != nil {
		return tx, errExeFail
	}

	row, err := rs.RowsAffected()
	if err != nil {
		return tx, errExeFail
	}

	if row <= 0 {
		return tx, errExeFail
	}

	return tx, nil
}

func Parameterize(index int, column int) string {
	format := ""
	var args []any
	for i := 1; i <= column; i++ {
		format += "$%d,"
		args = append(args, index*column+i)
	}
	format = fmt.Sprintf("(%s),", strings.TrimSuffix(format, ","))
	return fmt.Sprintf(format, args...)
}
