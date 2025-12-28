package panel

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RDLxxx/Nyx/nyx-server/conf"

	"github.com/gorilla/websocket"
)

func PanelCommands(message string, ip string, ws *websocket.Conn, messageType int) {
	switch {
	case message == "Initial":
		log.Printf("Initial(admn) from %s", ip)
		if err := ws.WriteMessage(messageType, []byte("client auth: true")); err != nil {
			log.Println("echo>", err)
		}
	case message == "getsalt":
		log.Printf("getsalt from %s", ip)
		if err := ws.WriteMessage(messageType, []byte(conf.Saltstr)); err != nil {
			log.Println("echo>", err)
		}
	}
}

func UnknownPanelCommands(message string, ip string, ws *websocket.Conn, messageType int) {
	switch {
	case strings.HasPrefix(message, "regpanel|"):
		hpass := message[9:]
		log.Printf("regpanel from %s, hpass %s; salt %s", ip, hpass, conf.Saltstr)

		err := RegisterClientSimple(ip, hpass)
		if err != nil {
			log.Printf("Error registering client %s: %v", ip, err)
		}
	case message == "getsalt":
		log.Printf("getsalt from %s", ip)
		if err := ws.WriteMessage(messageType, []byte(conf.Saltstr)); err != nil {
			log.Println("echo>", err)
		}
	case message == "Initial":
		log.Printf("Initial(unkw) from %s", ip)
		is := conf.IsGoodClient(ip)
		stris := strconv.FormatBool(is)
		if err := ws.WriteMessage(messageType, []byte(stris)); err != nil {
			log.Println("echo>", err)
		}
	}

}

func RegisterClientSimple(ip, pass string) error {
	panel, err := conf.GetPanel()
	if err != nil {
		return err
	}

	// Проверяем существующего клиента
	if adm, exists := panel.Admins[ip]; exists {
		if adm.RegisteredViaGood {
			return fmt.Errorf("IP %s уже зарегистрирован корректно", ip)
		}
		// Если клиент существует, но RegisteredViaGood = false,
		// продолжаем выполнение для перерегистрации
	}

	// Этот код выполнится для:
	// 1. Нового клиента (IP не существует)
	// 2. Существующего клиента с RegisteredViaGood = false (перерегистрация)
	rvg := conf.VerifyPassword(pass)

	panel.Admins[ip] = conf.Admin{
		RegisteredViaGood: rvg,
	}

	// Сохраняем данные
	file, err := os.Create("panel.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(panel); err != nil {
		return err
	}

	return nil
}
