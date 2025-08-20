package rest

import (
	"context"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	_ "github.com/samuael/shemach/backend/api"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/backend/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(
	rules middleware.Rules,
	userhandler IUserHandler,
	resourceHandler IResourceHandler,
) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/register-verify-phone", userhandler.RegisterPhoneVerify)
	router.POST("/api/register", userhandler.RegisterUser)
	router.POST("/api/confirm-registration", userhandler.ConfirmCode)
	router.POST("/api/login", userhandler.Login)
	router.GET("/api/user/detail/:userID", userhandler.GetUserByID)

	router.GET("/api/user/login-methods", userhandler.GetUserLoginMethods)
	router.GET("/api/forgot-password/send-shortcode", userhandler.GenerateForgotPasswordShortCode)
	router.POST("/api/forgot-password/confirm-shortcode", userhandler.ForgotPasswordShortcodeConfirmation)
	router.POST("/api/password/new", rules.Authenticated(userhandler.ChangePassword))

	router.GET("/api/image/:id", resourceHandler.GetImageByID)

	// Global CORS handler for preflight requests
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.WriteHeader(http.StatusOK)
	})
	// router.PanicHandler = PanicHandler
	return router
}

type StaticFilesServer struct {
	FileHandler http.Handler
}

func (c StaticFilesServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	FilterDirectory(c.FileHandler.ServeHTTP)
}

// AccessControl ... a method.
func AccessControl(h httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		h(w, r, params)
	})
}
func FilterDirectory(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "../") ||
			strings.HasSuffix(r.URL.Path, "/") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`<html>UnAuthorized!</html>`))
			return
		}
		handler(w, r)
	}
}

// SessionFromContext extracts session information from context or returns an error message otherwise
func SessionFromContext(ctx context.Context) (*types.Session, error) {
	sessionVal, okay := ctx.Value(types.AuthParam("session")).(*types.Session)
	if !okay {
		return nil, auth.ErrSessionNotFound
	}
	return sessionVal, nil
}

// func PanicHandler(w http.ResponseWriter, r *http.Request, result interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	panicMessage := fmt.Sprintf("Error occured: %v", result)
// 	println(panicMessage)
// 	w.Write(helper.MarshalThis(&types.StatusMsg{Success: false, Status: types.ScInternalError, Error: panicMessage}))
// }
