package clock

import (
	"fmt"
)

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	ret := Clock{hour: h, minute: m}
	ret.timeRule()
	return ret
}

func (clock *Clock) timeRule() {
	addHour := (clock.minute / 60)
	clock.hour += addHour
	clock.minute = clock.minute % 60
	if clock.minute < 0 {
		clock.hour--
		clock.minute = 60 + clock.minute
	}
	clock.hour = clock.hour % 24
	if clock.hour < 0 {
		clock.hour = 24 + clock.hour
	}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(n int) Clock {
	clock.minute += n
	clock.timeRule()
	return clock
}

func (clock Clock) Subtract(n int) Clock {
	clock.minute -= n
	clock.timeRule()
	return clock
}

/*
package clock

import "fmt"

type Clock int

const minInHour = 60
const minInDay = 24 * 60

//New returns a new Clock.
func New(h int, m int) Clock {
	c := (h*minInHour + m) % minInDay

	if c < 0 {
		c += minInDay
	}

	return Clock(c)
}

//String returns a string representation of the Clock.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/minInHour, c%minInHour)
}

//Add returns a new clock adding the minutes.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

//Subtract returns a new clock subtracting the minutes.
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}
*/
