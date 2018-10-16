package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

type JoinRoomAction struct {
	models.BaseAction
	Id   string `json:"id"`
	Pass string `json:"pass"`
}

// JoinRoom lets a player join a room if it is not full and the right password was provided
func JoinRoom(player *models.BasePlayer, message []byte, mt int) {
	data := JoinRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionJoinRoom, "invalid json", mt, player)
		return
	}
	err = global.Rooms.AddPlayer(data.Id, data.Pass, player)
	if err != nil {
		if err == util.ErrRoomNotFound {
			util.LogToConsole(err.Error())
			models.SendJsonResponse(false, util.ActionJoinRoom, err.Error(), mt, player)
			return
		}
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionJoinRoom, "invalid room password", mt, player)
		return
	}
	player.SetState(util.StateLobby)
	models.SendJsonResponse(true, util.ActionJoinRoom, "joined room", mt, player)
}
