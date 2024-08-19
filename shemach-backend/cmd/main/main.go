package main

import (
	"html/template"
	"os"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samuael/shemach/shemach-backend/cmd/main/service"
	"github.com/samuael/shemach/shemach-backend/pkg/admin"
	"github.com/samuael/shemach/shemach-backend/pkg/agent"
	"github.com/samuael/shemach/shemach-backend/pkg/http/rest"
	"github.com/samuael/shemach/shemach-backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/shemach-backend/pkg/http/rest/middleware"
	"github.com/samuael/shemach/shemach-backend/pkg/merchant"
	"github.com/samuael/shemach/shemach-backend/pkg/message"
	"github.com/samuael/shemach/shemach-backend/pkg/product"
	"github.com/samuael/shemach/shemach-backend/pkg/resource"
	"github.com/samuael/shemach/shemach-backend/pkg/session"
	"github.com/samuael/shemach/shemach-backend/pkg/storage/pgx_storage"
	pgxstorage "github.com/samuael/shemach/shemach-backend/pkg/storage/pgx_storage"
	"github.com/samuael/shemach/shemach-backend/pkg/store"
	"github.com/samuael/shemach/shemach-backend/pkg/subscriber"
	"github.com/samuael/shemach/shemach-backend/pkg/superadmin"
	"github.com/samuael/shemach/shemach-backend/pkg/user"
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
		conn, connError = pgxstorage.NewStorage()
		if connError != nil {
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

	userrepo := pgx_storage.NewUserRepo(conn)
	userservice := user.NewUserService(userrepo)

	subscriberRepo := pgx_storage.NewSubscriberRepo(conn)
	subscriberService := subscriber.NewSubscriberService(subscriberRepo)
	superadminrepo := pgx_storage.NewSuperadminRepo(conn)
	superadminservice := superadmin.NewSuperadminService(superadminrepo)

	otpService := service.NewOtpService(subscriberService, userservice)

	superadminhandler := rest.NewSuperadminHandler(superadminservice, authenticator, userservice)

	productrepo := pgx_storage.NewProductRepo(conn)
	productservice := product.NewProductService(productrepo)
	go otpService.Run()
	messagerepo := pgx_storage.NewMessageRepo(conn)
	messageservice := message.NewMessageService(messagerepo)

	messagehandler := rest.NewMessageHandler(messageservice, subscriberService)

	producthandler := rest.NewProductHandler(productservice)

	adminrepo := pgx_storage.NewAdminRepo(conn)
	adminservice := admin.NewAdminService(adminrepo)
	adminhandler := rest.NewAdminHandler(adminservice)

	agentrepo := pgx_storage.NewAgentRepo(conn)
	agentservice := agent.NewAgentService(agentrepo)
	agenthandler := rest.NewAgentHandler(agentservice, userservice)

	merchantrepo := pgx_storage.NewMerchantRepo(conn)
	merchantservice := merchant.NewMerchantService(merchantrepo)
	merchanthandler := rest.NewMerchantHandler(merchantservice, userservice)

	storerepo := pgx_storage.NewStoreRepo(conn)
	storeservice := store.NewStoreService(storerepo)
	storehandler := rest.NewStoreHandler(storeservice)

	resourcerepo := pgx_storage.NewResourceRepo(conn)
	resourceservice := resource.NewResourceService(resourcerepo)

	resourcehandler := rest.NewResourceHandler(resourceservice)

	userhandler := rest.NewUserHandler(templates,
		userservice, authenticator,
		adminservice, superadminservice,
		agentservice, merchantservice,
		storeservice)
	rest.Route(rules,
		superadminhandler,
		producthandler,
		messagehandler,
		userhandler,
		adminhandler,
		agenthandler,
		merchanthandler,
		storehandler,
		resourcehandler,
	).Run(":8080")
}
