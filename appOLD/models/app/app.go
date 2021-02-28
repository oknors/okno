package app

import (
	"github.com/oknors/okno/app/config"
	"github.com/oknors/okno/app/jdb"
	"net/http"
)

type OKNO struct {
	Configuration *config.Config
	Database      *jdb.JDB
	Hosts         []*Host
	Server        *http.Server
}
