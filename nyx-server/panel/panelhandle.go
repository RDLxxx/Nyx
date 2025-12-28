package panel

import (
	"log"
	"net/http"
	"time"

	"github.com/RDLxxx/Nyx/nyx-server/conf"

	"github.com/gorilla/websocket"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {

	ws, err := conf.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	ip := conf.GetClientIP(r)

	isGoodClient := conf.IsGoodClient(ip)

	for {
		ws.SetReadDeadline(time.Now().Add(60 * time.Second))

		messageType, message, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("WebSocket error from %s: %v", ip, err)
			}
			break
		}

		if isGoodClient {
			PanelCommands(string(message), ip, ws, messageType)
		} else {
			UnknownPanelCommands(string(message), ip, ws, messageType)
		}
	}
}
