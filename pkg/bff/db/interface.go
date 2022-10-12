package db

//go:generate mockgen -destination=mocks/db_mock.go . DB
type DB interface {
	Query(query string) (string, error)
}
