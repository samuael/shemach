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
	transactionhandler ITransactionHandler,
) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowOrigins:     []string{"http://localhost:8080"},
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
	router.GET("/api/connection/subscriber/:id", communicationHandler.SubscriberHandleWebsocketConnection)
	router.GET("/api/connection/admins/:id", communicationHandler.AdminsHandleWebsocketConnection)

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
	router.POST("/api/dictionary/translate", dictionaryhandler.Translate)
	router.GET("/api/dictionaries", dictionaryhandler.GetRecentDictionaries)
	// Store related routes
	router.POST("/api/admin/store/new", rules.Authenticated(), rules.Authorized(), storehandler.CreateStore)
	router.GET("/api/merchant/stores", rules.Authenticated(), storehandler.GetMerchantStores)
	router.GET("/api/store", rules.Authenticated(), storehandler.GetMerchantByID)
	router.GET("/api/merchants", rules.Authenticated(), merchanthandler.MerchantsSearch)
	router.GET("/api/agents", rules.Authenticated(), agenthandler.AgentsSearch)

	// Crop related routes
	// This routes are applicable for only Merchants and Agents
	router.POST("/api/cxp/post/new", rules.Authenticated(), rules.Authorized(), crophandler.CreateProduct)
	router.POST("/api/cxp/post/images/:postid", rules.Authenticated(), rules.Authorized(), crophandler.UploadProductImages)

	router.GET("/api/posts", rules.Authenticated(), crophandler.Getposts)
	router.GET("/api/cxp/posts", rules.Authenticated(), crophandler.GetMyPosts)

	router.GET("/api/post/:id", rules.Authenticated(), crophandler.GetPostByID)

	router.GET("/post/image/:id", rules.Authenticated(), resourcehandler.GetProductImage)
	router.GET("/post/image/:id/blurred/", rules.Authenticated(), resourcehandler.GetBlurredImage)

	router.GET("/api/merchant/product/subscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.SubscribeForProduct)
	router.GET("/api/merchant/product/unsubscribe/:id", rules.Authenticated(), rules.Authorized(), merchanthandler.UnsubscriberForProduct)

	// transaction related routes
	router.POST("/api/merchant/transaction/new", rules.Authenticated(), rules.Authorized(), transactionhandler.CreateTransaction)
	router.GET("/api/cxp/transactions", rules.Authenticated(), rules.Authorized(), transactionhandler.GetMyActiveTransactions)
	router.DELETE("/api/cxp/transaction/:id", rules.Authenticated(), rules.Authorized(), transactionhandler.DeclineTransaction)
	router.POST("/api/cxp/transaction/amend", rules.Authenticated(), rules.Authorized(), transactionhandler.TransactionAmendmenRequest)
	router.GET("/api/merchant/transaction/request/accept/:id", rules.Authenticated(), rules.Authorized(), transactionhandler.AcceptAmendmentRequest)
	router.POST("/api/merchant/transaction/request/ammend", rules.Authenticated(), rules.Authorized(), transactionhandler.PerformAmend)
	// ------------------------------------------------------------------------------------------------------------------------
	router.POST("/api/cxp/transaction/request/kebd", rules.Authenticated(), rules.Authorized(), transactionhandler.RequestKebd)
	router.POST("/api/merchant/transaction/kebd/request/amendment", rules.Authenticated(), rules.Authorized(), transactionhandler.RequestKebdRequestAmendment)
	router.POST("/api/merchant/transaction/kebd/amend", rules.Authenticated(), rules.Authorized(), transactionhandler.AmendKebdRequest)

	router.POST("/api/merchant/transaction/request/guarantee", rules.Authenticated(), rules.Authorized(), transactionhandler.RequestGuaranteePayment)
	router.GET("/api/cxp/transaction/seller/accept/:id", rules.Authenticated(), rules.Authorized(), transactionhandler.SellerAcceptedTransaction)
	router.GET("/api/merchant/transaction/buyer/accept/:id", rules.Authenticated(), rules.Authorized(), transactionhandler.BuyerAcceptTransaction)
	router.GET("/api/cxp/mytransactions", rules.Authenticated(), rules.Authorized(), transactionhandler.GetMyTransactionNotifications)
	router.GET("/api/cxp/transaction/reactivate/:id", rules.Authenticated(), rules.Authorized(), transactionhandler.ReactivateTransaction)

	router.RouterGroup.Use(FilterDirectory())
	{
		router.StaticFS("/images/", http.Dir(os.Getenv("ASSETS_DIRECTORY")))
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
