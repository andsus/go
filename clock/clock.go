package clock

// import (
// 	"fmt"
// )

// // Clock is type
// type Clock struct {
// 	m int
// }

// // New function to create Clock
// func New(hour, minute int) (c Clock) {
// 	m := hour*60 + minute
// 	m %= 24 * 60
// 	if m < 0 {
// 		m += 24 * 60
// 	}
// 	return Clock{m}
// }

// // Add function add to Clock, can be negative
// func (c Clock) Add(minutes int) Clock {
// 	return New(0, c.m+minutes)
// }

// // Subtract function add to Clock, can be negative
// func (c Clock) Subtract(minutes int) Clock {
// 	return c.Add(-minutes)
// }

// // String prints the time a Clock represents
// func (c Clock) String() string {
// 	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
// }

import "fmt"

const minutesInAnHour = 60
const hoursInADay = 24
const minutesInADay = minutesInAnHour * hoursInADay

// Clock that handles time (hours and minutes) without dates
type Clock struct {
	minutes int
}

// New creates a new clock
func New(hours int, minutes int) Clock {
	minutes = hours*minutesInAnHour + minutes
	wrappedMinutes := (minutes%minutesInADay + minutesInADay) % minutesInADay
	return Clock{wrappedMinutes}
}

// func new(minutes int) Clock {
// 	wrappedMinutes := (minutes%minutesInADay + minutesInADay) % minutesInADay
// 	return Clock{wrappedMinutes}
// }

func (c Clock) String() string {
	hours, minutes := divide(c.minutes, minutesInAnHour)
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

type quotient = int
type remainder = int

func divide(dividend int, divisor int) (quotient, remainder) {
	return dividend / divisor, dividend % divisor
}

// Add a given amount of minutes to this Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

// Subtract a given amount of minutes from this Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minutes-minutes)
}
