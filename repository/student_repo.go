package repository

import (
	"github.com/supunz/go-crud/dao"
	"github.com/supunz/go-crud/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StudentRepo - Student repository
type StudentRepo struct {
	Database *mgo.Database
}

//GetStudentRepository - returns student repository
func GetStudentRepository() (StudentRepo, error) {
	db, err := db.GetDatabase()
	return StudentRepo{Database: db}, err
}

//Select - Select students from db
func (repo StudentRepo) Select() ([]dao.Student, error) {
	studentCollection := repo.Database.C("student")
	var students []dao.Student
	err := studentCollection.Find(nil).All(&students)
	return students, err
}

//Insert - Insert Student to db
func (repo StudentRepo) Insert(student dao.Student) error {
	studentCollection := repo.Database.C("student")
	return studentCollection.Insert(student)
}

//Update - Update student
func (repo StudentRepo) Update(student dao.Student) error {
	//todo - implement here
	return nil
}

//Remove - Delete student from db
func (repo StudentRepo) Remove(student dao.Student) error {
	studentCollection := repo.Database.C("student")
	return studentCollection.Remove(bson.M{"Id": student.StudentID})
}

//Close - close database
func (repo StudentRepo) Close() {
	if repo.Database != nil && repo.Database.Session != nil {
		repo.Database.Session.Close()
	}
}
