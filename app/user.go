package app

import "github.com/go-sql-driver/mysql"

// User defines a person.
type User struct {
	ID        uint32         `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	StatusID  uint8          `db:"status_id"`
	CreatedAt mysql.NullTime `db:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at"`
	DeletedAt mysql.NullTime `db:"deleted_at"`
}

// UserService
type UserService interface {
	UserByEmail(email string) (*User, error)
	CreateUser(u *User) error
}
