package main

import (
	"log"
	"net/http"
	"nyx-server/panel"
)

func main() {
	http.HandleFunc("/ws", panel.HandleConnections)

	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
