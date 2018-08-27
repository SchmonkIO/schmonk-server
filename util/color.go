package util

import "math/rand"

const letterBytes = "abcdef0123456789"

// GetRandomColor generates a random hex color and returns the string
func GetRandomColor() string {
	return "#" + randStringBytes(6)
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
