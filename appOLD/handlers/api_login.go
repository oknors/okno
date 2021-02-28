package handlers

import (
	"encoding/json"
	"net/http"
)

//User struct declaration
type User struct {
	ID       string
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
}

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

//func (hs Handlers)TestAPI()func(w http.ResponseWriter, r *http.Request) {
//    return func(w http.ResponseWriter, r *http.Request) {
//        w.Write([]byte("API live and kicking"))
//    }
//}

func (hs Handlers) APILogin(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := &Request{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	//json.NewEncoder(w).Encode(hs.FindOne(user.Email, user.Password))
}
