package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"os"
)

func GetClientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
func GetPanel() (*Panel, error) {
	filename := "panel.json"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return &Panel{
			Admins: make(map[string]Admin),
		}, nil
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(data) == 0 {
		return &Panel{
			Admins: make(map[string]Admin),
		}, nil
	}

	var panel Panel
	if err := json.Unmarshal(data, &panel); err != nil {
		log.Printf("Warning: parsing error, creating new panel: %v", err)
		return &Panel{
			Admins: make(map[string]Admin),
		}, nil
	}

	if panel.Admins == nil {
		panel.Admins = make(map[string]Admin)
	}

	return &panel, nil
}
