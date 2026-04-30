package game

import (
	"github.com/firefly-zero/firefly-go/firefly"
)

type Score struct {
	// How many enemies the players are yet to kill.
	val int
}

func newScore() Score {
	return Score{val: 30}
}

func (s *Score) decrement() {
	s.val--
	if s.val == 0 {
		if iAmAlive() {
			title.show("victory!")
		} else {
			title.show("victory but without you")
		}
	}
}

// Decrease the score to the given value.
func (s *Score) decreaseTo(v int) {
	if s.val > v {
		s.val = v
	} else {
		s.decrement()
	}
}

func (s Score) render() {
	if hub {
		return
	}
	t := formatScore(s.val)
	p := firefly.P(firefly.Width/2-font.CharWidth(), 10)
	font.DrawBytes(t[:], p, firefly.ColorBlack)
}

func formatScore(i int) [2]byte {
	return [...]byte{'0' + byte(i/10), '0' + byte(i%10)}
}
