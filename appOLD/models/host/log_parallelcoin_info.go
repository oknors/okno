package host

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/tpl"
	"net/http"
)

func logParallelcoinINFO() *Host {
	////////////////
	// log.parallelcoin.INFO
	////////////////
	h := &Host{
		Name: "ParallelCoin Log",
		Slug: "log_parallelcoin_info",
		Host: "log.parallelcoin.info",
	}
	h.Routes = func(r *mux.Router) {
		s := r.Host("log.parallelcoin.info").Subrouter()
		s.HandleFunc("/", h.HomeHandler)

	}

	return h
}

// HomeHandler handles a request for (?)
func (h *Host) HomeHandler(w http.ResponseWriter, r *http.Request) {
	col := mux.Vars(r)["col"]
	//id := mux.Vars(r)["slug"]

	posts := h.Read(h.Slug, col)

	tpl.TemplateHandler().ExecuteTemplate(w, "index_gohtml", posts)
	fmt.Println("asasasas", posts)
}
