package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

var (
	bucketName = []byte("cars")
)

type DB struct {
	conn *bbolt.DB
}

func NewDB(fileName string) (*DB, error) {
	conn, err := bbolt.Open(fileName, 0666, nil)
	if err != nil {
		return nil, err
	}

	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		conn.Close()
		return nil, err
	}

	db := DB{conn}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Insert(e Event) error {
	key := fmt.Sprintf("%s-%s", e.ID, uuid.NewString())
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(e); err != nil {
		return err
	}

	err := db.conn.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put([]byte(key), buf.Bytes())
	})
	return err
}

func (db *DB) Get(id string) ([]Event, error) {
	var events []Event

	err := db.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()
		prefix := []byte(id + "-")

		for k, v := c.Seek(prefix); k != nil; k, v = c.Next() {
			var e Event
			if err := gob.NewDecoder(bytes.NewReader(v)).Decode(&e); err != nil {
				return err
			}
			events = append(events, e)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return events, nil
}
