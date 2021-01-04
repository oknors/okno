package main

import (
	"fmt"
	"github.com/oknors/okno/app"
	"log"
)

func main() {

	okno := app.NewOKNO()

	//log.Fatal(srv.ListenAndServeTLS("./cfg/server.pem", "./cfg/server.key"))
	fmt.Println("Listening on port:")

	log.Fatal(okno.Server.ListenAndServe())
	// port := 9898
	// fmt.Println("Listening on port:", port)
	// log.Fatal(http.ListenAndServe(":"+port, handlers.CORS()(r)))
}
