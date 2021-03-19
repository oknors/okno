package rts

import (
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
	"github.com/oknors/okno/app/our/tools"

)

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func NXFrame(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/coin/" + mux.Vars(r)["coin"] + "/" + mux.Vars(r)["frame"]
		p.ServeHTTP(w, r)
	}
}

func NXLib(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/" + mux.Vars(r)["libs"] + "/" + mux.Vars(r)["lib"]
		p.ServeHTTP(w, r)
	}
}

func NXWebPackHmr(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/__webpack_hmr"
		p.ServeHTTP(w, r)
	}
}

func NXNuxt(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nfile
		p.ServeHTTP(w, r)
	}
}
func NXNuxtSub(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nsub := vars["nsub"]
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nsub + "/" + nfile
		p.ServeHTTP(w, r)
	}
}
func NXNuxtSSub(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nsub := vars["nsub"]
		nssub := vars["nssub"]
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nsub + "/" + nssub + "/" + nfile
		p.ServeHTTP(w, r)
	}
}
func NXNuxtSSSub(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nsub := vars["nsub"]
		nssub := vars["nssub"]
		nsssub := vars["nsssub"]
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nsub + "/" + nssub + "/" + nsssub + "/" + nfile
		p.ServeHTTP(w, r)
	}
}
func NXNuxtSSSSub(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nsub := vars["nsub"]
		nssub := vars["nssub"]
		nsssub := vars["nsssub"]
		nssssub := vars["nssssub"]
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nsub + "/" + nssub + "/" + nsssub + "/" + nssssub + "/" + nfile
		p.ServeHTTP(w, r)
	}
}
func NXNuxtSSSSSub(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nsub := vars["nsub"]
		nssub := vars["nssub"]
		nsssub := vars["nsssub"]
		nssssub := vars["nssssub"]
		nsssssub := vars["nsssssub"]
		nfile := vars["nfile"]
		r.URL.Path = "/_nuxt/" + nsub + "/" + nssub + "/" + nsssub + "/" + nssssub + "/" + nsssssub + "/" + nfile
		p.ServeHTTP(w, r)
	}
}


func Frames(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["frame"] + "S" + mux.Vars(r)["file"]
	url := "http://127.0.0.1:3553/coin/" + mux.Vars(r)["frame"] + "/" + mux.Vars(r)["file"]
	data := tools.GetData(name, url)
	w.Write([]byte(data))
}

