package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/session"
	// "github.com/samuael/Project/RegistrationSystem/platforms/helper"
)

// Authenticator representing the Methods to be implemented by the authenticators
type Authenticator interface {
	SaveSubscriberSession(writer http.ResponseWriter, session *model.SubscriberSession) bool
	SaveEmailConfirmationSession(session *model.EmailConfirmationSession) (string, bool)
	SaveSession(writer http.ResponseWriter, session *model.Session) bool
	DeleteSession(writer http.ResponseWriter, request *http.Request) bool

	GetSession(request *http.Request) (*model.Session, error)
	GetSubscriberSession(request *http.Request) (*model.SubscriberSession, error)
	GetEmailSession(token string) (*model.EmailConfirmationSession, error)

	RandomToken() string
	ValidateToken(tokenstring string) bool

	LogoutSubscriberSession(request *http.Request) error
	LogoutSession(request *http.Request) error
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
	er := sessh.SessionService.SaveSession(context.Background(), &model.SaveSession{
		UserID: int(session.ID),
		Token:  tokenString,
	})
	if er != nil {
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

// SaveSubscriberSession to save the Session in the User Session Header
func (sessh *authenticator) SaveSubscriberSession(writer http.ResponseWriter, session *model.SubscriberSession) bool {
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
	tokenString, err := token.SignedString([]byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")))
	if err != nil {
		return false
	}
	er := sessh.SessionService.SaveSubscriberSession(context.Background(), &model.SaveSubscriberSession{
		SubscriberID: int(session.ID),
		Token:        tokenString,
	})
	if er != nil {
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

// SaveEmailConfirmationSession to save the Session in the User Session Header
func (sessh *authenticator) SaveEmailConfirmationSession(session *model.EmailConfirmationSession) (string, bool) {
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
		var username string
		var ok bool
		token := request.Header.Get("Authorization")
		token = strings.Trim(strings.TrimPrefix(token, "Bearer "), " ")
		if token == "" {
			username, token, ok = request.BasicAuth()
			println("The Token ", username, token)
			if token == "" || !ok {
				return nil, nil
			}
		}
		session := &model.Session{}
		tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SESSION_SECRET_KEY")), nil
		})
		if err != nil {
			return nil, err
		}
		if tkn.Valid {
			dsession, er := sessh.SessionService.GetSessionByUserID(context.Background(), uint(session.ID))
			if dsession == nil || er != nil {
				return nil, er
			}
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
	if tkn.Valid {
		dsession, er := sessh.SessionService.GetSessionByUserID(context.Background(), uint(session.ID))
		if dsession == nil || er != nil {
			return nil, er
		}
		return session, nil
	}
	return nil, errors.New(" invalid login session ")
}

// GetSubscriberSession(request *http.Request) (*model.Session, error)
func (sessh *authenticator) GetSubscriberSession(request *http.Request) (*model.SubscriberSession, error) {
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
				return nil, nil
			}
		}
		session := &model.SubscriberSession{}
		tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")), nil
		})
		if err != nil {
			return nil, err
		}
		if tkn.Valid {
			dsession, er := sessh.SessionService.GetSubscriberSessionByUserID(context.Background(), uint(session.ID))
			if dsession == nil || er != nil {
				return nil, er
			}
			return session, nil
		}
		return nil, errors.New(" invalid login session ")
	}
	tknStr := cookie.Value
	session := &model.SubscriberSession{}
	tkn, err := jwt.ParseWithClaims(tknStr, session, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	if tkn.Valid {
		dsession, er := sessh.SessionService.GetSubscriberSessionByUserID(context.Background(), uint(session.ID))
		if dsession == nil || er != nil {
			return nil, er
		}
		return session, nil
	}
	return nil, errors.New(" invalid login session ")
}

// GetEmailSession(request *http.Request) (*model.Session, error)
func (sessh *authenticator) GetEmailSession(token string) (*model.EmailConfirmationSession, error) {
	// go and check for the authorization header
	if token == "" {
		return nil, nil
	}
	session := &model.EmailConfirmationSession{}
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

// LogoutSubscriberSession(request *http.Request) (*model.Session, error)
func (sessh *authenticator) LogoutSubscriberSession(request *http.Request) error {
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
				return nil
			}
		}
		session := &model.SubscriberSession{}
		tkn, err := jwt.ParseWithClaims(token, session, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SUBSCRIBER_SESSION_SECRET_KEY")), nil
		})
		if err != nil {
			return err
		}
		if tkn.Valid {
			er := sessh.SessionService.DeleteSubscriberSession(context.Background(), 0, uint(session.ID))
			if er != nil {
				return er
			}
			return nil
		}
		return errors.New(" invalid login session ")
	}
	tknStr := cookie.Value
	session := &model.SubscriberSession{}
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
		er := sessh.SessionService.DeleteSubscriberSession(context.Background(), 0, uint(session.ID))
		if er != nil {
			return er
		}
		return nil
	}
	return errors.New(" invalid login session ")
}

// LogoutSession(request *http.Request) (*model.Session, error)
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
		session := &model.SubscriberSession{}
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
	session := &model.SubscriberSession{}
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
