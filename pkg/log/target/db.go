package target

import "io"

type DB struct {
	DB    io.Writer
	Table string
}

func (db *DB) Write(p []byte) (int, error) {
	return db.DB.Write(p)
}
