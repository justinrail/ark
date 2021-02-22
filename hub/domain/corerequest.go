package domain

import "time"

//CoreRequest request message
type CoreRequest struct {
	Destination string
	Target      string
	Category    int
	birthTime   time.Time
	Content     interface{}
}
