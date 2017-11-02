package repository

import (
	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/db"
	"gopkg.in/mgo.v2"
)

//BookRepo - Book repository
type BookRepo struct {
	Database *mgo.Database
	Name     string
}

//GetBookRepository - returns book repository
func GetBookRepository() (BookRepo, error) {
	db, err := db.GetDatabase()
	return BookRepo{Database: db, Name: "book"}, err
}

//Select - Select books from db
func (repo BookRepo) Select() ([]dao.Book, error) {
	//todo - implement here
	return nil, nil
}

//Insert - Insert books to db
func (repo BookRepo) Insert(book dao.Book) error {
	//todo - implement here
	return nil
}

//Update - Update books
func (repo BookRepo) Update(book dao.Book) error {
	//todo - implement here
	return nil
}

//Remove - Delete books from db
func (repo BookRepo) Remove(book dao.Book) error {
	//todo - implement here
	return nil
}

//Close - close database
func (repo BookRepo) Close() {
	if repo.Database != nil && repo.Database.Session != nil {
		repo.Database.Session.Close()
	}
}
