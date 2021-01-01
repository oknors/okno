package app

import (
	"fmt"
	"github.com/p9c/okno/app/config"
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/app/models/app"
	"github.com/p9c/okno/app/models/host"
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

func NewOKNO() *app.OKNO {
	conf, err := config.GetConfiguration()
	if err != nil {
		fmt.Println("Error", err)
	}
	o := &app.OKNO{
		Configuration: &conf,
		Database:      jdb.NewJDB("./" + conf.DBName),
	}
	o.Hosts = host.GetHosts(o.Database)

	srv := &http.Server{
		//Handler: o.Handlers(),
		//Handler: interceptHandler(o.Handlers(), defaultErrorHandler),
		//Handler: interceptHandler(o.Handlers(), defaultErrorHandler),
		Handler: o.Handler(),
		//Handler: handlers.CORS()(handlers.CompressHandler(o.Handlers())),
		// Handler: cacheHandler(handlers.CORS()(handlers.CompressHandler(r))),
		// Handler: handlers.CompressHandler(r),
		Addr: ":" + conf.AppPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	o.Server = srv
	return o
}
