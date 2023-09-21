package domain

import "time"

type Order struct {
	ID        int64
	Total     int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
