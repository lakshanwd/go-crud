package repository

import (
	"database/sql"

	"github.com/supunz/go-crud/db"
)

//StudentRepo - Student repository
type StudentRepo struct {
	Database *sql.DB
	Name     string
}

//GetStudentRepository - returns student repository
func GetStudentRepository() (StudentRepo, error) {
	db, err := db.GetDatabase()
	return StudentRepo{Database: db, Name: "student"}, err
}

//Select - Select students from db
func (repo StudentRepo) Select() ([]interface{}, error) {
	//to-do implement here
	return nil, nil
}

//Insert - Insert Student to db
func (repo StudentRepo) Insert(doc interface{}) error {
	//to-do implement here
	return nil
}

//Update - Update student
func (repo StudentRepo) Update(doc interface{}) error {
	//todo - implement here
	return nil
}

//Remove - Delete student from db
func (repo StudentRepo) Remove(doc interface{}) error {
	//to-do implement here
	return nil
}

//Close - close database
func (repo StudentRepo) Close() {

}
