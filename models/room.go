package models

import (
	"sync"

	"github.com/schmonk.io/schmonk-server/config"
	"github.com/schmonk.io/schmonk-server/util"
	"gopkg.in/mgo.v2/bson"
)

// Room is struct for a room
type Room struct {
	ID      bson.ObjectId         `json:"_id"`
	Name    string                `json:"name"`
	Pass    string                `json:"pass"`
	Map     string                `json:"map"`
	Slots   int                   `json:"slots"`
	Owner   string                `json:"owner"`
	Players map[string]RoomPlayer `json:"players"`
	Mut     *sync.Mutex           `json:"-"`
}

// CreateRoom creates and returns a new room
func CreateRoom(name, pass, owner string, slots int) Room {
	r := Room{}
	r.SetID()
	r.SetName(name)
	r.SetPass(pass)
	r.SetOwner(owner)
	r.SetSlots(slots)
	return r
}

// GetID returns the room id as a string
func (r *Room) GetID() string {
	return r.ID.Hex()
}

// SetID sets a random id for the room
func (r *Room) SetID() {
	r.ID = bson.NewObjectId()
}

// SetName sets the name for the room
func (r *Room) SetName(name string) error {
	if len(name) <= config.Config.Game.NameLength {
		r.Name = name
		return nil
	} else {
		return util.ErrNameToLong
	}
}

// SetPass sets the password for the room
func (r *Room) SetPass(pass string) {
	r.Pass = pass
}

// SetOwner sets the owner id of the room
func (r *Room) SetOwner(oID string) {
	r.Owner = oID
}

// SetSlots sets the number of slots available for the room
func (r *Room) SetSlots(quantity int) {
	if quantity <= config.Config.Game.SlotsPerRoom {
		r.Slots = quantity
	} else {
		r.Slots = config.Config.Game.SlotsPerRoom
	}
}

// AddPlayer adds a new player to the room
func (r *Room) AddPlayer(player RoomPlayer, pass string) error {
	r.Mut.Lock()
	if len(r.Players) < r.Slots {
		if r.Pass != "" {
			if r.Pass != pass {
				return util.ErrPassWrong
			}
		}
		r.Players[player.GetID()] = player
		r.Mut.Unlock()
		return nil
	} else {
		r.Mut.Unlock()
		return util.ErrNoSlots
	}
}

// GetPlayerCount returns the number of players in the room
func (r *Room) GetPlayerCount() int {
	r.Mut.Lock()
	count := len(r.Players)
	r.Mut.Unlock()
	return count
}

// RemovePlayer removes the player from the room
func (r *Room) RemovePlayer(player RoomPlayer) error {
	r.Mut.Lock()
	if player.GetID() == r.Owner {
		delete(r.Players, player.GetID())
		for _, p := range r.Players {
			r.Owner = p.GetID()
			break
		}
		return nil
	}
	delete(r.Players, player.GetID())
	if len(r.Players) <= 0 {
		return util.ErrNoPlayer
	}
	r.Mut.Unlock()
	return nil
}
