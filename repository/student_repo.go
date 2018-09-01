package repository

import (
	"container/list"
	"database/sql"

	"github.com/lakshanwd/go-crud/dao"
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
	reader := func(rows *sql.Rows, collection *list.List) error {
		var student dao.Student
		var address dao.Address
		student.Address = &address
		err := rows.Scan(&student.StudentID, &student.Name, &student.Age, &address.StreetAddress, &address.StreetAddress2, &address.City, &address.State, &address.ZipCode, &address.Country)
		collection.PushBack(student)
		return err
	}
	return executeReader("select student_id, student_name, student_age, street_address, street_address_2, city, state, zip_code, country from tbl_student", reader)
}

//Insert - Insert Student to db
func (repo StudentRepo) Insert(doc interface{}) (int64, error) {
	student := doc.(dao.Student)
	address := student.Address
	return executeInsert("insert into tbl_student (student_name, student_age, street_address, street_address_2, city, state, zip_code, country) values (?,?,?,?,?,?,?,?)", student.Name, student.Age, address.StreetAddress, address.StreetAddress2, address.City, address.State, address.ZipCode, address.Country)
}

//Update - Update student
func (repo StudentRepo) Update(doc interface{}) (int64, error) {
	student := doc.(dao.Student)
	address := student.Address
	return executeUpdateDelete("update tbl_student set student_name=?, student_age=?, street_address=?, street_address_2=?, city=?, state=?, zip_code=?, country=? where student_id=?", student.Name, student.Age, address.StreetAddress, address.StreetAddress2, address.City, address.State, address.ZipCode, address.Country, student.StudentID)
}

//Remove - Delete student from db
func (repo StudentRepo) Remove(doc interface{}) (int64, error) {
	student := doc.(dao.Student)
	return executeUpdateDelete("delete from tbl_student where student_id=?", student.StudentID)
}
