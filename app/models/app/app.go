package app

import (
	"github.com/oknors/okno/app/config"
	"github.com/oknors/okno/app/jdb"
	"github.com/oknors/okno/app/models/host"
	"net/http"
)

type OKNO struct {
	Configuration *config.Config
	Database      *jdb.JDB
	Hosts         []*host.Host
	Server        *http.Server
}
