package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
)

// GetRooms returns all available rooms to the requesting player
func GetRooms(player *models.BasePlayer, mt int) {
	models.SendJsonResponseRoomList(true, global.Rooms.GetRooms(), mt, player)
}
