package host

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/jdb"
	"net/http"
)

type Host struct {
	Name   string
	Slug   string
	Host   string
	Routes func(r *mux.Router)
}

func GetHosts(j *jdb.JDB) []*Host {
	return []*Host{
		//authOKNO(db),
		//admin(j),
		oknoRS(),
		sokno(),
		beliRS(),
		bitNodesNET(),
		comhttpORG(),
		comhttpUS(),

		djordjeMarcetinCOM(),
		marcetinCOM(),

		parallelcoinIO(),
		parallelcoinINFO(),
		gitParallelcoinINFO(),
		whitepaperParallelcoinINFO(),
		downloadParallelcoinINFO(),
		docsParallelcoinINFO(),
		explorerParallelcoinINFO(),
		logParallelcoinINFO(),
		legacyParallelcoinINFO(),
		punqRS(),
		solutionsRS(),
		vesicaPiescesORG(),
	}
}

func (h *Host) testRoutes(r *mux.Router) {
	s := h.sub(r)
	s.Host(h.Host).Path("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Host: %v\n", h.Name)
		})
}

func (h *Host) sub(r *mux.Router) *mux.Router {
	s := r.Host(h.Host).Subrouter()
	return s
}

func (h *Host) static() func(r *mux.Router) {
	return func(r *mux.Router) {
		s := h.sub(r)
		s.StrictSlash(true).PathPrefix("/").Handler(http.FileServer(http.Dir("js/public/" + h.Slug)))
	}
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
