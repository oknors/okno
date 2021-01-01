package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func (o *OKNO) Handler() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	for _, host := range o.Hosts {
		host.Routes(r)
	}
	o.api(r)
	return handlers.CORS()(handlers.CompressHandler(interceptHandler(r, defaultErrorHandler)))
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

type interceptResponseWriter struct {
	http.ResponseWriter
	errH func(http.ResponseWriter, int)
}

func (w *interceptResponseWriter) WriteHeader(status int) {
	if status >= http.StatusBadRequest {
		w.errH(w.ResponseWriter, status)
		w.errH = nil
	} else {
		w.ResponseWriter.WriteHeader(status)
	}
}

type ErrorHandler func(http.ResponseWriter, int)

func (w *interceptResponseWriter) Write(p []byte) (n int, err error) {
	if w.errH == nil {
		return len(p), nil
	}
	return w.ResponseWriter.Write(p)
}

func defaultErrorHandler(w http.ResponseWriter, status int) {
	//    const tpl = `
	//<!DOCTYPE html>
	//<html>
	//	<head>
	//		<meta charset="UTF-8">
	//	</head>
	//	<body>
	//		<div>{{ . }}</div>
	//
	//	</body>
	//</html>`
	t := template.Must(template.ParseFiles("js/errors/error.html"))
	//check := func(err error) {
	//    if err != nil {
	//        log.Fatal(err)
	//    }
	//}

	//t, err := template.New("webpage").Parse(tpl)
	//check(err)

	//err := template.GetRenderer("./js/errors/error.html", "").Render(w, "error", map[string]interface{}{"status": status})

	//tmpl := template.Must(template.ParseFiles("js/errors/error.html"))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t.Execute(w, map[string]interface{}{"status": status})
	//})
}

func interceptHandler(next http.Handler, errH ErrorHandler) http.Handler {
	if errH == nil {
		errH = defaultErrorHandler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(&interceptResponseWriter{w, errH}, r)
	})
}
