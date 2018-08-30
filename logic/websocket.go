package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/schmonk.io/schmonk-server/config"
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/models"
	"github.com/schmonk.io/schmonk-server/util"
)

var upgrader = websocket.Upgrader{}

// InitSocket initializes a new socket connection
func InitSocket(c *gin.Context) {
	if !config.Config.Server.CORS {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	}

	con, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	util.LogToConsole("socket upgrade request")
	if err != nil {
		util.LogToConsole("socket upgrade failed:", err)
		return
	}
	gSlots := global.Players.GetPlayerCount()
	if gSlots < config.Config.Server.Slots {
		CreateSocketPlayer(con)
	} else {
		con.WriteMessage(1, []byte("Slots exceeded"))
		con.Close()
		return
	}
}

// CreateSocketPlayer creates the player for the websocket connection
func CreateSocketPlayer(con *websocket.Conn) {
	util.LogToConsole("Connected Players:", global.Players.GetPlayerCount())
	player := models.CreateBasePlayer(con)
	defer con.Close()
	PlayerLoop(&player)
}
