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
		models.SendJsonResponse(false, "invalid json", mt, player)
		return
	}
	player.Name = data.Name
	global.Players.AddPlayer(player)
	models.SendJsonResponse(true, "set user", mt, player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
