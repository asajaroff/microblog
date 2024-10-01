package models

import (
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	IsAdmin    bool      `json:"is_admin"`
	IsActive   bool      `json:"is_active"`
	IsDisabled bool      `json:"is_disabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Blog struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	OwnerID    uint      `json"owner_id"`
	IsActive   bool      `json:"is_active"`
	IsDisabled bool      `json:"is_disabled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Post struct {
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
