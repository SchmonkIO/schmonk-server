package models

import (
	"encoding/json"
)

type StatusResponse struct {
	Status bool   `json:"status"`
	Action string `json:"action"`
}

type StatusResponseMessage struct {
	StatusResponse
	Message string `json:"message"`
}

type StatusResponseRoomList struct {
	StatusResponse
	Message RoomList `json:"roomList"`
}

type StatusResponsePlayerID struct {
	StatusResponse
	PlayerID string `json:"playerid"`
}

type RoomList struct {
	Rooms map[string]*Room `json:"rooms"`
}

func SendJsonResponse(status bool, action string, message string, mt int, player *BasePlayer) {
	resp := StatusResponseMessage{}
	resp.Status = status
	resp.Action = action
	resp.Message = message
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte("error serializing json"))
	}
	player.Connection.WriteMessage(mt, bytes)
}

func SendJsonResponsePlayeriD(status bool, action string, id string, mt int, player *BasePlayer) {
	resp := StatusResponsePlayerID{}
	resp.Status = status
	resp.Action = action
	resp.PlayerID = id
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte("error serializing json"))
	}
	player.Connection.WriteMessage(mt, bytes)
}

func SendJsonResponseRoomList(status bool, action string, rooms map[string]*Room, mt int, player *BasePlayer) {
	resp := StatusResponseRoomList{}
	rList := RoomList{}
	rList.Rooms = rooms
	resp.Status = status
	resp.Action = action
	resp.Message = rList
	bytes, err := json.Marshal(resp)
	if err != nil {
		player.Connection.WriteMessage(mt, []byte("error serializing json"))
	}
	player.Connection.WriteMessage(mt, bytes)
}
