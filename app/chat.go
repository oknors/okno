package app

import (
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//initialize clients and broadcast channel
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Upgrader will allow us to turn a normal http connection into a WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}



func (o *OKNO) chat(rt *mux.Router) {
		rt.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			site := mux.Vars(r)["site"]
			if site == "chat" {
				// Upgrade http request to WebSocket
				ws, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					log.Fatal(err)
				}
				// Close connection when function is returned
				defer ws.Close()

				// Add new client to clients dictionary
				clients[ws] = true

				// Listen for any message and publish it to broadcast channel
				for {
					var msg Message
					err := ws.ReadJSON(&msg)
					if err != nil {
						log.Printf("error: %v", err)
						delete(clients, ws)
						break
					}
					broadcast <- msg
				}
			}
		})
	go handleMessages()
}

//Grab messages from broadcast channel and relays message
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
