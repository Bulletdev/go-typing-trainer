package anticheat

import "time"

type AntiCheat struct {
	MaxAPM float64
}

func NewAntiCheat() *AntiCheat {
	return &AntiCheat{
		MaxAPM: 300, // Adjust this value based on realistic typing speeds
	}
}

func (ac *AntiCheat) DetectCheating(keyPresses int, duration time.Duration) bool {
	if duration.Minutes() == 0 {
		return false
	}
	apm := float64(keyPresses) / duration.Minutes()
	return apm > ac.MaxAPM
}
