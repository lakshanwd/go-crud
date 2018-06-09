package repository

import "container/list"

//Repo - abstract repository interface
type Repo interface {
	Select(doc interface{}) (*list.List, error)
	Insert(doc interface{}) error
	Update(doc interface{}) error
	Remove(doc interface{}) error
}
