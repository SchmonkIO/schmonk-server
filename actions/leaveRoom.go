package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

type LeaveRoomAction struct {
	models.BaseAction
}

func LeaveRoom(player *models.RoomPlayer, message []byte, mt int) {
	data := LeaveRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, "invalid json", mt, &player.BasePlayer)
		return
	}
	err = global.Rooms.RemovePlayer(player)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, err.Error(), mt, &player.BasePlayer)
		return
	}
	player.SetState(util.StateRoomList)
	models.SendJsonResponse(true, "left room", mt, &player.BasePlayer)
}
