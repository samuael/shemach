package rest

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	_ "github.com/samuael/shemach/shemach-backend/api"
	"github.com/samuael/shemach/shemach-backend/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(
	rules middleware.Rules,
	superadminhandler ISuperadminHandler,
	producthandler IProductHandler,
	messagehandler IMessageHandler,
	userhandler IUserHandler,
	adminhandler IAdminHandler,
	agenthandler IAgentHandler,
	merchanthandler IMerchantHandler,
	storehandler IStoreHandler,
	resourcehandler IResourceHandler,
) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Content-type", "*"},
		AllowCredentials: true,
	}))
	router.GET("/api/superadmin", rules.Authenticated(), superadminhandler.GetSuperadminByID)
	router.GET("/api/superadmins", rules.Authenticated(), superadminhandler.GetSystemSuperadmin)

	// -------------------------------------------------------------------------------
	router.POST("/api/login", userhandler.Login)

	// product related end points
	router.POST("/api/superadmin/product/new", rules.Authenticated(), rules.Authorized(), producthandler.CreateProductInstance)
	router.GET("/api/product", producthandler.GetProductByID)
	router.GET("/api/products", producthandler.GetProducts)
	// router.GET("/api/product/subscribe", rules.AuthenticatedSubscriber(), producthandler.SubscribeForProduct)
	// router.GET("/api/product/unsubscribe", rules.AuthenticatedSubscriber(), producthandler.UnsubscriberForProduct)
	router.GET("/api/product/search", producthandler.SearchProduct)
	router.GET("/api/product/units", producthandler.GetProductUnits)
	router.PUT("/api/infoadmin/product", rules.Authenticated(), rules.Authorized(), producthandler.UpdateProduct)

	// message related end points
	router.GET("/api/messages", rules.AuthenticatedSubscriber(), messagehandler.GetRecentMessages)
	router.DELETE("/api/message/:id", rules.Authenticated(), messagehandler.DeleteMessageByID)
	router.POST("/api/message/new", rules.Authenticated(), messagehandler.SendMessage)
	router.GET("/api/admins/messages", rules.Authenticated(), messagehandler.GetAllMessages)

	// user realted methods
	router.PUT("/api/user/password", rules.Authenticated(), userhandler.ChangePassword)
	router.PUT("/api/user/profile/picture", rules.Authenticated(), userhandler.UpdateProfilePicture)
	router.DELETE("/api/user/profile/picture", rules.Authenticated(), userhandler.DeleteProfilePicture)
	router.PUT("/api/user", rules.Authenticated(), userhandler.UpdateProfile)
	router.GET("/api/user/:id", rules.Authenticated(), userhandler.GetUserByID)

	// admin related routes
	router.POST("/api/superadmin/admin/new/", rules.Authenticated(), rules.Authorized(), adminhandler.RegisterAdmin)
	router.GET("/api/admins", rules.Authenticated(), adminhandler.ListAdmins)
	router.DELETE("/api/superadmin/admin", rules.Authenticated(), rules.Authorized(), adminhandler.DeleteAdminByID)
	router.GET("/api/admin", rules.Authenticated(), rules.Authorized(), adminhandler.GetAdminByID)

	// agents related enpoints
	router.POST("/api/admin/agent/new", rules.Authenticated(), rules.Authorized(), agenthandler.RegisterAgent)
	router.DELETE("/api/admin/agent/:id", rules.Authenticated(), rules.Authorized(), agenthandler.DeleteAgentByID)

	// merchants related endpoints
	router.POST("/api/admin/merchant/new", rules.Authenticated(), rules.Authorized(), merchanthandler.RegisterMerchant)
	router.DELETE("/api/admin/merchant/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.DeleteMerchantByID)

	// CXP(  Commodity exchage participant ) related routes
	router.POST("/api/cxp/account/confirm", userhandler.ConfirmTempoCXP)
	router.GET("/api/user/account/email/confirm", userhandler.ConfirmEmail)

	// Store related routes
	router.POST("/api/admin/store/new", rules.Authenticated(), rules.Authorized(), storehandler.CreateStore)
	router.GET("/api/merchant/stores", rules.Authenticated(), storehandler.GetMerchantStores)
	router.GET("/api/store", rules.Authenticated(), storehandler.GetMerchantByID)
	router.GET("/api/merchants", rules.Authenticated(), merchanthandler.MerchantsSearch)
	router.GET("/api/store/merchant/:storeid", rules.Authenticated(), userhandler.GetMerchantByStoreID)
	router.GET("/api/agents", rules.Authenticated(), agenthandler.AgentsSearch)

	router.GET("/post/image/:id", rules.Authenticated(), resourcehandler.GetProductImage)
	router.GET("/post/image/:id/blurred/" /*rules.Authenticated(), */, resourcehandler.GetBlurredImage)

	router.GET("/api/merchant/product/subscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.SubscribeForProduct)
	router.GET("/api/merchant/product/unsubscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.UnsubscriberForProduct)

	router.GET("/api/logout", rules.Authenticated(), userhandler.Logout)
	router.GET("/api/subscriber/logout", rules.AuthenticatedSubscriber(), userhandler.LogoutSubscriber)

	//
	router.RouterGroup.Use(FilterDirectory())
	{
		router.StaticFS("/images/", http.Dir(os.Getenv("ASSETS_DIRECTORY")+"/images/"))
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
