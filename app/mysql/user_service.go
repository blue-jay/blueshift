package mysql

import (
	"database/sql"
	"fmt"

	app "github.com/blue-jay/blueshift"

	"github.com/jmoiron/sqlx"
)

const (
	// table is the table name.
	table = "user"
)

// UserService represents a PostgreSQL implementation of myapp.UserService.
type UserService struct {
	DB *sqlx.DB
}

// UserByEmail returns a user for an email.
func (s *UserService) UserByEmail(email string) (*app.User, bool, error) {
	r := &app.User{}
	err := s.DB.Get(&r, fmt.Sprintf(`
		SELECT id, password, status_id, first_name
		FROM %v
		WHERE email = ?
			AND deleted_at IS NULL
		LIMIT 1
		`, table),
		email)
	return r, err == sql.ErrNoRows, err
}

// CreateUser creates a user.
func (s *UserService) CreateUser(u *app.User) error {
	_, err := s.DB.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(first_name, last_name, email, password)
		VALUES
		(?,?,?,?)
		`, table),
		u.FirstName, u.LastName, u.Email, u.Password)
	return err
}
