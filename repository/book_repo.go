package repository

import (
	"database/sql"

	"github.com/supunz/go-crud/db"
)

//BookRepo - Book repository
type BookRepo struct {
	Database *sql.DB
	Name     string
}

//GetBookRepository - returns book repository
func GetBookRepository() (BookRepo, error) {
	db, err := db.GetDatabase()
	return BookRepo{Database: db, Name: "tbl_book"}, err
}

//Select - Select books from db
func (repo BookRepo) Select() ([]interface{}, error) {
	//to-do implement here
	return nil, nil
}

//Insert - Insert books to db
func (repo BookRepo) Insert(doc interface{}) error {
	//to-do implement here
	return nil
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) error {
	//to-do implement here
	return nil
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) error {
	//to-do implement here
	return nil
}

//Close - close database
func (repo BookRepo) Close() {

}
