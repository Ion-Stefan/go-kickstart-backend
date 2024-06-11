package types

import "time"

type UserStore interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
}

type User struct {
	CreatedAt time.Time `json:"createdAt"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ID        int       `json:"id"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6,max=127"`
}
