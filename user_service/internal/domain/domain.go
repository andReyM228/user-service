package domain

import "time"

type User struct {
	ID        int
	Name      string
	Surname   string
	Phone     string
	Email     string
	CreatedAt time.Time `db:"created_at"`
}

type Car struct {
	ID        int
	Name      string
	Model     string
	CreatedAt time.Time `db:"created_at"`
}
