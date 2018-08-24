package repository

import (
	"container/list"
	"database/sql"

	"github.com/lakshanwd/go-crud/dao"
	"github.com/lakshanwd/go-crud/db"
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

	reader := func(rows *sql.Rows, collection *list.List) error {
		var book dao.Book
		err = rows.Scan(&book.BookID, &book.BookName, &book.Author)
		collection.PushBack(book)
		return err
	}

	return executeReader(db, "select book_id, book_name, book_author from tbl_book", reader)
}

//Insert - Insert books to db
func (repo BookRepo) Insert(doc interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return -1, err
	}
	defer db.Close()
	book := doc.(dao.Book)
	return executeInsert(db, "insert into tbl_book(book_name, book_author) values (?,?)", book.BookName, book.Author)
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) (int64, error) {
	db, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	book := doc.(dao.Book)

	return executeUpdateDelete(db, "update tbl_book set book_name=?, book_author=? where book_id=?", book.BookName, book.Author, book.BookID)
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) (int64, error) {
	book := doc.(dao.Book)
	db, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	return executeUpdateDelete(db, "delete from tbl_book where book_id=?", book.BookID)
}
