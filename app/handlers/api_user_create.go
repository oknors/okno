package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Create inserts a new post to the database
//CreateUser function -- create a new user
func (hs *Handlers) APICreateUser(w http.ResponseWriter, r *http.Request) {
	hs.Collection("users")
	//hs.JDB.Collection("users")
	//user := &User{}
	//json.NewDecoder(r.Body).Decode(user)
	//
	//pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	fmt.Println(err)
	//	err := ErrorResponse{
	//		Err: "Password Encryption  failed",
	//	}
	//	json.NewEncoder(w).Encode(err)
	//}
	//user.ID = uuid.Must(uuid.NewV4(), nil).String()
	//user.Password = string(pass)
	//
	//createdUser := hs.Write(user.ID, user)
	//var errMessage = createdUser.Error
	//
	//if createdUser.Error != nil {
	//	fmt.Println(errMessage)
	//}
	//json.NewEncoder(w).Encode(createdUser)

	fmt.Println("assssssssssssss")
}

//FetchUser function
func (hs Handlers) APIFetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	//db.Preload("auths").Find(&users)
	us, err := hs.ReadAll()
	if err != nil {
	}

	for _, userInterface := range us {
		var u User
		if err := json.Unmarshal([]byte(userInterface), &u); err != nil {
			fmt.Println("Error", err)
		}
		users = append(users, u)
	}
	json.NewEncoder(w).Encode(users)
}

func (hs Handlers) APIUpdateUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//var id = params["id"]
	//user, err := hs.Read(id)
	//if err != nil {
	//}
	//json.NewDecoder(r.Body).Decode(user)
	//hs.Write(id, user)
	//json.NewEncoder(w).Encode(&user)
}
func (hs Handlers) APIDeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	//user,err := hs.Read(id)
	//if err != nil {}
	hs.Delete(id)
	json.NewEncoder(w).Encode("User deleted")
}
func (hs Handlers) APIGetUser(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//var id = params["id"]
	//var user User
	//db.First(&user, id)
	//json.NewEncoder(w).Encode(&user)
	json.NewEncoder(w).Encode("User dess")

}
