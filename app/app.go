package app

import (
	"net/http"
)

type OKNO struct {
	Hosts         map[string]Host
	Server        *http.Server
}
