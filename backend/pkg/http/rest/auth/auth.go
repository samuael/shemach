package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/session"
	"github.com/samuael/shemach/backend/pkg/user"
)

var (
	ErrAuthenticationTokenNotFound  = errors.New("authentication token not found")
	ErrAuthorizationSessionNotFound = errors.New("authorization session not found")
	ErrInvalidSessionToken          = errors.New("invalid session token")
	ErrSessionNotFound              = errors.New("session not found")
)

// Authenticator representing the Methods to be implemented by the authenticators
type Authenticator interface {
	SaveRegistrationSession(writer http.ResponseWriter, session *types.TempoRegistrationSession) (string, bool)
	SaveEmailConfirmationSession(session *types.EmailConfirmationSession) (string, bool)
	SaveSession(writer http.ResponseWriter, session *types.Session) (string, error)
	// SaveSubscriberSession(writer http.ResponseWriter, session *types.SubscriberSession) bool
	DeleteSession(writer http.ResponseWriter, request *http.Request) bool

	GetSession(request *http.Request) (*types.Session, error)
	GetEmailSession(token string) (*types.EmailConfirmationSession, error)

	RandomToken() string
	ValidateToken(tokenstring string) bool

	LogoutSession(request *http.Request) error
	GetTempoRegistrationSession(request *http.Request) (*types.TempoRegistrationSession, error)
}

// authenticator representing the Cookie methods and handler in jwt
type authenticator struct {
	SessionService session.ISessionService
}

// NewCookieHandler representing New Cookie thing
func NewAuthenticator(
	sessionService session.ISessionService,
) Authenticator {
	return &authenticator{
		SessionService: sessionService,
	}
}

// SaveRegistrationSession to save the Session in the User Session Header
func (sessh *authenticator) SaveRegistrationSession(writer http.ResponseWriter, session *types.TempoRegistrationSession) (string, bool) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(12 * time.Hour)
	session.StandardClaims = jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		return "", false
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
	return tokenString, true
}

// SaveSession to save the Session in the User Session Header
func (sessh *authenticator) SaveSession(writer http.ResponseWriter, session *types.Session) (authToken string, err error) {
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
	authToken, err = token.SignedString([]byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		return authToken, err
	}
	err = sessh.SessionService.SaveSession(context.Background(), &types.SaveSession{
		UserID: session.ID,
		Token:  authToken,
	})
	if err != nil {
		return authToken, err
	}
	// Setting the bearer Authorization Token to  the header
	writer.Header().Set("Authorization", "Bearer "+authToken)
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	cookie := http.Cookie{
		Name:     os.Getenv("COOKIE_NAME"),
		Value:    authToken,
		Expires:  expirationTime,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(writer, &cookie)
	return authToken, nil
}

// SaveEmailConfirmationSession to save the Session in the User Session Header
func (sessh *authenticator) SaveEmailConfirmationSession(session *types.EmailConfirmationSession) (string, bool) {
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
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_EMAIL_CONFIRMATION_KEY")))
	if err != nil {
		return tokenString, false
	}
	return tokenString, true
}

// DeleteSession representing del
func (sessh *authenticator) DeleteSession(writer http.ResponseWriter, request *http.Request) bool {
	session := types.Session{}
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

// GetTempoRegistrationSession returns a session struct from user registration authorization token
func (sessh *authenticator) GetTempoRegistrationSession(request *http.Request) (*types.TempoRegistrationSession, error) {
	var tknStr string
	cookie, err := request.Cookie(os.Getenv("COOKIE_NAME"))
	defer recover()
	if err != nil {
		// go and check for the authorization header
		tknStr = request.Header.Get("Authorization")
		tknStr = strings.Trim(strings.TrimPrefix(tknStr, "Bearer "), " ")
		if tknStr == "" {
			return nil, user.ErrMissingAuthorizationHeader
		}
	} else {
		tknStr = cookie.Value
	}

	session := &types.TempoRegistrationSession{}
	tkn, err := jwt.ParseWithClaims(tknStr, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if tkn.Valid {
		return session, nil
	}
	return nil, user.ErrInvalidSession
}

// GetSession returns a Session Struct Having the Data of the User
func (sessh *authenticator) GetSession(request *http.Request) (*types.Session, error) {
	cookie, err := request.Cookie(os.Getenv("COOKIE_NAME"))
	defer recover()
	var tokenString string
	if err != nil {
		// go and check for the authorization header
		var ok bool
		tokenString = request.Header.Get("Authorization")
		tokenString = strings.Trim(strings.TrimPrefix(tokenString, "Bearer "), " ")
		if tokenString == "" {
			_, tokenString, ok = request.BasicAuth()
			if tokenString == "" || !ok {
				return nil, ErrAuthenticationTokenNotFound
			}
		}
	} else {
		tokenString = cookie.Value
	}
	session := &types.Session{}
	tkn, err := jwt.ParseWithClaims(tokenString, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_SECRET_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if tkn.Valid {
		var ssession *types.SaveSession
		ssession, err = sessh.SessionService.GetSessionByUserID(context.Background(), uint(session.ID))
		if ssession == nil {
			return nil, ErrAuthorizationSessionNotFound
		} else if err != nil {
			return nil, err
		}
		return session, nil
	}
	return nil, ErrInvalidSessionToken
}

// GetEmailSession(request *http.Request) (*types.Session, error)
func (sessh *authenticator) GetEmailSession(token string) (*types.EmailConfirmationSession, error) {
	// go and check for the authorization header
	if token == "" {
		return nil, nil
	}
	session := &types.EmailConfirmationSession{}
	tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SESSION_EMAIL_CONFIRMATION_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if tkn.Valid {
		return session, nil
	}
	return nil, errors.New("invalid login session")
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

// LogoutSession deletes a session instance from database
func (sessh *authenticator) LogoutSession(request *http.Request) error {
	cookie, err := request.Cookie(os.Getenv("COOKIE_NAME"))
	defer recover()
	if err != nil {
		// go and check for the authorization header
		// var username string
		var ok bool
		token := request.Header.Get("Authorization")
		token = strings.Trim(strings.TrimPrefix(token, "Bearer "), " ")
		if token == "" {
			_, token, ok = request.BasicAuth()
			if token == "" || !ok {
				return errors.New("not authenticated")
			}
		}
		session := &types.Session{}
		tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")), nil
		})
		if err != nil {
			return err
		}
		if tkn.Valid {
			er := sessh.SessionService.DeleteSesssion(context.Background(), 0, uint(session.ID))
			if er != nil {
				return er
			}
			return nil
		}
		return errors.New(" invalid login session ")
	}
	tknStr := cookie.Value
	session := &types.Session{}
	tkn, err := jwt.ParseWithClaims(tknStr, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return err
		}
		return err
	}
	if tkn.Valid {
		er := sessh.SessionService.DeleteSesssion(context.Background(), 0, uint(session.ID))
		if er != nil {
			return er
		}
		return nil
	}
	return errors.New(" invalid login session ")
}
