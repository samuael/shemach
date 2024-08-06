package middleware

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/model"
	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
)

type Rules interface {
	Authenticated() gin.HandlerFunc
	AuthenticatedSubscriber() gin.HandlerFunc
	Authorized() gin.HandlerFunc
	HasPermission(path, role, method string) bool
}

type rules struct {
	auth auth.Authenticator
}

func NewRules(auth auth.Authenticator) Rules {
	return &rules{auth}
}

// LoggedIn simple middleware to push value to the context
func (m rules) Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := m.auth.GetSession(c.Request)
		if err != nil || t == nil {
			if err != nil {
				println(err.Error())
			} else {
				println("The session is nil")
			}
			http.Error(c.Writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "session", t)
		success := m.auth.SaveSession(c.Writer, t)
		if !success {
			c.Abort()
			return
		}
		ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*5))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (m rules) AuthenticatedSubscriber() gin.HandlerFunc {
	return func(c *gin.Context) {
		t, err := m.auth.GetSubscriberSession(c.Request)
		if err != nil || t == nil {
			if err != nil {
				println(err.Error())
			}
			http.Error(c.Writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "session", t)
		success := m.auth.SaveSubscriberSession(c.Writer, t)
		if !success {
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (m rules) AuthenticatedEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Trim(" ", strings.TrimPrefix("Bearer ", (c.Request.Header.Get("Authorization"))))
		t, err := m.auth.GetEmailSession(token)
		if err != nil || t == nil {
			if err != nil {
				println(err.Error())
			}
			http.Error(c.Writer, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			c.Abort()
			return
		}
		ctx := context.WithValue(c.Request.Context(), "session", t)
		_, success := m.auth.SaveEmailConfirmationSession(t)
		if !success {
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// Authorized checks if a user has proper authority to access a give route
func (m *rules) Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := m.auth.GetSession(c.Request)
		if err != nil || session == nil {
			if err != nil {
				println(err.Error())
			}
			http.Error(c.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			c.Abort()
			return
		}
		permitted := m.HasPermission(c.Request.URL.Path, session.Role, c.Request.Method)
		if !permitted {
			http.Error(c.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			c.Abort()
			return
		}
		if c.Request.Method == http.MethodPost {
			erro := c.Request.ParseForm()
			if erro != nil {
				http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func (m *rules) HasPermission(path, role, method string) bool {
	// method = strings.ToUpper(method)

	// if permission := state.Authorities[path]; permission != nil {
	if strings.HasPrefix(path, "/api/superadmin/") && (role == state.SUPERADMIN) {
		return true
	} else if strings.HasPrefix(path, "/api/infoadmin/") && (role == state.INFO_ADMIN) {
		return true
	} else if strings.HasPrefix(path, "/api/admin/") && (role == state.ADMIN) {
		return true
	} else if (strings.HasPrefix(path, "/api/agent/") || strings.HasPrefix(path, "/api/cxp/")) && (role == state.AGENT) {
		return true
	} else if (strings.HasPrefix(path, "/api/merchant/") || strings.HasPrefix(path, "/api/cxp/")) && (role == state.MERCHANT) {
		return true
	}
	return false
	// }
	// return false
}

// Logout function api Logging out
// METHOD GET
// VAriables NONE
func (m rules) Logout(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Authorization", "")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(helper.MarshalThis(model.LoginResponse{Success: true}))
}
