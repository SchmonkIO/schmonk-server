package models

import "github.com/schmonk.io/schmonk-server/util"

// RoomPlayer is struct for a player in a room
type RoomPlayer struct {
	BasePlayer
	PosX  float32 `json:"posx"`
	PosY  float32 `json:"posy"`
	Color string  `json:"color"`
}

// CreateRoomPlayer creates and returns a new room player
func CreateRoomPlayer(bp BasePlayer) RoomPlayer {
	rp := RoomPlayer{}
	rp.BasePlayer = bp
	rp.SetColor()
	return rp
}

// SetColor sets a random color for the player
func (rp *RoomPlayer) SetColor() {
	rp.Color = util.GetRandomColor()
}
