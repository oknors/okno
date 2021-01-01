package jdb

import (
	"fmt"
	"github.com/gorilla/schema"
	scribble "github.com/nanobox-io/golang-scribble"
)

var decoder = schema.NewDecoder()

// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

type JDB struct {
	db   *scribble.Driver
	path string
	col  string
}

func NewJDB(path string) *JDB {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(path, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	return &JDB{
		db: db,
	}
}

// Change host of JDB
func (j *JDB) Host(h string) *JDB {
	j.path = h
	return j
}
