package global

import (
	"encoding/json"
	"sync"

	"github.com/schmonk.io/schmonk-server/models"
)

// Rooms is a list of all available rooms
var Rooms RoomList

// RoomList is a struct for a list of Rooms
type RoomList struct {
	Rooms map[string]*models.Room `json:"rooms"`
	Mut   *sync.Mutex             `json:"-"`
}

func CreateGlobalRoomList() {
	grList := RoomList{}
	grList.Rooms = map[string]*models.Room{}
	grList.Mut = &sync.Mutex{}
	Rooms = grList
}

func (rl *RoomList) Marshal() ([]byte, error) {
	rl.Mut.Lock()
	bytes, err := json.Marshal(rl)
	rl.Mut.Unlock()
	return bytes, err
}

func (rl *RoomList) GetRooms() map[string]*models.Room {
	rl.Mut.Lock()
	rooms := rl.Rooms
	rl.Mut.Unlock()
	return rooms
}

// AddRoom adds a room to the global room list
func (rl *RoomList) AddRoom(room *models.Room) {
	rl.Mut.Lock()
	rl.Rooms[room.GetID()] = room
	rl.Mut.Unlock()
}

// RemoveRoom removes a room from the global room list
func (rl *RoomList) RemoveRoom(room models.Room) {
	rl.Mut.Lock()
	delete(rl.Rooms, room.GetID())
	rl.Mut.Unlock()
}

// GetRoom returns a room from the global room list
func (rl *RoomList) GetRoom(rID string) models.Room {
	rl.Mut.Lock()
	r := *rl.Rooms[rID]
	rl.Mut.Unlock()
	return r
}

func (rl *RoomList) AddPlayer(rID, pass string, player *models.BasePlayer) error {
	rl.Mut.Lock()
	r := rl.Rooms[rID]
	err := r.AddPlayer(player, pass)
	if err != nil {
		rl.Mut.Unlock()
		return err
	}
	rl.Rooms[rID] = r
	rl.Mut.Unlock()
	return nil
}
