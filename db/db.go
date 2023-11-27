package db

import (
	"context"

	"github.com/kentro-lab/cloudy-neigh/proto/vector"
)

type Vector = vector.Vector

type Knn struct {
	Vector Vector
	K      int
}

type SearchRequest struct {
	Knn Knn
}

type SearchResonse struct{}

type WriteOptions struct{}

type DB struct{}

func (d *DB) Get(ctx context.Context, key []byte) ([]byte, error) {
	var r []byte
	return r, nil
}

func (d *DB) Set(ctx context.Context, key []byte, value []byte, o *WriteOptions) error {
	return nil
}

func (d *DB) Delete(ctx context.Context, key []byte, o *WriteOptions) error {
	return nil
}

// Generic Search function including search by Vector.
func (d *DB) Search(ctx context.Context, req *SearchRequest) (*SearchResonse, error) {
	return &SearchResonse{}, nil
}

func (d *DB) Reader() Reader {
	// This should be a new instance of DB snapshot.
	return d
}

func (d *DB) Close() error {
	return nil
}

type Reader interface {
	Get(ctx context.Context, key []byte) ([]byte, error)
	Search(ctx context.Context, req *SearchRequest) (*SearchResonse, error)
	Close() error
}

type Writer interface {
	Set(ctx context.Context, key []byte, value []byte, o *WriteOptions) error
	Delete(ctx context.Context, key []byte, o *WriteOptions) error
	Reader() Reader
	Close() error
}

var _ Reader = (*DB)(nil)
var _ Writer = (*DB)(nil)

func NewDB() (*DB, error) {
	return &DB{}, nil
}
