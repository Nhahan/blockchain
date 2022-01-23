package db

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {

	}
	return db
}
