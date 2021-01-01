package main

import (
	"github.com/p9c/okno/app"
	"log"
)

func main() {

	okno := app.NewOKNO()

	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))

	log.Fatal(okno.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
