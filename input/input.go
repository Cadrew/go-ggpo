// This file is an example on how we could implement a link between the inputs of a game/emulator and GGPO
package input

import (
	"github.com/Cadrew/go-ggpo/ggpo/ggponet"
)

// MaxPlayers is the maximum number of players to poll input for
const MaxPlayers = ggponet.GGPO_MAX_PLAYERS

type inputstate [MaxPlayers][ActionLast]bool

// Input state for all the players
var (
	NewState inputstate // input state for the current frame
	OldState inputstate // input state for the previous frame
	Released inputstate // keys just released during this frame
	Pressed  inputstate // keys just pressed during this frame
)

// Hot keys
const (
	// ActionLast is used for iterating
	ActionLast uint32 = 15 // this is an example, it represents the number of inputs on a controller/keyboard
)

// Resets all buttons to false
func reset(state inputstate) inputstate {
	for p := range state {
		for k := range state[p] {
			state[p][k] = false
		}
	}
	return state
}

func keyboardToPlayer(st inputstate, p int) inputstate {
	// get inputs from keyboard
	return st
}

// pollKeyboard processes keyboard keys
func PollKeyboard(st inputstate) inputstate {
	st = keyboardToPlayer(st, 0)
	return st
}

// Compute the keys pressed or released during this frame
func getPressedReleased(new inputstate, old inputstate) (inputstate, inputstate) {
	for p := range new {
		for k := range new[p] {
			Pressed[p][k] = new[p][k] && !old[p][k]
			Released[p][k] = !new[p][k] && old[p][k]
		}
	}
	return Pressed, Released
}

// Poll calculates the input state. It is meant to be called for each frame.
func Poll() {
	Pressed, Released = getPressedReleased(NewState, OldState)
	OldState = NewState
}

func Reset() {
	NewState = reset(NewState)
}

// State is a callback example
// It returns 1 if the button corresponding to the parameters is pressed
func State(port uint, index uint, id uint) int16 {
	if id >= 255 || index > 0 {
		return 0
	}

	if NewState[port][id] {
		return 1
	}
	return 0
}
