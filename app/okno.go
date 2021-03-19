package app

import (
	"github.com/oknors/okno/app/cfg"
	"github.com/oknors/okno/app/jdb"
	//csrc "github.com/oknors/okno/app/jorm/c/src"
	"github.com/oknors/okno/pkg/utl"
	"net/http"
	"time"
)

const (
	// HTTPMethodOverrideHeader is a commonly used
	// http header to override a request method.
	HTTPMethodOverrideHeader = "X-HTTP-Method-Override"
	// HTTPMethodOverrideFormKey is a commonly used
	// HTML form key to override a request method.
	HTTPMethodOverrideFormKey = "_method"
)

func NewOKNO() *OKNO {
	//jdb.JDB.Write("conf", "conf", cfg.CONFIG)
	err := jdb.JDB.Read("conf", "conf", &cfg.CONFIG)
	utl.ErrorLog(err)

	//go csrc.GetCoinSources()

	//fmt.Println(":ajdeeeeee", cfg.CONFIG)
	//go u.CloudFlare()
	o := &OKNO{

	}
	o.Hosts = o.GetHosts()

	srv := &http.Server{
		Handler: o.Handler(),
		Addr: ":" + cfg.CONFIG.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
