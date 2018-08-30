package logic

import (
	"strconv"

	"github.com/schmonk.io/schmonk-server/actions"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// PlayerLoop handles every websocket message and calls appropriate functions
func PlayerLoop(player *models.BasePlayer) {
	for {
		mt, message, err := player.Connection.ReadMessage()
		util.LogToConsole("MT: " + strconv.Itoa(mt))
		util.LogToConsole("Message: " + string(message))
		if err != nil {
			if mt == -1 {
				actions.Disconnect(mt, player)
				util.LogToConsole("disconnect player:", err)
			} else {
				util.LogToConsole("message read error:", err)
			}
			break
		}
		if player.Name == "" {
			baseAction := models.BaseAction{}
			err := baseAction.Unmarshal(message)
			if err != nil {
				models.SendJsonResponse(false, "invalid json", mt, player)
				continue
			}
			if !baseAction.Check("setUser") {
				models.SendJsonResponse(false, "set name first", mt, player)
				continue
			}
			actions.SetUser(player, message, mt)
		} else {
			ActionChooser(player, message, mt)
		}
	}
}

// ActionChooser handles every action per player and calls corresponding functions
func ActionChooser(player *models.BasePlayer, message []byte, mt int) {
	baseAction := models.BaseAction{}
	err := baseAction.Unmarshal(message)
	if err != nil {
		models.SendJsonResponse(false, "invalid json", mt, player)
		return
	}
	switch baseAction.Action {
	case "createRoom":
		util.LogToConsole("createRoom")
		actions.CreateRoom(player, message, mt)
	case "getRooms":
		util.LogToConsole("getRooms")
		actions.GetRooms(player, mt)
	case "joinRoom":
		util.LogToConsole("joinRoom")
		actions.JoinRoom(player, message, mt)
	default:
		util.LogToConsole("Not implemented")
		models.SendJsonResponse(false, "action not implemented", mt, player)
	}
}
