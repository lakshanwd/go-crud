package repository

import (
	"container/list"
	"log"

	"../dao"
	"../db"
)

//BookRepo - Book repository
type BookRepo struct {
	Name string
}

//GetBookRepository - returns book repository
func GetBookRepository() BookRepo {
	return BookRepo{Name: "tbl_book"}
}

//Select - Select books from db
func (repo BookRepo) Select() (*list.List, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("select book_id, book_name, book_author from tbl_book")
	if err != nil {
		log.Printf("%v", err)
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
	db, err := db.GetDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	book := doc.(dao.Book)
	stmt, err := db.Prepare("insert into tbl_book(book_name, book_author) values (?,?)")
	if err != nil {
		return nil
	}
	defer stmt.Close()
	_, err = stmt.Exec(book.BookName, book.Author)
	return err
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) error {
	book := doc.(dao.Book)
	db, err := db.GetDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("update tbl_book set book_name=?, book_author=? where book_id=?")
	defer stmt.Close()
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(book.BookName, book.Author, book.BookID)
	return err
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) error {
	book := doc.(dao.Book)
	db, err := db.GetDatabase()
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("delete from tbl_book where book_id=?")
	defer stmt.Close()
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(book.BookID)
	return err
}
