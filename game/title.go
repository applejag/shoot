package game

import "github.com/firefly-zero/firefly-go/firefly"

type Title struct {
	msg string
	ttl int
}

func (t *Title) show(msg string) {
	*t = Title{
		msg: msg,
		ttl: 180,
	}
}

func (t *Title) hide() {
	t.ttl = 0
}

func (t Title) shouldShow() bool {
	return t.ttl > 0
}

func (t *Title) update() {
	t.ttl--
	if t.ttl <= 0 {
		openHub()
	}
}

func (t Title) render() {
	x := (firefly.Width - font.LineWidth(t.msg)) / 2
	firefly.DrawText(t.msg, font, firefly.P(x, 80), firefly.ColorBlack)
}
