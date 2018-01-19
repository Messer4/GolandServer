package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


type Message struct {
	Amount    string `json:"amount"`
	Adress string `json:"adress"`
}

func main() {
	
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	
	http.HandleFunc("/ws", handleConnections)

	
	go handleMessages()

	
	log.Println("http server started on :8010")
	err := http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	defer ws.Close()

	

	for {
		var msg Message
		
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			
			break
		}
log.Printf(msg.Amount);
log.Printf(msg.Adress);

		
		
	}
}

func handleMessages() {
	
}
