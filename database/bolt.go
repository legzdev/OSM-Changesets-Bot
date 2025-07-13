package database

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/legzdev/OSM-Changesets-Bot/types"
)

type Bolt struct {
	engine  *bolt.DB
	buckets buckets
}

type buckets struct {
	changesets []byte
}

func NewBolt(path string) (*Bolt, error) {
	engine, err := bolt.Open(path, 0666, bolt.DefaultOptions)
	if err != nil {
		return nil, err
	}

	db := &Bolt{
		engine: engine,
		buckets: buckets{
			changesets: []byte("changesets"),
		},
	}

	tx, err := engine.Begin(true)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists(db.buckets.changesets)
	if err != nil {
		return nil, err
	}

	return db, tx.Commit()
}

func (db *Bolt) GetLatest() (types.ChangesetID, error) {
	tx, err := db.engine.Begin(true)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(db.buckets.changesets)

	value := bucket.Get([]byte("latest"))
	if value == nil {
		return 0, ErrNotFound
	}

	return strconv.ParseInt(string(value), 10, 64)
}

func (db *Bolt) SetLatest(id types.ChangesetID) error {
	tx, err := db.engine.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(db.buckets.changesets)

	err = bucket.Put([]byte("latest"), fmt.Append([]byte{}, id))
	if err != nil {
		return err
	}

	return tx.Commit()
}
