package database

import "database/sql"

type Db interface {
	Connect() (*sql.DB, error)
	Close() error
}

type DbImpl struct {
	db *sql.DB
}

func (d DbImpl) Connect() (*sql.DB, error) {
	dbConn := "root:123456@tcp(127.0.0.1:3306)/web"
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		_ = d.Close()
		return nil, err
	}

	return d.db, nil
}

func (d DbImpl) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}
