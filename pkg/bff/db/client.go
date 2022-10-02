package db

type DB interface {
	Query(query string) (string, error)
}
