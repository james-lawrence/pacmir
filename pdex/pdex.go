// Package pdex provides a package index.

package pdex

type Record struct {
	Digest []byte
	SHA256 []byte
	Name   string
}

// DB package information db.
type DB struct {
	root string
}

func New(dir string) *DB {
	return &DB{root: dir}
}

func (t *DB) Insert(precord Record) error {
	return nil
}

func (t *DB) Get(name string) (Record, error) {
	return Record{}, nil
}

