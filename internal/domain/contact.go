package domain

import "time"

type Contact struct {
	ID        uint
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
