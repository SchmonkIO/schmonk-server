package logic

import (
	"math/rand"
	"time"
)

/*

Possible ground tiles:
- path: normal path to walk on
- grass: normal ground, not walkable
- shop_path: path tile where you can enter the shop
- star_path: path tile where you get a star for free
- intersection_path: path tile where you can choose a direction

Generate Map in multiple rounds
- 1. Set the normal path down
- 2. Fill other space with grass
- 3. Replace some normal path tiles with special ones


{
	"row1": [
		"column1": "path",
		"column2": "...",
		...
	],
	"row2": [
		"column1": "...",
		"column2": "...",
		...
	],
	...
}
*/

func GenWorld() {
	height := 32
	width := 32
	startX, startY := GenStartPoint(height, width)
	spaceTillLeftWall := startX - 1
	leftEnd
}

func GenStartPoint(height, width int) (int, int) {
	y := generateRandomNumber(height / 2)
	y = y + (height / 2)
	x := generateRandomNumber(width / 2)
	x = x + (width / 2)
	return x, y
}

func generateRandomNumber(maxValue int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(maxValue + 1)
}
