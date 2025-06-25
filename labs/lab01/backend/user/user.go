package user

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrInvalidEmail is returned when the email format is invalid
	ErrInvalidEmail = errors.New("invalid email format")
	// ErrInvalidAge is returned when the age is invalid
	ErrInvalidAge = errors.New("invalid age: must be between 0 and 150")
	// ErrEmptyName is returned when the name is empty
	ErrEmptyName = errors.New("name cannot be empty")
)

// User represents a user in the system
type User struct {
	Name  string
	Age   int
	Email string
}

// NewUser creates a new user with validation
func NewUser(name string, age int, email string) (*User, error) {
	user := &User{Name: name, Age: age, Email: email}
	err := user.Validate()
	if err == nil {
		return user, nil
	}
	return nil, err
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
	if !IsValidEmail(u.Email) {
		return ErrInvalidEmail
	} else if u.Age < 0 || u.Age > 150 {
		return ErrInvalidAge
	} else if u.Name == "" {
		return ErrEmptyName
	}
	return nil
}

// String returns a string representation of the user
func (u *User) String() string {

	return u.Name + "/n" + strconv.FormatInt(int64(u.Age), 10) + "/n" + u.Email
}

// IsValidEmail checks if the email format is valid
func IsValidEmail(email string) bool {
	if strings.Contains(email, "@") {
		domen := strings.Split(email, "@")[1]
		if strings.Contains(domen, ".") {
			return true
		}
	}
	return false
}
