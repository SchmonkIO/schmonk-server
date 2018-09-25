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
		if player.State == util.StateUndefined {
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
		} else if player.State != util.StateUndefined {
			ActionChooser(player, message, mt)
		} else {
			models.SendJsonResponse(false, "invalid player state", mt, player)
			continue
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
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole("createRoom")
		actions.CreateRoom(player, message, mt)
	case "getRooms":
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole("getRooms")
		actions.GetRooms(player, mt)
	case "joinRoom":
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole("joinRoom")
		actions.JoinRoom(player, message, mt)
	case "leaveRoom":
		if player.State == util.StateLobby || player.State == util.StateReady || player.State == util.StatePlaying {
			util.LogToConsole("leaveRoom")
			actions.JoinRoom(player, message, mt)
		}
		models.SendJsonResponse(false, "action not possible at this state", mt, player)
		return
	default:
		util.LogToConsole("Not implemented")
		models.SendJsonResponse(false, "action not implemented", mt, player)
	}
}
