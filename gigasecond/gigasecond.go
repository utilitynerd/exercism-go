// Package gigasecond provides fucntions to work with gigaseconds
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = 1000000000

// AddGigasecond returns the time 1,000,000,000 seconds later
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond * time.Second)
}
