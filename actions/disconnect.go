package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

func Disconnect(mt int, player *models.BasePlayer) {
	global.Players.RemovePlayer(player)
	//SendPlayerList()
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
