package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

//var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Amount    string `json:"amount"`
	Adress string `json:"adress"`
	//Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8010 and log any errors
	log.Println("http server started on :8010")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	//clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			//delete(clients, ws)
			break
		}
log.Printf(msg.Amount);
log.Printf(msg.Adress);
//log.Printf(msg.Message);
		
		// Send the newly received message to the broadcast channel
		//broadcast <- msg
	}
}

func handleMessages() {
	
}
