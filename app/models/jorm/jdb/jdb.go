package jdb

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/oknors/okno/app/models/jorm"
	"github.com/oknors/okno/app/models/jorm/cfg"
)

// DB is the central database for jorm
var DB, _ = scribble.New(cfg.Dir, nil)

// // ReadData appends 'data' path prefix for a database read
func Read(collection, resource string) interface{} {
	ic := mod.Cache{}
	if err := DB.Read(cfg.Web+collection, resource, &ic); err != nil {
		fmt.Println("Error", err)
	}
	return ic.Data
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoin(slug string, v interface{}, d interface{}) bool {
	dc := mod.Cache{Data: d}
	return DB.Write(cfg.Web+"/coins", slug, v) == nil &&
		DB.Write(cfg.Web+"/data/"+slug, "info", dc) == nil
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoinImg(slug string, i interface{}) bool {
	ic := mod.Cache{Data: i}
	return DB.Write(cfg.Web+"/data/"+slug, "logo", ic) == nil
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoinData(slug, data string, d interface{}) bool {
	dc := mod.Cache{Data: d}
	return DB.Write(cfg.Web+"/data/"+slug, data, dc) == nil
}

// WriteExchange appends 'exchanges' path prefix for a database write
func WriteExchange(slug string, v interface{}) bool {
	return DB.Write(cfg.Web+"/exchanges", slug, v) == nil
}

// ReadCoins reads in all coin data in and converts to bytes for unmarshalling
func ReadData(v string) [][]byte {
	s, err := DB.ReadAll(v)
	if err != nil {
		fmt.Println("error reading data:", err.Error())
	}
	b := make([][]byte, len(s))
	for i := range s {
		b[i] = []byte(s[i])
	}
	return b
}
