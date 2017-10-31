package dao

//Book - book dao
type Book struct {
	BookID   int    `json:"bookID,omitempty"`
	BookName string `json:"bookName,omitempty"`
	Author   string `json:"author,omitempty"`
}
