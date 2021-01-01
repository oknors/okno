package app

import (
	"github.com/p9c/okno/app/config"
	"github.com/p9c/okno/app/jdb"
	"github.com/p9c/okno/app/models/host"
	"net/http"
)

type OKNO struct {
	Configuration *config.Config
	Database      *jdb.JDB
	Hosts         []*host.Host
	Server        *http.Server
}
