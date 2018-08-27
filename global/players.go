package global

import (
	"sync"

	"github.com/schmonk.io/schmonk-server/models"
)

// Players is a list of all connected players
var Players PlayerList{}

// PlayerList is a struct for a list of BasePlayers
type PlayerList struct {
	Players map[string]models.BasePlayer
	Mut     *sync.Mutex
}

// AddPlayer adds a new player to the global player list
func (pl *PlayerList) AddPlayer(player models.BasePlayer) {
	pl.Mut.Lock()
	_, ok := m[player.GetID()]
	if !ok {
		pl.Players[player.GetID()] = player
	}
	pl.Mut.Unlock()
}

// RemovePlayer removes a player from the global player list
func (pl *PlayerList) RemovePlayer(player models.BasePlayer) {
	pl.Mut.Lock()
	delete(pl.Players, player.GetID())
	pl.Mut.Unlock()
}

// GetPlayer returns a player from the global player list
func (pl *PlayerList) GetPlayer(pID string) models.BasePlayer {
	pl.Mut.Lock()
	bp := pl.Players[pID]
	pl.Mut.Unlock()
	return bp
}

// ModifyPlayer modifies an existing player in the global player list
func (pl *PlayerList) ModifyPlayer(player models.BasePlayer) {
	pl.Mut.Lock()
	pl.Players[player.GetID()] = player
	pl.Mut.Unlock()
}
