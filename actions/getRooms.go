package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

// GetRooms returns all available rooms to the requesting player
func GetRooms(player *models.BasePlayer, mt int) {
	models.SendJsonResponseRoomList(true, util.ActionGetRooms, global.Rooms.GetRooms(), mt, player)
}
