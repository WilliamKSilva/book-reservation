package domain

import "time"

type Book struct {
	ID            string
	Title         string
	Author        string
	Genre         string
	PublishedDate time.Time
}
