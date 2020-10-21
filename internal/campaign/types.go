package campaign

import "time"

type Click struct {
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Timestamp   time.Time `json:"timestamp"`
}
