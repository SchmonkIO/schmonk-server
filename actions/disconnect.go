package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// Disconnect disconnects a player from the server
func Disconnect(mt int, player *models.BasePlayer) {
	global.Players.RemovePlayer(player)
	//SendPlayerList()
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
