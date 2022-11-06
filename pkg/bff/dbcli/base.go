package dbcli

import (
	"database/sql"
	"fmt"
)

// shared logic for sql DBs
type BaseDBCli struct {
	driver, url string
	rootClient  *sql.DB
}

// DO NOT CALL. NO DRIVERS LOADED
func NewBaseDBCli(driver, hostName string) (*BaseDBCli, error) {
	db := &BaseDBCli{
		driver: driver,
		// root (user)
		// password (password)
		// using this string will connect w/o specifying a db,
		// concat with a dbName to access that db
		url: fmt.Sprintf("root:password@tcp(%s:3306)/", hostName),
	}

	var err error
	db.rootClient, err = db.makeCli("")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (b *BaseDBCli) makeDB(dbName string) error {
	_, err := b.rootClient.Exec(fmt.Sprintf("CREATE DATABASE [IF NOT EXISTS] %s", dbName))
	return err
}

func (b *BaseDBCli) makeCli(dbName string) (*sql.DB, error) {
	return sql.Open(b.driver, b.url+dbName)
}

// EnsureDBandClient ensures that the DB exists, then creates a client for it
func (b *BaseDBCli) EnsureDBandCli(dbName string) (*sql.DB, error) {
	err := b.makeDB(dbName)
	if err != nil {
		return nil, err
	}
	return b.makeCli(dbName)
}
