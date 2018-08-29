package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// GetRooms returns all available rooms to the requesting player
func GetRooms(player *models.BasePlayer, mt int) {
	bytes, err := global.Rooms.Marshal()
	if err != nil {
		util.LogToConsole(err.Error())
		player.Connection.WriteMessage(mt, []byte("failed to serialize room list"))
		return
	}
	player.Connection.WriteMessage(mt, bytes)
}
