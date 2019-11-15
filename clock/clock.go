package clock

import "fmt"

// Clock implements a 24 hour clock, with minute resolution and no time zone
type Clock struct {
	minutes int
}

// New returns a new Clock.
func New(hours, minutes int) Clock {
	minutes += hours * 60
	minutes %= 24*60
	if minutes < 0 {
		minutes += 24*60
	}
	return Clock{minutes}
}

// String returns a string representation of a Clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add returns a new Clock with minutes added
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract returns a new Clock with minutes Subtracted
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
