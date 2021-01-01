package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (o *OKNO) api(r *mux.Router) {
	////////////////
	// api.okno.RS
	////////////////
	s := r.Host("api.okno.rs").Subrouter()
	s.StrictSlash(true)
	//sh.Use(CommonMiddleware)
	so := s.PathPrefix("/cms").Subrouter()
	so.HandleFunc("/hosts", o.ReadHosts).Methods("GET")
	sd := s.PathPrefix("/data").Subrouter()
	sh := sd.PathPrefix("/{host}").Subrouter()
	sh.HandleFunc("/{col}", o.Database.ReadCollection).Methods("GET")
	sh.HandleFunc("/{col}/{slug}", o.Database.Read).Methods("GET")
	sh.HandleFunc("/{col}/{slug}", o.Database.Create).Methods("POST")
	sh.HandleFunc("/{col}/{slug}", o.Database.Update).Methods("PUT")
	sh.HandleFunc("/{col}/{slug}", o.Database.Delete).Methods("DELETE")
}

// Read all items from the database, unmarshaling the response.
func (o *OKNO) ReadHosts(w http.ResponseWriter, r *http.Request) {
	var hosts []map[string]string
	for _, h := range o.Hosts {
		host := map[string]string{
			"host": h.Host,
			"slug": h.Slug,
			"name": h.Name,
		}
		hosts = append(hosts, host)
	}
	js, err := json.Marshal(hosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
