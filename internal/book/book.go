package internal

import "time"

type Post struct {
	ID            string
	Title         string
	Author        string
	Genre         string
	PublishedDate time.Time
}
