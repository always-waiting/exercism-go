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
