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
				models.SendJsonResponse(false, util.ActionNone, "invalid json", mt, player)
				continue
			}
			if !baseAction.Check(util.ActionSetUser) {
				models.SendJsonResponse(false, util.ActionNone, "set name first", mt, player)
				continue
			}
			actions.SetUser(player, message, mt)
		} else if player.State != util.StateUndefined {
			ActionChooser(player, message, mt)
		} else {
			models.SendJsonResponse(false, util.ActionNone, "invalid player state", mt, player)
			continue
		}
	}
}

// ActionChooser handles every action per player and calls corresponding functions
func ActionChooser(player *models.BasePlayer, message []byte, mt int) {
	baseAction := models.BaseAction{}
	err := baseAction.Unmarshal(message)
	if err != nil {
		models.SendJsonResponse(false, util.ActionNone, "invalid json", mt, player)
		return
	}
	switch baseAction.Action {
	case util.ActionCreateRoom:
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, util.ActionCreateRoom, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole(util.ActionCreateRoom)
		actions.CreateRoom(player, message, mt)
	case util.ActionGetRooms:
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, util.ActionGetRooms, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole(util.ActionGetRooms)
		actions.GetRooms(player, mt)
	case util.ActionJoinRoom:
		if player.State != util.StateRoomList {
			models.SendJsonResponse(false, util.ActionJoinRoom, "action not possible at this state", mt, player)
			return
		}
		util.LogToConsole(util.ActionJoinRoom)
		actions.JoinRoom(player, message, mt)
	case util.ActionLeaveRoom:
		if player.State == util.StateLobby || player.State == util.StateReady || player.State == util.StatePlaying {
			util.LogToConsole(util.ActionLeaveRoom)
			actions.JoinRoom(player, message, mt)
		}
		models.SendJsonResponse(false, util.ActionLeaveRoom, "action not possible at this state", mt, player)
		return
	default:
		util.LogToConsole("Not implemented")
		models.SendJsonResponse(false, util.ActionNone, "action not implemented", mt, player)
	}
}
