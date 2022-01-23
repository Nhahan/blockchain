package db

import (
	"github.com/Nhahan/blockchain/utils"
	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open("blockcahin.db", 0600, nil)
		utils.HandleErr(err)
		db = dbPointer
	}
	return db
}
