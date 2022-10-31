package db

//go:generate mockgen -destination=mocks/db_mock.go . DB
type DB interface {
	CreateDB(dbNames string) error
	Query(dbName, query string) (string, error)
}
