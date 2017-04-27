package bolt

import (
	"os"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

// DB represents a handle to a Bolt database.
type DB struct {
	db *bolt.DB

	Path string
	Now  func() time.Time
}

// NewDB returns a new instance of DB.
func NewDB() *DB {
	return &DB{
		Now: time.Now,
	}
}

// Open opens and initializes the database.
func (db *DB) Open() error {
	// Create parent directory, if necessary.
	if err := os.MkdirAll(filepath.Dir(db.Path), 0700); err != nil {
		return err
	}

	// Open bolt database.
	d, err := bolt.Open(db.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	db.db = d

	return nil
}

// Close closes the database.
func (db *DB) Close() error {
	if db.db != nil {
		db.db.Close()
	}
	return nil
}

// Begin starts a new transaction.
func (db *DB) Begin(writable bool) (*Tx, error) {
	tx, err := db.db.Begin(writable)
	if err != nil {
		return nil, err
	}
	return &Tx{Tx: tx, Now: db.Now()}, nil
}

// Tx is a wrapper for bolt.Tx.
type Tx struct {
	*bolt.Tx
	Now time.Time
}