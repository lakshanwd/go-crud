package repository

import (
	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/db"
	"gopkg.in/mgo.v2"
)

//StudentBookRepo - StudentBook repository
type StudentBookRepo struct {
	Database *mgo.Database
	Name     string
}

//GetStudentBookRepository -sdsf
func GetStudentBookRepository() (StudentBookRepo, error) {
	db, err := db.GetDatabase()
	return StudentBookRepo{Database: db, Name: "student_book"}, err
}

//Select - Select books from db
func (repo StudentBookRepo) Select() ([]dao.StudentBook, error) {
	//todo - implement here
	return nil, nil
}

//Insert - Insert books to db
func (repo StudentBookRepo) Insert(studentBook dao.StudentBook) error {
	//todo - implement here
	return nil
}

//Update - Update books
func (repo StudentBookRepo) Update(studentBook dao.StudentBook) error {
	//todo - implement here
	return nil
}

//Remove - Delete books from db
func (repo StudentBookRepo) Remove(studentBook dao.StudentBook) error {
	//todo - implement here
	return nil
}

//Close - close database
func (repo StudentBookRepo) Close() {
	if repo.Database != nil && repo.Database.Session != nil {
		repo.Database.Session.Close()
	}
}
