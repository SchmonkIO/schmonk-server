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

func JoinRoom(player *models.BasePlayer, message []byte, mt int) {
	data := JoinRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, "invalid json", mt, player)
		return
	}
	err = global.Rooms.AddPlayer(data.Id, data.Pass, player)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, "invalid room password", mt, player)
		return
	}
	models.SendJsonResponse(true, "joined room", mt, player)
}
