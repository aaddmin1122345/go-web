package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbInit interface {
	Conn() (*sql.DB, error)
	Close() error
}

type DbInitImpl struct {
	Db *sql.DB
}

//func (d DbInitImpl) SetDb(db *sql.DB) {
//	d.Db = db
//}

func (d *DbInitImpl) Conn() (*sql.DB, error) {
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
	d.Db = db
	return d.Db, nil
}

func (d *DbInitImpl) Close() error {
	if d.Db != nil {
		return d.Db.Close()
	}
	return nil
}
