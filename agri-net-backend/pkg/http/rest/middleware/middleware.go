package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/model"
	"github.com/samuael/Project/RegistrationSystem/pkg/constants/state"
	"github.com/samuael/Project/RegistrationSystem/pkg/http/rest/auth"
	"github.com/samuael/Project/RegistrationSystem/platforms/helper"
)

type Rules interface {
	Authenticated() gin.HandlerFunc
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
		// ctx.Deadline( time.Now().Add(time.Second * 5))
		c.Request = c.Request.WithContext(ctx)
		log.Println(" Authenticated ... ")
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
		role := session.Role
		permitted := m.HasPermission(c.Request.URL.Path, role, c.Request.Method)
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
		println(" Authorized ... ")
		c.Next()
	}
}

func (m *rules) HasPermission(path, role, method string) bool {
	// method = strings.ToUpper(method)

	// if permission := state.Authorities[path]; permission != nil {
	if strings.HasPrefix(path, "/api/admin/") && (role == state.SUPERADMIN) {
		return true
	} else if strings.HasPrefix(path, "/api/") && !(strings.HasPrefix(path, "/api/admin/")) && (role == state.ADMIN || role == state.SUPERADMIN) {
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
