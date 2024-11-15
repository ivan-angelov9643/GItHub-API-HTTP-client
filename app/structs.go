package app

import "time"

type User struct {
	Login     string `json:"login"`
	Followers int    `json:"followers"`
}

type Repository struct {
	Name      string    `json:"name"`
	Forks     int       `json:"forks_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
