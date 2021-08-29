package clock

import (
	"fmt"
)

// Clock is type
type Clock struct {
	m int
}

// New function to create Clock
func New(hour, minute int) (c Clock) {
	m := hour*60 + minute
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

// Add function add to Clock, can be negative
func (c Clock) Add(minutes int) Clock {
	return New(0, c.m+minutes)
}

// Subtract function add to Clock, can be negative
func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-minutes)
}

// String prints the time a Clock represents
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
