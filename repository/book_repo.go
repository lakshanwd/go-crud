package repository

import (
	"container/list"
	"database/sql"
	"log"

	util "github.com/lakshanwd/db-helper/mysql"
	"github.com/lakshanwd/go-crud/dao"
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
	reader := func(rows *sql.Rows, collection *list.List) {
		var book dao.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Author)
		collection.PushBack(book)
		log.Fatal(err)
	}
	return util.ExecuteReader(DbConnection, "select book_id, book_name, book_author from tbl_book", reader)
}

//Insert - Insert books to db
func (repo BookRepo) Insert(doc interface{}) (int64, error) {
	book := doc.(dao.Book)
	return util.ExecuteInsert(DbConnection, "insert into tbl_book(book_name, book_author) values (?,?)", book.BookName, book.Author)
}

//Update - Update books
func (repo BookRepo) Update(doc interface{}) (int64, error) {
	book := doc.(dao.Book)
	return util.ExecuteUpdateDelete(DbConnection, "update tbl_book set book_name=?, book_author=? where book_id=?", book.BookName, book.Author, book.BookID)
}

//Remove - Delete books from db
func (repo BookRepo) Remove(doc interface{}) (int64, error) {
	book := doc.(dao.Book)
	return util.ExecuteUpdateDelete(DbConnection, "delete from tbl_book where book_id=?", book.BookID)
}
