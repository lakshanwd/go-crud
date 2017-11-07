package repository

//Repo - abstract repository interface
type Repo interface {
	Select(doc interface{}) ([]interface{}, error)
	Insert(doc interface{}) error
	Update(doc interface{}) error
	Remove(doc interface{}) error
	Close(repo interface{})
}
