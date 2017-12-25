package db

import "database/sql"

//_ github.com/go-sql-driver/mysql this is neaded
import _ "github.com/go-sql-driver/mysql"

//GetDatabase - returns a Database object
func GetDatabase() (*sql.DB, error) {
	return sql.Open("mysql", "root:123@/library")
}
