package repository

import (
	"container/list"
	"database/sql"

	"github.com/lakshanwd/go-crud/db"
)

//Repo - abstract repository interface
type Repo interface {
	Select(doc interface{}) (*list.List, error)
	Insert(doc interface{}) (int64, error)
	Update(doc interface{}) (int64, error)
	Remove(doc interface{}) (int64, error)
}

//DbConnection - Database Connectin Pool
var DbConnection *sql.DB

//SetupRepo - setup database connections
func SetupRepo() (err error) {
	DbConnection, err = db.GetDatabase()
	return
}

//CloseRepo - close database connections
func CloseRepo() {
	if DbConnection != nil {
		defer DbConnection.Close()
	}
}
