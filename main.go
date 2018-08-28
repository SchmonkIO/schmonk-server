package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/schmonk.io/schmonk-server/config"
	"github.com/schmonk.io/schmonk-server/global"
	"github.com/schmonk.io/schmonk-server/logic"
)

func main() {
	setup()
	sAddress := config.Config.Server.IP + ":" + strconv.Itoa(config.Config.Server.Port)
	router := gin.New()
	if config.Config.Server.Debug {
		router = gin.Default()
	}
	router.GET("/ws", func(c *gin.Context) {
		logic.InitSocket(c)
	})
	global.CreateGlobalPlayerList()
	global.CreateGlobalRoomList()
	log.Fatal(router.Run(sAddress))
}

func setup() {
	err := config.ReadConfig("server.conf")
	if err != nil {
		fmt.Println("[Failure] Could not read config file")
		os.Exit(1)
	}
	log.SetFlags(0)
}
