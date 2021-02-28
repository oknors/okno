package host

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func authOKNO() *Host {
	////////////////
	// auth.okno.RS
	////////////////
	h := &Host{
		Name: "Authentication",
		Slug: "auth",
		Host: "a.okno.rs",
	}
	//hs := handlers.Handlers{jdb.NewJDB(db, h.Slug)}
	r := func(r *mux.Router) {
		s := h.sub(r)
		s.Use(CommonMiddleware)

		s.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("js/auth/"))))

		s.HandleFunc("/", homeHandler)
		s.HandleFunc("/metacortex", metacortexHandler)
		s.HandleFunc("/agents/{name}", agentsHandler)

		r.HandleFunc("/authenticate", authenticate)

		r.Handle("/api/megacity", authMiddleware(megacityHandler))
		r.Handle("/api/levrai", authMiddleware(levraiHandler))

	}
	h.Routes = r
	return h
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Matrix!"))
}
func metacortexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mr Anderson's not so secure workplace!"))
}
func agentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("My name is agent " + vars["name"]))
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	password := r.FormValue("password")
	if len(name) == 0 || len(password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide name and password to obtain the token"))
		return
	}
	if (name == "user123" && password == "user123") || (name == "morpheus" && password == "lawrence") {
		token, err := getToken(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Token: " + token))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}

func getToken(name string) (string, error) {
	signingKey := []byte("keymaker")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "redpill",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("keymaker")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

var megacityHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Megacity!"))
})

var levraiHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	if name != "neo" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Only Neo can enter the Merovingian's restaurant!"))
		return
	}
	w.Write([]byte("Welcome to the LeVrai!"))
})
