package db

import (
	"fmt"

	"github.com/Nhahan/blockchain/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		db, err := bolt.Open(dbName, 0600, nil)
		utils.HandleErr(err)

		err = db.Update(func(t *bolt.Tx) error {
			_, err = t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b", hash, data)
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte("checkpoint"), data)
		return err
	})
	utils.HandleErr(err)
}
