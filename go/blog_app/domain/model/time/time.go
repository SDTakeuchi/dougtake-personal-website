package time

import (
	"time"
)

type Time struct {
	time.Time
}

func (t Time) StringDay() string {
	return t.Format("Jan 02, 2006")
}

func (t Time) StringHour() string {
	return t.Format("Jan 02, 2006 15:04")
}
