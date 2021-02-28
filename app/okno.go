package app

import (
	"fmt"
	"github.com/oknors/okno/app/config"
	"github.com/oknors/okno/app/jdb"
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
	conf, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("Error", err)
	}
	j, err := jdb.NewJDB(conf.Path, nil)
	if err != nil {
		utl.ErrorLog(err)
	}
	o := &OKNO{
		Configuration: &conf,
		Database:      j,
	}
	o.Hosts = o.GetHosts()

	srv := &http.Server{
		//Handler: o.Handlers(),
		//Handler: interceptHandler(o.Handlers(), defaultErrorHandler),
		//Handler: interceptHandler(o.Handlers(), defaultErrorHandler),
		Handler: o.Handler(),
		//Handler: handlers.CORS()(handlers.CompressHandler(o.Handler())),
		//Handler: cacheHandler(handlers.CORS()(handlers.CompressHandler(r))),
		// Handler: handlers.CompressHandler(r),
		Addr: ":" + conf.AppPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
