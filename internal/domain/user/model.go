package user

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// validate ensure the data is valid
func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}
	if u.Email == "" {
		return errors.New("Email is required")
	}
	return nil
}
