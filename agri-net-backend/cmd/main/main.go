package main

import (
	"html/template"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/agri-net/agri-net-backend/cmd/main/service"
	"github.com/samuael/agri-net/agri-net-backend/cmd/main/service/message_broadcast_service"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/middleware"
	"github.com/samuael/agri-net/agri-net-backend/pkg/infoadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/message"
	"github.com/samuael/agri-net/agri-net-backend/pkg/product"
	"github.com/samuael/agri-net/agri-net-backend/pkg/storage/pgx_storage"
	pgxstorage "github.com/samuael/agri-net/agri-net-backend/pkg/storage/pgx_storage"
	"github.com/samuael/agri-net/agri-net-backend/pkg/subscriber"
	"github.com/samuael/agri-net/agri-net-backend/pkg/superadmin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/user"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var once sync.Once
var conn *pgxpool.Pool
var connError error

var templates *template.Template

func main() {
	once.Do(func() {
		conn, connError = pgxstorage.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		if connError != nil {
			os.Exit(1)
		}
		templates = template.Must(template.ParseGlob(os.Getenv("PATH_TO_TEMPLATES") + "*.html"))
	})
	defer conn.Close()

	authenticator := auth.NewAuthenticator()
	rules := middleware.NewRules(authenticator)

	subscriberRepo := pgx_storage.NewSubscriberRepo(conn)
	subscriberService := subscriber.NewSubscriberService(subscriberRepo)
	superadminrepo := pgx_storage.NewSuperadminRepo(conn)
	superadminservice := superadmin.NewSuperadminService(superadminrepo)
	userrepo := pgx_storage.NewUserRepo(conn)
	userservice := user.NewUserService(userrepo)

	otpService := service.NewOtpService(subscriberService, userservice)

	superadminhandler := rest.NewSuperadminHandler(superadminservice, authenticator, userservice)
	subscriberhandler := rest.NewSubscriberHandler(authenticator, subscriberService, otpService)

	productrepo := pgx_storage.NewProductRepo(conn)
	productservice := product.NewProductService(productrepo)
	go otpService.Run()
	messagerepo := pgx_storage.NewMessageRepo(conn)
	messageservice := message.NewMessageService(messagerepo)
	messagehandler := rest.NewMessageHandler(messageservice, subscriberService)

	infoadminrepo := pgx_storage.NewInfoadminRepo(conn)
	infoadminservice := infoadmin.NewInfoadminService(infoadminrepo)
	infoadminhandler := rest.NewInfoAdminHandler(infoadminservice)

	broadcastHub := message_broadcast_service.NewMainBroadcastHub(messageservice)
	producthandler := rest.NewProductHandler(productservice, broadcastHub)

	adminrepo := pgx_storage.NewAdminRepo(conn)
	adminservice := admin.NewAdminService(adminrepo)
	adminhandler := rest.NewAdminHandler(adminservice)

	userhandler := rest.NewUserHandler(userservice, authenticator)

	communicationHandler := message_broadcast_service.NewClientConnectionHandler(
		subscriberService,
		broadcastHub,
	)

	go broadcastHub.Run()
	rest.Route(rules, subscriberhandler,
		superadminhandler, producthandler,
		communicationHandler, messagehandler,
		infoadminhandler,
		userhandler,
		adminhandler).Run(":8080")
}
