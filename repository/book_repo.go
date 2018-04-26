package repository

import (
	"container/list"
	"database/sql"

	"../dao"
	"../db"
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
func (repo BookRepo) Select() (*list.List, error) {
	rows, err := repo.Database.Query("select book_id, book_name, book_author from tbl_book")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := list.New()
	for rows.Next() {
		var book dao.Book
		err = rows.Scan(&book.BookID, &book.BookName, &book.Author)
		books.PushBack(book)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, err
}

//Insert - Insert books to db
func (repo BookRepo) Insert(doc interface{}) error {
	book := doc.(dao.Book)
	stmt, err := repo.Database.Prepare("insert into tbl_book(book_name, book_author) values (?,?)")
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(book.BookName, book.Author)
	return err
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) error {
	book := doc.(dao.Book)
	stmt, err := repo.Database.Prepare("update tbl_book set book_name=?, book_author=? where book_id=?")
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(book.BookName, book.Author, book.BookID)
	return err
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) error {
	book := doc.(dao.Book)
	stmt, err := repo.Database.Prepare("delete from tbl_book where book_id=?")
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(book.BookID)
	return err
}

//Close - close database
func (repo BookRepo) Close() {
	repo.Database.Close()
}
