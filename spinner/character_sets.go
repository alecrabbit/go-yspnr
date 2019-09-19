package spinner

import (
	"time"
)

const (
	clockOneOClock = '\U0001F550'
	clockOneThirty = '\U0001F55C'
)

// Declare spinner types
const (
	Arrows int = iota
	Blocks
	BouncingBlock
	RotatingCircle
	Clock
	HalfClock
)

// CharSets contains the available character sets
var CharSets = map[int][]string{
	Arrows:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	Blocks:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	BouncingBlock:  {"▖", "▘", "▝", "▗"},
	RotatingCircle:  {"◐", "◓", "◑", "◒"},
}
type S struct {
	CharSet 	[]string
	Interval 	time.Duration
}
// Settings contains the spinner type specific settings
var SettingsSets = map[int][]string{
	Arrows:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	Blocks:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	BouncingBlock:  {"▖", "▘", "▝", "▗"},
	RotatingCircle:  {"◐", "◓", "◑", "◒"},
}

func init() {
	for i := rune(0); i < 12; i++ {
		CharSets[Clock] = append(CharSets[Clock], string([]rune{clockOneOClock + i}))
		CharSets[HalfClock] = append(CharSets[HalfClock], string([]rune{clockOneOClock + i}), string([]rune{clockOneThirty + i}))
	}
	// fmt.Println(CharSets[Clock])
	// fmt.Println(CharSets[HalfClock])
}
