package test

import (
	"database/sql"
	"fmt"
	"strings"
)

func rowsContains(rows *sql.Rows, expected string) error {
	var text string
	line := new(string)
	for rows.Next() {
		err := rows.Scan(line)
		if err != nil {
			return fmt.Errorf("Scan Error: %s", err.Error())
		}
		text += *line
	}
	if !strings.Contains(text, expected) {
		return fmt.Errorf("%s does not contain %s", text, expected)
	}

	return nil
}
