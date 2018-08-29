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
		player.Connection.WriteMessage(mt, []byte("invalid json"))
		return
	}
	player.Name = data.Name
	global.Players.AddPlayer(player)
	player.Connection.WriteMessage(mt, []byte("set user"))
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
