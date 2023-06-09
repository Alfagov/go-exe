package bowling

import (
	"errors"
)

type Bowling struct {
	rolls        []int
	pins         int
	currentRoll  int
	currentFrame int
	ended        bool
}

// Calculate the score of a game of bowling.
func (b *Bowling) Score() (int, error) {
	if !b.ended {
		return 0, errors.New("game is not ended")
	}
	return score(b.rolls, 0), nil
}

// Calculate the score given a list of rolls and a previous score.
func score(rolls []int, prev int) int {
	if len(rolls) == 3 {
		return sum(rolls...)
	}
	roll, rolls := rolls[0], rolls[1:]
	if roll == 10 {
		return roll + sum(rolls[0:2]...) + score(rolls, 0)
	} else if roll+prev == 10 {
		return roll + rolls[0] + score(rolls, 0)
	} else if prev == 0 {
		return roll + score(rolls, roll)
	}
	return roll + score(rolls, 0)
}

// is called each time the player rolls a ball. The argument is the number of pins knocked down
func (b *Bowling) Roll(pins int) error {

	// check game is still valid
	switch {
	case b.ended:
		return errors.New("game is ended")
	case pins < 0 || pins > 10:
		return errors.New("pins is out of range")
	case pins > b.pins:
		return errors.New("too many pins")
	}

	// add roll to game
	b.rolls = append(b.rolls, pins)
	// update game state pins
	b.pins -= pins

	// check if endame
	if b.currentFrame == 10 &&
		((b.currentRoll == 2 && sum(b.rolls[len(b.rolls)-2:]...) < 10) ||
			b.currentRoll == 3) {
		b.ended = true
	} else {
		// Continue frame
		if b.currentFrame == 10 {
			b.currentRoll++
			if b.pins == 0 {
				b.pins = 10
			}
		} else {
			// strike or spare or second roll of the frame
			if pins == 10 || b.pins == 0 || b.currentRoll == 2 {
				b.currentFrame++
				b.currentRoll = 1
				b.pins = 10
			} else {
				b.currentRoll = 2
			}
		}
	}
	return nil
}

// Create a new game of bowling
func NewGame() *Bowling {
	return &Bowling{nil, 10, 1, 1, false}
}
