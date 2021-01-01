package handlers

//import (
//	"encoding/json"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"golang.org/x/crypto/bcrypt"
//	"time"
//)
//
//func (hs Handlers) FindOne(email, password string) (resp map[string]interface{}) {
//	hs.Collection("users")
//	user := &User{}
//	us, err := hs.ReadAll()
//	if err != nil {
//	}
//	for _, userInterface := range us {
//		u := User{}
//		if err := json.Unmarshal([]byte(userInterface), &u); err != nil {
//			fmt.Println("Error", err)
//		}
//		if u.Email != email {
//			resp = map[string]interface{}{"message": "Email address not found"}
//			return
//		} else {
//			user = &u
//		}
//	}
//	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
//	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
//		var resp = map[string]interface{}{"message": "Invalid login credentials. Please try again"}
//		return resp
//	}
//	tk := &models.Token{
//		UserID: user.ID,
//		Name:   user.Name,
//		Email:  user.Email,
//		StandardClaims: &jwt.StandardClaims{
//			ExpiresAt: expiresAt,
//		},
//	}
//	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
//	tokenString, error := token.SignedString([]byte("secret"))
//	if error != nil {
//		fmt.Println(error)
//	}
//
//	//var resp = map[string]interface{}{"status": false, "message": "logged in"}
//
//	resp = map[string]interface{}{"status": false, "token": tokenString}
//	return resp
//}
