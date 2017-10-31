package db

import (
	"gopkg.in/mgo.v2"
)

// GetDatabase returns a Database object
//
// Once the mgo.Database is not useful anymore, mgo.Database.Close must be called to release the
// resources appropriately.
func GetDatabase() (*mgo.Database, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB("library"), nil
}
