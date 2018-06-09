package db

import (
	"database/sql"
	//_ github.com/go-sql-driver/mysql this is neaded
	_ "github.com/go-sql-driver/mysql"
)

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	if db, err := sql.Open("mysql", "developer:123@/library"); err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(20)
		return db, nil
	} else {
		return nil, err
	}
}
