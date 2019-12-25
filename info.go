package toc

import "time"

// Info describe table of contents, words & chars counters.
type Info struct {
	Headers      []Header
	Words, Chars int
}

// Duration will return the approximate text reading time at a given speed
// reading (characters per minute).
func (info Info) Duration(speed int) time.Duration {
	return (time.Duration(info.Chars) * time.Minute / time.Duration(speed)).
		Round(time.Second)
}
