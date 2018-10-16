package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// CreateRoomAction is the action that is used when creating rooms
type CreateRoomAction struct {
	models.BaseAction
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Map   string `json:"map"`
	Slots int    `json:"slots"`
}

// CreateRoom gets called to create a new room
func CreateRoom(player *models.BasePlayer, message []byte, mt int) {
	data := CreateRoomAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionCreateRoom, "invalid json", mt, player)
		return
	}
	r, err := models.CreateRoom(data.Name, data.Pass, player.GetID(), data.Slots)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionCreateRoom, err.Error(), mt, player)
		return
	}
	err = r.AddPlayer(player, data.Pass)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionCreateRoom, err.Error(), mt, player)
		return
	}
	global.Rooms.AddRoom(&r)
	player.SetState(util.StateLobby)
	models.SendJsonResponse(true, util.ActionCreateRoom, "created room", mt, player)
}
