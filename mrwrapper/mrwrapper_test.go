package mrwrapper_test

import (
	"fmt"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
	"strings"
	"testing"
)

type Model struct {
	Column1 string `db:"column1"`
	Column2 string `db:"column2"`
	Column3 string `db:"column3"`
}

func BuildInsertQuery(models []Model) (string, []interface{}) {
	query := "INSERT INTO my_table (column1, column2, column3) VALUES "
	var values []any

	column := 3
	for i, model := range models {
		query += mrwrapper.Parameterize(i, column)
		values = append(values, model.Column1, model.Column2, model.Column3)
	}

	// Remove the trailing comma from the query string
	query = strings.TrimSuffix(query, ",")

	return query, values
}

func TestBuildInsertQuery(t *testing.T) {
	models := []Model{
		{Column1: "value1", Column2: "value2", Column3: "value3"},
		{Column1: "value4", Column2: "value5", Column3: "value6"},
		{Column1: "value7", Column2: "value8", Column3: "value9"},
	}

	query, values := BuildInsertQuery(models)

	fmt.Println(query)
	fmt.Println(values)
}

func TestParameterize(t *testing.T) {
	actual := mrwrapper.Parameterize(0, 4)

	fmt.Println(actual)
}
