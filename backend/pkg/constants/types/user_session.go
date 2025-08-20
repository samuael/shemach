package types

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
	TempoRegistrationSession struct {
		jwt.StandardClaims
		ID    int64
		Phone string
	}

	Session struct {
		jwt.StandardClaims
		ID   uint64
		Role uint8
	}

	// ForgotPasswordSession
	ForgotPasswordSession struct {
		jwt.StandardClaims
		ID    uint64
		Email string
		Time  time.Time
	}
	// EmailConfirmationSession
	EmailConfirmationSession struct {
		jwt.StandardClaims
		*EmailConfirmation
	}
)
