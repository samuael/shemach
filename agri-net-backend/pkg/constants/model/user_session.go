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
		ID        uint64
		Firstname string
		Lastname  string
		Email     string
		Role      string
		Lang      string
	}
	// SubscriberSession
	SubscriberSession struct {
		jwt.StandardClaims
		ID       uint64
		Phone    string
		Fullname string
		Lang     string
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
	// EmailConfirmation
	EmailConfirmation struct {
		ID           string `json:"id"`
		UserID       string `json:"user_id"`
		Email        string `json:"email"`
		OldEmail     string `json:"old_email"`
		CreatedAt    uint64 `json:"created_at"`
		IsNewAccount bool   `json:"is_new_account"`
	}
)
