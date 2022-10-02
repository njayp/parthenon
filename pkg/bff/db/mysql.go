package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MYSQL struct {
	driver, source string
}

func NewMYSQL() *MYSQL {
	return &MYSQL{
		driver: "mysql",
		source: "root:password@tcp(mysql.default:3306)/testo",
	}
}

func (m *MYSQL) Query(query string) (string, error) {
	db, err := sql.Open(m.driver, m.source)
	if err != nil {
		return "", err
	}
	defer db.Close()
	results, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer results.Close()

	var text string
	line := new(string)
	for results.Next() {
		err = results.Scan(line)
		if err != nil {
			return "", err
		}
		text = fmt.Sprintf("%s%s\n", text, *line)
	}
	return text, nil
}
