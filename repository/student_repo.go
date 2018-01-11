package repository

import (
	"database/sql"

	"github.com/supunz/go-crud/dao"
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
	rows, err := repo.Database.Query("select student_id, student_name, student_age, street_address, street_address_2, city, state, zip_code, country from tbl_student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student dao.Student
		var address dao.Address
		student.Address = &address
		err = rows.Scan(&student.StudentID, &student.Name, &student.Age, &address.StreetAddress, &address.StreetAddress2, &address.City, &address.State, &address.ZipCode, &address.Country)
		//todo assign student to a list
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return nil, err
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
