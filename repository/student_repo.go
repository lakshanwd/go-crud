package repository

import (
	"container/list"

	"../dao"
	"../db"
)

//StudentRepo - Student repository
type StudentRepo struct {
	Name string
}

//GetStudentRepository - returns student repository
func GetStudentRepository() StudentRepo {
	return StudentRepo{Name: "student"}
}

//Select - Select students from db
func (repo StudentRepo) Select() (*list.List, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select student_id, student_name, student_age, street_address, street_address_2, city, state, zip_code, country from tbl_student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := list.New()
	for rows.Next() {
		var student dao.Student
		var address dao.Address
		student.Address = &address
		err = rows.Scan(&student.StudentID, &student.Name, &student.Age, &address.StreetAddress, &address.StreetAddress2, &address.City, &address.State, &address.ZipCode, &address.Country)
		students.PushBack(student)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return students, err
}

//Insert - Insert Student to db
func (repo StudentRepo) Insert(doc interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	student := doc.(dao.Student)
	address := student.Address

	return executeInsert(db, "insert into tbl_student (student_name, student_age, street_address, street_address_2, city, state, zip_code, country) values (?,?,?,?,?,?,?,?)",
		student.Name, student.Age, address.StreetAddress, address.StreetAddress2, address.City, address.State, address.ZipCode, address.Country)
}

//Update - Update student
func (repo StudentRepo) Update(doc interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	student := doc.(dao.Student)
	address := student.Address
	return executeUpdateDelete(db, "update tbl_student set student_name=?, student_age=?, street_address=?, street_address_2=?, city=?, state=?, zip_code=?, country=? where student_id=?",
		student.Name, student.Age, address.StreetAddress, address.StreetAddress2, address.City, address.State, address.ZipCode, address.Country, student.StudentID)
}

//Remove - Delete student from db
func (repo StudentRepo) Remove(doc interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	student := doc.(dao.Student)

	return executeUpdateDelete(db, "delete from tbl_student where student_id=?", student.StudentID)
}
