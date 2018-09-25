package models

import (
	"encoding/json"
)

type StatusResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type StatusResponseRoomList struct {
	Status  bool     `json:"status"`
	Message RoomList `json:"message"`
}

type RoomList struct {
	Rooms map[string]*Room `json:"rooms"`
}

func SendJsonResponse(status bool, message string, mt int, player *BasePlayer) {
	resp := StatusResponse{}
	resp.Status = status
	resp.Message = message
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte("error serializing json"))
	}
	player.Connection.WriteMessage(mt, bytes)
}

func SendJsonResponseRoomList(status bool, rooms map[string]*Room, mt int, player *BasePlayer) {
	resp := StatusResponseRoomList{}
	rList := RoomList{}
	rList.Rooms = rooms
	resp.Status = status
	resp.Message = rList
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte("error serializing json"))
	}
	player.Connection.WriteMessage(mt, bytes)
}
