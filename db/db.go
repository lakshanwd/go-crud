package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "root:1234@/library")
}
