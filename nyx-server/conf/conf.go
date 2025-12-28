package conf

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

type Admin struct {
	RegisteredViaGood bool `json:"rvg"`
}
type Panel struct {
	Admins map[string]Admin `json:"admins"`
}

var Saltstr = Saltgen()

const MachinePassword = "test123"
