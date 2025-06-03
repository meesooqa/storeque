package db_provider

import "database/sql"

type DBProvider interface {
	Connect() (*sql.DB, error)
}
