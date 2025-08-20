package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/backend/cmd/main/service"
	"github.com/samuael/shemach/backend/pkg/http/rest"
	"github.com/samuael/shemach/backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/backend/pkg/http/rest/middleware"
	"github.com/samuael/shemach/backend/pkg/resource"
	"github.com/samuael/shemach/backend/pkg/session"
	"github.com/samuael/shemach/backend/pkg/storage/pgx_storage"
	"github.com/samuael/shemach/backend/pkg/storage/pgxconn"
	"github.com/samuael/shemach/backend/pkg/user"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

var once sync.Once
var conn *pgxpool.Pool
var connError error

var templates *template.Template

func main() {
	once.Do(func() {
		conn, connError = pgxconn.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		if connError != nil {
			println(connError.Error())
			os.Exit(1)
		}
		templates = template.Must(template.ParseGlob(os.Getenv("PATH_TO_TEMPLATES") + "*.html"))
	})
	defer conn.Close()
	defer os.Exit(0)

	sessionRepo := pgx_storage.NewSessionRepo(conn)
	sessionService := session.NewSessionService(sessionRepo)
	authenticator := auth.NewAuthenticator(sessionService)

	rules := middleware.NewRules(authenticator)
	userrepo := user.NewUserRepo(conn)
	userservice := user.NewUserService(userrepo)

	otpService := service.NewOtpService(userservice, sessionService)
	go otpService.Run()
	defer func() {
		println("Closing Short Code Management Service ...")
		otpService.DeleteChannel <- true
	}()

	resourcerepo := pgx_storage.NewResourceRepo(conn)
	resourceservice := resource.NewResourceService(resourcerepo)
	resourcehandler := rest.NewResourceHandler(resourceservice)

	userhandler := rest.NewUserHandler(templates, userservice, authenticator, otpService)

	router := rest.Route(rules,
		userhandler,
		resourcehandler,
	)
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		println(err.Error())
	}
	log.Fatal(err)
}
