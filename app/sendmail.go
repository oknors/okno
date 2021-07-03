package app

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/gorilla/mux"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Feedback is feedback struct
type Feedback struct {
	Name    string
	Email   string
	Message string
	URL string
}


func (o *OKNO) sendmail(r *mux.Router) {
	s := r.Host("sm.okno.rs").Subrouter()
	s.HandleFunc("/", sendmailHandler).Methods("POST")

	s.Headers("Access-Control-Allow-Origin", "*")
	s.Headers("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	s.Headers("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

//Handler is the default handler
func sendmailHandler(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/api/sendmail" || r.Method != "POST" {
	//	http.Error(w, "404 not found.", http.StatusNotFound)
	//	return
	//}

	var fb Feedback
	err := json.NewDecoder(r.Body).Decode(&fb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, body, err := SendMail(fb)
	if err != nil {
		println("Error sending Email: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(res)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(body))
	return
}

//SendMail sends the email using Sendgrid
func SendMail(f Feedback) (res int, out string, err error) {
	from := mail.NewEmail(f.Name, f.Email)
	subject := "ParallelCoin Info"
	to := mail.NewEmail("ParallelCoin Info", "info@parallelcoin.info")
	msgbody := "Page: " + f.URL + "<br/>" + "Message: <br/>" + f.Message 
	message := mail.NewSingleEmail(from, subject, to, "", msgbody)
	client := sendgrid.NewSendClient(os.Getenv("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"))
	r, err := client.Send(message)
	if err != nil {
		return r.StatusCode, r.Body, err
	}
	return r.StatusCode, r.Body, nil
}
