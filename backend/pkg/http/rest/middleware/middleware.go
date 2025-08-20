package middleware

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/backend/platforms/helper"
)

type Rules interface {
	CheckIfAuthentication(httprouter.Handle) httprouter.Handle
	Authenticated(httprouter.Handle) httprouter.Handle
	Authorized(httprouter.Handle) httprouter.Handle
	HasPermission(path, role, method string) bool
}

type rules struct {
	auth auth.Authenticator
}

func NewRules(auth auth.Authenticator) Rules {
	return &rules{auth}
}

func (m rules) CheckIfAuthentication(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Gets the token from header and checks the validity of the session
		session, err := m.auth.GetSession(r)
		if err == nil && session != nil {
			ctx := r.Context()
			ctx = context.WithValue(ctx, types.AuthParam("session"), session)
			r = r.WithContext(ctx)
		}
		handler(w, r, params)
	}
}

// LoggedIn simple middleware to push value to the context
func (m rules) Authenticated(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		println("Authenticated Header: ", string(helper.MarshalThis(r.Header)))
		// Gets the token from header and checks the validity of the session
		session, err := m.auth.GetSession(r)
		if err != nil || session == nil {
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, types.AuthParam("session"), session)
		r = r.WithContext(ctx)
		handler(w, r, params)
	}
}

func (m rules) AuthenticatedEmail() httprouter.Handle {
	return nil
}

// Authorized checks if a user has proper authority to access a give route
func (m *rules) Authorized(httprouter.Handle) httprouter.Handle {
	return nil
	// func(c *gin.Context) {
	// 	session, err := m.auth.GetSession(c.Request)
	// 	if err != nil || session == nil {
	// 		if err != nil {
	// 			println(err.Error())
	// 		}
	// 		http.Error(c.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	// 		c.Abort()
	// 		return
	// 	}
	// 	permitted := m.HasPermission(c.Request.URL.Path, session.Role, c.Request.Method)
	// 	if !permitted {
	// 		http.Error(c.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	// 		c.Abort()
	// 		return
	// 	}
	// 	if c.Request.Method == http.MethodPost {
	// 		erro := c.Request.ParseForm()
	// 		if erro != nil {
	// 			http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 			c.Abort()
	// 			return
	// 		}
	// 	}
	// 	c.Next()
	// }
}

func (m *rules) HasPermission(path, role, method string) bool {
	// if strings.HasPrefix(path, "/api/superadmin/") && (role == state.SUPERADMIN) {
	// 	return true
	// } else if strings.HasPrefix(path, "/api/infoadmin/") && (role == state.INFO_ADMIN) {
	// 	return true
	// } else if strings.HasPrefix(path, "/api/admin/") && (role == state.ADMIN) {
	// 	return true
	// } else if (strings.HasPrefix(path, "/api/agent/") || strings.HasPrefix(path, "/api/cxp/")) && (role == state.AGENT) {
	// 	return true
	// } else if (strings.HasPrefix(path, "/api/merchant/") || strings.HasPrefix(path, "/api/cxp/")) && (role == state.MERCHANT) {
	// 	return true
	// }
	return false
}

// Logout function api Logging out
// METHOD GET
// VAriables NONE
func (m rules) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "")
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(types.LoginResponse{Success: true}))
}
