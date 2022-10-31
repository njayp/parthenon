package mysqlcli

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MYSQLCli struct {
	driver, address string
	conns           map[string]*sql.DB
}

func NewMYSQLCli(address string) *MYSQLCli {
	return &MYSQLCli{
		driver: "mysql",
		// root (user)
		// password (password)
		// using this string will connect w/o specifying a db,
		// concat with a dbName to access that db
		address: fmt.Sprintf("root:password@tcp(%s:3306)/", address),
		conns:   make(map[string]*sql.DB),
	}
}

func (m *MYSQLCli) ensureConn(dbName string) (conn *sql.DB, err error) {
	conn, ok := m.conns[dbName]
	if !ok {
		// conn is mia, so we need to create it
		conn, err = sql.Open(m.driver, m.address+dbName)
		if err != nil {
			return
		}
		// store conn for future use
		m.conns[dbName] = conn
	}
	return
}

func (m *MYSQLCli) CreateDB(dbName string) error {
	conn, err := m.ensureConn("")
	if err != nil {
		return err
	}

	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	return err
}

func (m *MYSQLCli) Query(dbName, query string) (text string, err error) {
	conn, err := m.ensureConn(dbName)
	results, err := conn.Query(query)
	if err != nil {
		return
	}
	defer results.Close()

	line := new(string)
	for results.Next() {
		err = results.Scan(line)
		if err != nil {
			return "", err
		}
		text = fmt.Sprintf("%s%s\n", text, *line)
	}

	return
}
