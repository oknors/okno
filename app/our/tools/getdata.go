package tools

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/p9c/our/jdb"
)

type proxyPages struct {
	Time time.Time `json:"time"`
	Data []byte    `json:"data"`
}

var PxP = make(map[string]proxyPages)

func GetData(page, url string) template.HTML {
	gdb, err := jdb.OpenDB()
	if err != nil {
	}
	gPage := proxyPages{}
	gdb.Read("cache", "pages", &gPage)
	pTime := PxP[page].Time
	timeNow := time.Now()
	pTp := pTime.Add(time.Duration(1 * time.Minute))
	if timeNow.After(pTp) {
		gData, err := http.Get(url)
		if err != nil {
			fmt.Println("Get Proxy Dat Fail", gData)
		}
		defer gData.Body.Close()
		pD, err := ioutil.ReadAll(gData.Body)
		PxP[page] = proxyPages{
			Time: timeNow,
			Data: pD,
		}
		gdb.Write("cache", "pages", PxP)
	}
	return template.HTML(fmt.Sprint(string(PxP[page].Data)))
}
