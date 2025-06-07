package provider

import "database/sql"

// DBProvider defines an interface for database providers
type DBProvider interface {
	Connect() (*sql.DB, error)
}
