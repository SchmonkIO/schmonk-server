package actions

import (
	"encoding/json"
	
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

type JoinAction struct {
	models.BaseAction
	Name string `json:"name"`
}

func Join(player *models.BasePlayer, message []byte, mt int) {
	data := JoinAction{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		util.LogToConsole(err.Error())
		player.Connection.WriteMessage(mt, []byte("invalid json"))
		return
	}
	player.Name = data.Name
	global.Players.AddPlayer(player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
