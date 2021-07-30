package structure

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	// ID        string    `json:"id" sql:"id"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Username  string    `json:"username" validate:"min=3"`
	TokenHash string    `json:"tokenhash"`
	RefToken  string    `json:"reftoken"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
	// UserID    	string    `json:"userID"`
}

// type UserKey struct{}
type SignedDetails struct {
	Email    string
	Username string
	jwt.StandardClaims
}
