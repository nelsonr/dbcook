package lib

import (
	"database/sql"
)

// GetDBConnection tries to establish a connection to a database with the given
// driver name and data source name. The connection is tested before returning.
// A database handle is returned if the connection test is successful, otherwise
// an error is returned.
//
// The consumer is responsible for ensuring that the connection is properly closed by
// calling Close() on the returned database handle.
func GetDBConnection(driverName string, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
