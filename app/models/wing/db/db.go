package db

import (
	"encoding/json"
	"fmt"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/w-ingsolutions/c/model"
	"golang.org/x/text/unicode/norm"
	"unicode"
)

type DuoUIdb struct {
	DB     *scribble.Driver
	Folder string `json:"folder"`
	Name   string `json:"name"`
}

type Ddb interface {
	//DbReadAllTypes()map[int]model.WingCalGrupaRadova
	DbRead(folder, name string) model.WingVrstaRadova
	DbReadAll(folder string) map[int]model.WingVrstaRadova
	DbWrite(folder, name string, data interface{})
}

func DuoUIdbInit(dataDir string) (d *DuoUIdb) {
	d = new(DuoUIdb)
	db, err := scribble.New(dataDir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	d.DB = db
	return
}

var skip = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var safe = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

var _ Ddb = &DuoUIdb{}

//func (d *DuoUIdb) DbReadAllTypes() map[int]model.WingCalGrupaRadova{
//	items := make(map[int]model.WingCalGrupaRadova)
//	//types := []string{"assets", "config", "apps"}
//
//	//for t := range types {
//	//	items[t] = d.DbReadAll(t)
//	//}
//	for i := 1; i <= 31; i++ {
//		items[i] = d.DbReadAll(fmt.Sprint(i))
//	}
//	return items
//}
//func (d *DuoUIdb) DbReadTypeAll(f string) map[]model.WingVrstaRadova{
//	return d.DbReadAll(f)
//}

func (d *DuoUIdb) DbReadAll(folder string) map[int]model.WingVrstaRadova {
	itemsRaw, err := d.DB.ReadAll(folder)
	if err != nil {
		fmt.Println("Error", err)
	}
	items := make(map[int]model.WingVrstaRadova)
	for _, bt := range itemsRaw {
		item := model.WingVrstaRadova{}
		if err := json.Unmarshal([]byte(bt), &item); err != nil {
			fmt.Println("Error", err)
		}
		items[item.Id-1] = item
	}
	return items
}

func (d *DuoUIdb) DbRead(folder, name string) model.WingVrstaRadova {
	item := model.WingVrstaRadova{}
	if err := d.DB.Read(folder, name, &item); err != nil {
		fmt.Println("Error", err)
	}
	return item
}
func (d *DuoUIdb) DbWrite(folder, name string, data interface{}) {
	d.DB.Write(folder, name, data)
}

func slug(text string) string {
	buf := make([]rune, 0, len(text))
	dash := false
	for _, r := range norm.NFKD.String(text) {
		switch {
		case unicode.IsOneOf(safe, r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(skip, r):
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}
