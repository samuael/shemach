package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Session representing the Session to Be sent with the request body
// no saving of a session in the database so i Will use this session in place of
type (
	Session struct {
		jwt.StandardClaims
		ID       uint64
		Email    string
		Password string
		Role     string
	}

	ForgotPasswordSession struct {
		jwt.StandardClaims
		ID    uint64
		Email string
		Time  time.Time
	}
)
