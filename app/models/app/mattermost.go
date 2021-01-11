package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func (o *OKNO) mattermost(r *mux.Router) {
	////////////////
	// mattermost
	////////////////
	s := r.Host("mm.okno.rs").Subrouter()

	//s.HandleFunc("/", forward)
	//s.HandleFunc("/static/{file}", static)
	//s.Headers("X-Requested-With", "XMLHttpRequest")

	// start server
	s.HandleFunc("/", handleRequestAndRedirectHome)
	s.HandleFunc("/{path}", handleRequestAndRedirect)
	t := s.PathPrefix("/static").Subrouter()
	t.HandleFunc("/{path}", handleRequestAndRedirectStatic)
	t.Headers("X-Forwarded-For")
	t.Headers("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	t.Headers("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	t.Headers("Access-Control-Allow-Origin", "*")
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	//req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	//req.Host = url.Host
	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirectHome(w http.ResponseWriter, r *http.Request) {
	u := "http://192.168.192.169:8065"
	log.Printf("proxy_condition:", u)
	serveReverseProxy(u, w, r)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	u := "http://192.168.192.169:8065/" + v["path"]
	log.Printf("proxy_condition:", u)
	serveReverseProxy(u, w, r)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirectStatic(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	u := "http://192.168.192.169:8065/static/" + v["path"]

	log.Printf("proxy_condition:", u)

	serveReverseProxy(u, w, r)
}
