package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Session representing the Session to Be sent with the request body
// no saving of a session in the database so i Will use this session in place of
const (
	SUPERADMIN = iota
)

type (
	Session struct {
		jwt.StandardClaims
		ID    uint64
		Email string
		Role  string
	}

	SubscriberSession struct {
		jwt.StandardClaims
		ID       uint64
		Phone    string
		Fullname string
	}

	ForgotPasswordSession struct {
		jwt.StandardClaims
		ID    uint64
		Email string
		Time  time.Time
	}
)
