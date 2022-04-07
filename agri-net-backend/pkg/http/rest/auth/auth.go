package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/model"
	// "github.com/samuael/Project/RegistrationSystem/platforms/helper"
)

// Authenticator representing the Methods to be implemented by the authenticators
type Authenticator interface {
	SaveSession(writer http.ResponseWriter, session *model.Session) bool
	DeleteSession(writer http.ResponseWriter, request *http.Request) bool
	GetSession(request *http.Request) (*model.Session, error)
	RandomToken() string
	ValidateToken(tokenstring string) bool
}

// authenticator representing the Cookie methods and handler in jwt
type authenticator struct {
}

// NewCookieHandler representing New Cookie thing
func NewAuthenticator() Authenticator {
	return &authenticator{}
}

// SaveSession to save the Session in the User Session Header
func (sessh *authenticator) SaveSession(writer http.ResponseWriter, session *model.Session) bool {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)
	session.StandardClaims = jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
		// HttpOnly:  true,
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		return false
	}
	// Setting the bearer Authorization Token to  the header
	writer.Header().Set("Authorization", "Bearer "+tokenString)
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	cookie := http.Cookie{
		Name:     os.Getenv("COOKIE_NAME"),
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(writer, &cookie)
	return true
}

// DeleteSession representing del
func (sessh *authenticator) DeleteSession(writer http.ResponseWriter, request *http.Request) bool {
	session := model.Session{}
	expirationTime := time.Now().Add(-2400 * time.Hour)
	session.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		return false
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	cookie := http.Cookie{
		Name:    os.Getenv("COOKIE_NAME"),
		Value:   tokenString,
		Expires: expirationTime,
		// Domain:   host,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(writer, &cookie)
	return true
}

// GetSession returns a Session Struct Having the Data of the User
func (sessh *authenticator) GetSession(request *http.Request) (*model.Session, error) {
	cookie, err := request.Cookie(os.Getenv("COOKIE_NAME"))
	defer recover()
	if err != nil {
		// go and check for the authorization header
		token := request.Header.Get("Authorization")
		token = strings.Trim(strings.TrimPrefix(token, "Bearer "), " ")
		if token == "" {
			return nil, nil
		}
		session := &model.Session{}
		tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SESSION_SECRET_KEY")), nil
		})
		if err != nil {
			return nil, err
		}
		if tkn.Valid {
			return session, nil
		}
		return nil, errors.New(" invalid login session ")
	}
	tknStr := cookie.Value
	session := &model.Session{}
	tkn, err := jwt.ParseWithClaims(tknStr, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	return session, nil
}

// RandomToken random token Generator for CSRF and related technologies
func (sessh *authenticator) RandomToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(os.Getenv("CAR_INSPECTION_COOKIE_NAME"))
	return tokenString
}

// ValidateToken representing the Form Value
func (sessh *authenticator) ValidateToken(tokenstring string) bool {
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}
