package global

import (
	"sync"

	"github.com/schmonk.io/schmonk-server/models"
)

// Rooms is a list of all available rooms
var Rooms RoomList

// RoomList is a struct for a list of Rooms
type RoomList struct {
	Rooms map[string]*models.Room
	Mut   *sync.Mutex
}

func CreateGlobalRoomList() {
	grList := RoomList{}
	grList.Rooms = map[string]*models.Room{}
	grList.Mut = &sync.Mutex{}
	Rooms = grList
}

// AddRoom adds a room to the global room list
func (rl *RoomList) AddRoom(room *models.Room) {
	rl.Mut.Lock()
	rl.Rooms[room.GetID()] = room
	rl.Mut.Unlock()
}

// RemoveRoom removes a room from the global room list
func (rl RoomList) RemoveRoom(room models.Room) {
	rl.Mut.Lock()
	delete(rl.Rooms, room.GetID())
	rl.Mut.Unlock()
}

// GetRoom returns a room from the global room list
func (rl RoomList) GetRoom(rID string) *models.Room {
	rl.Mut.Lock()
	r := rl.Rooms[rID]
	rl.Mut.Unlock()
	return r
}
