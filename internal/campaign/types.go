package campaign

import "time"

type Click struct {
	Origin      string
	Destination string
	CreatedAt   time.Time
}
