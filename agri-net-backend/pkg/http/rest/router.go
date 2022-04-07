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
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/middleware"
)

// Route returns an http handler for the api.
func Route(rules middleware.Rules, adminhandler IAdminHandler, roundhandler IRoundHandler, studenthandler IStudentHandler, paymenthandler IPaymentHandler) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8080", "https://facebook.com"},
		AllowHeaders:     []string{"Content-type", "*"},
		AllowCredentials: true,
	}))
	router.GET("/api/password/forgot/", adminhandler.ForgotPassword)
	router.POST("/api/admin/login/", adminhandler.AdminLogin)
	router.GET("/api/deactivate/", adminhandler.DeactivateAccount)
	router.GET("/deactivate/", adminhandler.DeactivateAccount)

	router.GET("/api/logout/", rules.Authenticated(), adminhandler.Logout)
	router.PUT("/api/password/new/", rules.Authenticated(), adminhandler.ChangePassword)
	router.DELETE("/api/admin", rules.Authenticated(), adminhandler.DeleteAdmin)
	// Routes which needs Authentication and Authorization.
	router.POST("/api/admin/new/", rules.Authenticated(), rules.Authorized(), adminhandler.CreateAdmin)
	router.PUT("/api/profile/update/", adminhandler.UpdateAdmin)
	router.POST("/api/profile/picture/new/", rules.Authenticated(), rules.Authorized(), adminhandler.ChangeProfilePicture)
	router.DELETE("/api/profile/picture/", rules.Authenticated(), rules.Authorized(), adminhandler.DeleteProfilePicture)
	router.GET("/api/admins", rules.Authenticated(), adminhandler.GetAdminsOfSystem)

	router.POST("/api/admin/round/new/", rules.Authenticated(), rules.Authorized(), roundhandler.CreateRound)
	router.PUT("/api/admin/round/", rules.Authenticated(), rules.Authorized(), roundhandler.UpdateRound)
	router.GET("/api/category/rounds", roundhandler.GetRoundsOfCategory)
	router.GET("/api/round/activation", rules.Authenticated(), rules.Authorized(), roundhandler.ActivateRound)
	router.GET("/api/round/deactivation", rules.Authenticated(), rules.Authorized(), roundhandler.DeactivateRound)

	// --------------- Not tested -----------------------

	router.POST("/api/student/new/", rules.Authenticated(), rules.Authorized(), studenthandler.RegisterStudent)
	router.PUT("/api/student", rules.Authenticated(), rules.Authorized(), studenthandler.UpdateStudent)
	//
	router.GET("/api/round/students", rules.Authenticated(), studenthandler.GetStudentsOfRound)
	router.GET("/api/category/students", rules.Authenticated(), studenthandler.GetStudentsOfCategory)
	//
	router.PUT("/api/student/profile", rules.Authenticated(), rules.Authorized(), studenthandler.UpdateStudentProfilePicture)
	router.POST("/api/student/special_case", rules.Authenticated(), rules.Authorized(), studenthandler.CreateStudentsSpecialCaseForStudent)
	router.GET("/api/student", rules.Authenticated(), rules.Authorized(), studenthandler.GetStudentByID)
	router.PUT("/api/student/special_case", rules.Authenticated(), rules.Authorized(), studenthandler.UpdateStudentSpecialCaseInformation)

	router.POST("/api/payment/new", rules.Authenticated(), paymenthandler.Payin)
	router.GET("/api/student/payments", rules.Authenticated(), paymenthandler.GetStudentsPayments)

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
