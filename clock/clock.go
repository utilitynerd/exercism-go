package clock

import "fmt"

// Clock implements a 24 hour clock, with no time zone
type Clock struct {
	h int
	m int
}

// New returns a new Clock.
func New(hours, minutes int) Clock {
	if minutes >= 60 {
		hours += minutes / 60
		minutes %= 60
	}
	if minutes < 0 {
		if minutes%60 == 0 {
			hours += minutes / 60
			minutes = 0
		} else {
			hours += (minutes / 60) - 1
			minutes = 60 + minutes%60
		}
	}
	if hours < 0 {
		hours = 24 + hours%24
	}
	return Clock{hours % 24, minutes}
}

// String returns a string representation of a Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add returns a new Clock with minutes added
func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}

// Subtract returns a new Clock with minutes Subtracted
func (c Clock) Subtract(minutes int) Clock {
	return New(c.h, c.m-minutes)
}
