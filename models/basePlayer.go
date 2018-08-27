package models

import (
	"errors"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"

	"github.com/schmonk.io/schmonk-server/config"
)

// BasePlayer is struct for a basic player
type BasePlayer struct {
	ID   bson.ObjectId `json:"_id"`
	Name string        `json:"name"`
	//Status     string          `json:"status"`
	Connection *websocket.Conn `json:"-"`
}

// CreateBasePlayer creates and returns a new basic players
func CreateBasePlayer(con *websocket.Conn) BasePlayer {
	bp := BasePlayer{}
	bp.SetID()
	bp.SetConnection(con)
	return bp
}

// GetID returns the player id as a string
func (bp *BasePlayer) GetID() string {
	return bp.ID.Hex()
}

// SetID sets a random id for the player
func (bp *BasePlayer) SetID() {
	bp.ID = bson.NewObjectId()
}

// SetConnection sets the websocket connection for the player
func (bp *BasePlayer) SetConnection(con *websocket.Conn) {
	bp.Connection = con
}

// SetName sets the name for the player
func (bp *BasePlayer) SetName(name string) error {
	if len(name) <= config.Config.Game.NameLength {
		bp.Name = name
		return nil
	} else {
		return errors.New("Name to long")
	}
}