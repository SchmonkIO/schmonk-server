package logic

import (
	"strconv"

	"github.com/schmonk.io/schmonk-server/actions"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

func PlayerLoop(player *models.BasePlayer) {
	for {
		mt, message, err := player.Connection.ReadMessage()
		util.LogToConsole("MT: " + strconv.Itoa(mt))
		util.LogToConsole("Message: " + string(message))
		if err != nil {
			if mt == -1 {
				actions.Disconnect(mt, player)
				util.LogToConsole("disconnect player:", err)
			} else {
				util.LogToConsole("message read error:", err)
			}
			break
		}
		if player.Name == "" {
			//Naming Action -> Init player complete
			baseAction := models.BaseAction{}
			err := baseAction.Unmarshal(message)
			if err != nil {
				player.Connection.WriteMessage(mt, []byte("wrong json format"))
			} else {
				if baseAction.Check("join") {
					actions.Join(player, message, mt)
				} else {
					player.Connection.WriteMessage(mt, []byte("set name first"))
				}
			}
		} else {
			//Action choosing
			util.LogToConsole("Done did it!")
		}
	}
}
