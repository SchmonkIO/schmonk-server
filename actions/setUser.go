package actions

import (
	"encoding/json"

	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// SetUserAction is the action to set a users name and register the user at the server
type SetUserAction struct {
	models.BaseAction
	Name string `json:"name"`
}

// SetUser gets called to register a new user
func SetUser(player *models.BasePlayer, message []byte, mt int) {
	data := SetUserAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		models.SendJsonResponse(false, util.ActionSetUser, "invalid json", mt, player)
		return
	}
	player.Name = data.Name
	player.SetState(util.StateRoomList)
	global.Players.AddPlayer(player)
	models.SendJsonResponsePlayeriD(true, util.ActionSetUser, player.GetID(), mt, player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
