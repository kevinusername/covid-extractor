package state

import "time"

type Record struct {
	State     string
	Confirmed int
	Deaths    int
	Updated   time.Time
}
