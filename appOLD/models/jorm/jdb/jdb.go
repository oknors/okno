package jdb

import (
	"fmt"
	"github.com/oknors/okno/appOLD/models/jorm/cfg"
)

// DB is the central database for jorm
var DB, _ = New(cfg.Dir, nil)

// // ReadData appends 'data' path prefix for a database read
func Read(collection, resource string) (i interface{}) {
	if err := DB.Read(cfg.Web+collection, resource, &i); err != nil {
		fmt.Println("Error", err)
	}
	return
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoin(slug string, v interface{}, d interface{}) bool {
	return DB.Write(cfg.Web+"/coins", slug, v) == nil &&
		DB.Write(cfg.Web+"/data/"+slug, "info", d) == nil
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoinImg(slug string, i interface{}) bool {
	return DB.Write(cfg.Web+"/data/"+slug, "logo", i) == nil
}

// WriteCoin appends 'coins' path prefix for a database write
func WriteCoinData(slug, data string, d interface{}) bool {
	return DB.Write(cfg.Web+"/data/"+slug, data, d) == nil
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
