package rest

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	_ "github.com/samuael/agri-net/agri-net-backend/api"
	"github.com/samuael/agri-net/agri-net-backend/cmd/main/service/message_broadcast_service"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(
	rules middleware.Rules,
	subscriberhandler ISubscriberHandler,
	superadminhandler ISuperadminHandler,
	producthandler IProductHandler,
	communicationHandler message_broadcast_service.IClientConnetionHandler,
	messagehandler IMessageHandler,
	infoadminhandler IInfoadminHandler,
	userhandler IUserHandler,
	adminhandler IAdminHandler,
	agenthandler IAgentHandler,
	dictionaryhandler IDictionaryHandler,
	merchanthandler IMerchantHandler,
	storehandler IStoreHandler,
	crophandler ICropHandler,
	resourcehandler IResourceHandler,
) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8080", "https://facebook.com"},
		AllowHeaders:     []string{"Content-type", "*"},
		AllowCredentials: true,
	}))

	router.GET("/api/superadmin", rules.Authenticated(), superadminhandler.GetSuperadminByID)
	router.GET("/api/superadmins", rules.Authenticated(), superadminhandler.GetSystemSuperadmin)

	// subscriber related endpoints
	router.POST("/api/info/register", subscriberhandler.RegisterSubscriber)
	router.POST("/api/subscription/registration/confirm", subscriberhandler.ConfirmRegistrationSubscription)
	router.POST("/api/subscription/login", subscriberhandler.SubscriberLoginWithPhone)
	router.POST("/api/subscription/confirm", subscriberhandler.ConfirmLoginSubscription)
	// -------------------------------------------------------------------------------
	router.POST("/api/login", userhandler.Login)

	// product related end points
	router.POST("/api/superadmin/product/new", rules.Authenticated(), rules.Authorized(), producthandler.CreateProductInstance)
	router.GET("/api/product", producthandler.GetProductByID)
	router.GET("/api/products", producthandler.GetProducts)
	router.GET("/api/product/subscribe", rules.AuthenticatedSubscriber(), producthandler.SubscribeForProduct)
	router.GET("/api/product/unsubscribe", rules.AuthenticatedSubscriber(), producthandler.UnsubscriberForProduct)
	router.GET("/api/product/search", producthandler.SearchProduct)
	router.GET("/api/product/units", producthandler.GetProductUnits)
	router.PUT("/api/infoadmin/product", rules.Authenticated(), rules.Authorized(), producthandler.UpdateProduct)

	// web socket end points
	router.GET("/api/connection/subscriber", rules.AuthenticatedSubscriber(), communicationHandler.SubscriberHandleWebsocketConnection)
	router.GET("/api/connection/admins", rules.Authenticated(), communicationHandler.AdminsHandleWebsocketConnection)

	// message related end points
	router.GET("/api/messages", rules.AuthenticatedSubscriber(), messagehandler.GetRecentMessages)

	// infoadmin related routes
	router.POST("/api/superadmin/infoadmin/new", rules.Authenticated(), rules.Authorized(), infoadminhandler.Registerinfoadmin)
	router.GET("/api/infoadmins", infoadminhandler.ListInfoadmins)
	router.DELETE("/api/superadmin/infoadmin", infoadminhandler.DeleteInfoadminByID)
	router.GET("/api/infoadmin", rules.Authenticated(), infoadminhandler.GetInfoadminByID)

	// user realted methods
	router.PUT("/api/user/password", rules.Authenticated(), userhandler.ChangePassword)
	router.PUT("/api/user/profile/picture", rules.Authenticated(), userhandler.UpdateProfilePicture)
	router.DELETE("/api/user/profile/picture", rules.Authenticated(), userhandler.DeleteProfilePicture)
	router.PUT("/api/user", rules.Authenticated(), userhandler.UpdateProfile)

	// admin related routes
	router.POST("/api/superadmin/admin/new/", rules.Authenticated(), rules.Authorized(), adminhandler.RegisterAdmin)
	router.GET("/api/admins", rules.Authenticated(), adminhandler.ListAdmins)
	router.DELETE("/api/superadmin/admin", rules.Authenticated(), rules.Authorized(), adminhandler.DeleteAdminByID)
	router.GET("/api/admin", rules.Authenticated(), rules.Authorized(), adminhandler.GetAdminByID)

	// agents related enpoints
	router.POST("/api/admin/agent/new", rules.Authenticated(), rules.Authorized(), agenthandler.RegisterAgent)

	// merchants related endpoints
	router.POST("/api/admin/merchant/new", rules.Authenticated(), rules.Authorized(), merchanthandler.RegisterMerchant)

	// CXP(  Commodity exchage participant ) related routes
	router.POST("/api/cxp/account/confirm", userhandler.ConfirmTempoCXP)
	router.GET("/api/user/account/email/confirm", userhandler.ConfirmEmail)

	// Dictionary related routes
	router.POST("/api/superadmin/dictionary/new", rules.Authenticated(), rules.Authorized(), dictionaryhandler.CreateDictionary)
	router.PUT("/api/superadmin/dictionary", rules.Authenticated(), rules.Authorized(), dictionaryhandler.UpdateDictionary)
	router.DELETE("/api/superadmin/dictionary", rules.Authenticated(), dictionaryhandler.DeleteDictionary)
	router.GET("/api/dictionary/translate", dictionaryhandler.Translate)

	// Store related routes
	router.POST("/api/admin/store/new", rules.Authenticated(), rules.Authorized(), storehandler.CreateStore)
	router.GET("/api/merchant/stores", rules.Authenticated(), storehandler.GetMerchantStores)
	router.GET("/api/store", rules.Authenticated(), storehandler.GetMerchantByID)

	// Crop related routes
	// This routes are applicable for only Merchants and Agents
	router.POST("/api/cxp/post/new", rules.Authenticated(), rules.Authorized(), crophandler.CreateProduct)
	router.POST("/api/cxp/post/images/:postid", rules.Authenticated(), rules.Authorized(), crophandler.UploadProductImages)
	router.GET("/api/posts", rules.Authenticated(), crophandler.Getposts)
	router.GET("/api/post/:id", rules.Authenticated(), crophandler.GetPostByID)

	router.GET("/post/image/:id", rules.Authenticated(), resourcehandler.GetProductImage)
	router.GET("/post/image/:id/blurred/", rules.Authenticated(), resourcehandler.GetBlurredImage)

	router.GET("/api/merchant/product/subscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.SubscribeForProduct)
	router.GET("/api/merchant/product/unsubscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.UnsubscriberForProduct)

	router.RouterGroup.Use(FilterDirectory())
	{
		router.StaticFS("/images/", http.Dir(os.Getenv("ASSETS_DIRECTORY")+"images/"))
	}
	return router
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
func FilterDirectory() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(" Filter Directory ")
		if strings.HasSuffix(c.Request.URL.Path, "/") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
