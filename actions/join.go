package actions

import (
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

type JoinAction struct {
	models.BaseAction
	Name string `json:"name"`
}

func Join(player *models.BasePlayer, message []byte, mt int) {
	// data := models.NameRequest{}
	// err := json.Unmarshal(message, &data)
	// if err != nil {
	// 	util.LogToConsole(err.Error())
	// 	player.Connection.WriteMessage(mt, []byte("invalid json"))
	// } else {
	// 	player.Name = data.Name
	// 	player.Color = util.GetRandomColor()
	// 	player.Status = "lobby"
	// 	global.Mutex.Lock()
	// 	global.Players[player.ID.Hex()] = player
	// 	global.Mutex.Unlock()
	// 	util.LogToConsole(player.ID.Hex())
	// 	global.Mutex.Lock()
	// 	util.LogToConsole(global.Players[player.ID.Hex()])
	// 	global.Mutex.Unlock()
	// 	SendPlayerList(mt)
	// 	global.Mutex.Lock()
	// 	l := len(global.Players)
	// 	global.Mutex.Unlock()
	// 	util.LogToConsole("Connected Players:", l)
	// }

	data := JoinAction{}
	err := data.Unmarshal(message)
	if err != nil {
		util.LogToConsole(err.Error())
		player.Connection.WriteMessage(mt, []byte("invalid json"))
		return
	}
	player.Name = data.Name
	global.Players.AddPlayer(player)
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
}
