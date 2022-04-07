package rest

import (
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/admin"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/model"
	"github.com/samuael/agri-net/agri-net-backend/pkg/constants/state"
	"github.com/samuael/agri-net/agri-net-backend/pkg/http/rest/auth"
	"github.com/samuael/agri-net/agri-net-backend/platforms"
	"github.com/samuael/agri-net/agri-net-backend/platforms/form"
	"github.com/samuael/agri-net/agri-net-backend/platforms/hash"
	"github.com/samuael/agri-net/agri-net-backend/platforms/helper"
	"github.com/samuael/agri-net/agri-net-backend/platforms/mail"
)

// IAdminHandler interface
type IAdminHandler interface {
	AdminLogin(c *gin.Context)
	ChangePassword(c *gin.Context)
	ForgotPassword(c *gin.Context)
	DeactivateAccount(c *gin.Context)
	DeleteProfilePicture(c *gin.Context)
	ChangeProfilePicture(c *gin.Context)
	CreateAdmin(c *gin.Context)
	Logout(c *gin.Context)
	UpdateAdmin(c *gin.Context)
	//  uses "email": string   or  "id":int
	DeleteAdmin(c *gin.Context)
	GetAdminsOfSystem(c *gin.Context)
}

// AdminHandler ... |  ...
type AdminHandler struct {
	Authenticator auth.Authenticator
	AdminSer      admin.IAdminService
}

// NewAdminHandler ... | ...
func NewAdminHandler(auths auth.Authenticator, adminser admin.IAdminService) IAdminHandler {
	return &AdminHandler{
		AdminSer:      adminser,
		Authenticator: auths,
	}
}

/*
	AdminLogin to handle a login request for an admin ....
	METHOD : POST
	INPUT  : JSON
	INPUT : {
		"email"  : "email" ,
		"password"  : "passs"
	}
	OUTPUT : {
		"success" : true ,
		"message" : "Success message" ,
		"admin" : {
			"id" : 3 ,
			"email" : ""
		}
	}
*/
func (adminhr *AdminHandler) AdminLogin(c *gin.Context) {
	input := &struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	// LoginResponse ...
	resp := &model.LoginResponse{}
	resp.Success = false

	// jdecode := json.NewDecoder(request.Body)
	err := c.Bind(input)
	if err != nil || input.Email == "" || input.Password == "" {
		resp.Message = os.Getenv("INVALID_INPUT")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, "email", input.Email)
	newAdmin, err := adminhr.AdminSer.AdminByEmail(ctx)
	if err != nil || newAdmin == nil {
		resp.Success = false
		resp.Message = "Invalid Username or Password!"
		c.JSON(401, resp)
		return
	} else {
		if newAdmin == nil {
			goto InvalidUsernameOrPassword
		}
		// comparing the hashed password and the password
		matches := hash.ComparePassword(newAdmin.Password, input.Password)
		if !matches {
			goto InvalidUsernameOrPassword
		}
		session := &model.Session{
			ID:       uint64(newAdmin.ID),
			Email:    newAdmin.Email,
			Password: input.Password,
		}
		if newAdmin.Superadmin {
			session.Role = state.SUPERADMIN
		} else {
			session.Role = state.ADMIN
		}
		success := adminhr.Authenticator.SaveSession(c.Writer, session)
		if !success {
			resp.Message = os.Getenv("INTERNAL_SERVER_ERROR")
			resp.Success = false
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		resp.Success = true
		resp.Message = state.SuccesfulyLoggedIn
		resp.User = newAdmin
		c.JSON(200, resp)
		return
	}
	// InvalidUsernameOrPassword
InvalidUsernameOrPassword:
	{
		resp.Success = false
		resp.Message = state.InvalidUsernameORPassword
		c.JSON(404, resp)
		return
	}
}

// Logout || method GET /for an admin to log out
func (adminhr *AdminHandler) Logout(c *gin.Context) {
	adminhr.Authenticator.DeleteSession(c.Writer, c.Request)
}

// ChangePassword ... method to change the password for all the three roles
// METHOD  : PUT
// INPUT : JSON
/*
	{
		"old_password" : "theoldpassword" ,
		"new_password" : "new_password " ,
		"confirm_password" : "new_password_here"
	}

	OUTPUT : JSON

	{
		"success" : true ,
		"message" : "Password changed succesfuly "
	}
*/
func (adminhr *AdminHandler) ChangePassword(c *gin.Context) {
	ctx := c.Request.Context()
	session := ctx.Value("session").(*model.Session)

	res := &model.SimpleSuccessNotifier{
		Success: false,
	}
	input := &struct {
		Oldpassword     string `json:"old_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	jdecoder := json.NewDecoder(c.Request.Body)
	era := jdecoder.Decode(input)
	if era != nil || input.Oldpassword == "" || input.NewPassword == "" || input.ConfirmPassword == "" {
		res.Message = os.Getenv("BAD_REQUEST_BODY")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if input.ConfirmPassword != input.NewPassword {
		res.Message = os.Getenv("RE_CONFIRM_PASSWORD")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if len(input.NewPassword) < 4 {
		res.Message = "Password Length Must exceed 4 characters! "
		c.JSON(http.StatusBadRequest, res)
		return
	}
	var changesuccess bool
	ctx = context.WithValue(ctx, "user_id", uint(session.ID))
	hashed, era := hash.HashPassword(input.NewPassword)
	if era != nil {
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	ctx = context.WithValue(ctx, "password", hashed)
	changesuccess, era = adminhr.AdminSer.ChangePassword(ctx)
	if era != nil || !changesuccess {
		res.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		res.Success = false
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Message = "Password Changed Succesfuly!"
	res.Success = true
	c.JSON(http.StatusOK, res)
}

// Method ForgotPassword method GET
func (adminhr *AdminHandler) ForgotPassword(c *gin.Context) {
	input := &struct {
		Email string `json:"email"`
	}{}

	respo := &struct {
		Message string `json:"msg"`
	}{}
	if input.Email = c.Request.FormValue("email"); input.Email == "" {
		respo.Message = "Expected Email not found!"
		c.JSON(http.StatusBadRequest, respo)
		return
	}
	// session, _ := adminhr.Authenticator.GetSession(request)
	ctx := c.Request.Context()
	if !form.MatchesPattern(input.Email, form.EmailRX) {
		respo.Message = "Invalid email address!"
		c.JSON(http.StatusBadRequest, respo)
		return
	}
	ctx = context.WithValue(ctx, "email", input.Email)
	admin, er := adminhr.AdminSer.AdminByEmail(ctx)
	if admin != nil && er == nil {
		password := helper.GenerateRandomString(5, helper.NUMBERS)
		if success := mail.SendPasswordEmailSMTP([]string{admin.Email}, password, false, admin.Fullname, c.Request.Host); success {
			hashed, era := hash.HashPassword(password)
			if era != nil {
				respo.Message = os.Getenv("INTERNAL_SERVER_ERROR")
				c.JSON(http.StatusInternalServerError, respo)
				return
			}
			ctx = context.WithValue(ctx, "user_id", uint(admin.ID))
			ctx = context.WithValue(ctx, "password", hashed)
			changesuccess, era := adminhr.AdminSer.ChangePassword(ctx)
			if era != nil || !changesuccess {
				respo.Message = os.Getenv("INTERNAL_SERVER_ERROR")
				c.JSON(http.StatusInternalServerError, respo)
				return
			}
			respo.Message = "Email is sent to Your email account " + admin.Email
			c.JSON(http.StatusOK, respo)
			return
		}
		respo.Message = os.Getenv("INTERNAL_SERVER_ERROR")
		c.JSON(http.StatusInternalServerError, respo)
		return
	} else {
		c.JSON(http.StatusNotFound, respo)
		respo.Message = "Account with this ID doesn't exist"
		return
	}

}

// CreateAdmin creates admin instance.
func (adminhr *AdminHandler) CreateAdmin(c *gin.Context) {
	input := &struct {
		Fullname   string `json:"fullname"`
		Email      string `json:"email"`
		Superadmin bool   `json:"superadmin"`
	}{}
	resp := &model.LoginResponse{false, "Bad Request Body", nil}

	currentTime := platforms.GetCurrentEthiopianTime()

	if er := c.BindJSON(input); er == nil {
		fail := false
		if !form.MatchesPattern(input.Email, form.EmailRX) {
			resp.Message = "Invalid email address!"
			fail = true
		}
		if len(input.Fullname) <= 1 || len(strings.Split(strings.Trim(input.Fullname, " "), " ")) <= 1 {
			resp.Message = " Invalid Fullname \n Your full name should include yours and your father's name!"
			fail = true
		}
		if !fail {
			// Generate Random password
			password := "admin" //helper.GenerateRandomString(5, helper.NUMBERS)
			hash, er := helper.HashPassword(password)

			ctx := c.Request.Context()
			ctx = context.WithValue(ctx, "email", input.Email)
			if admin, err := adminhr.AdminSer.AdminByEmail(ctx); admin != nil || err == nil {
				resp.Message = "account with this email already exist."
				c.JSON(http.StatusUnauthorized, resp)
				return
			}
			if er != nil {
				resp.Message = " Internal Server error "
				resp.Success = false
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			admin := &model.Admin{
				Fullname:   input.Fullname,
				Email:      input.Email, //
				Password:   hash,
				Superadmin: input.Superadmin,
				CreatedAt:  currentTime,
			}
			// Send Email for the password if this doesn't work raise internal server error.
			// if success := mail.SendPasswordEmailSMTP([]string{admin.Email}, password, true, admin.Fullname, c.Request.Host); success {
			ctx = c.Request.Context()
			ctx = context.WithValue(ctx, "admin", admin)
			if admin, er = adminhr.AdminSer.CreateAdmin(ctx); admin != nil && er == nil {
				resp.Success = true
				resp.Message = func() string {
					if admin.Superadmin {
						return " Super Admin"
					}
					return " Admin "
				}() + " created succesfully!"
				resp.User = admin
				c.JSON(http.StatusOK, resp)
				return
			} else {
				if admin != nil && er != nil {
					resp.Message = er.Error()
				} else {
					resp.Message = "Internal server error!"
				}
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
			// } else {
			// 	resp.Message = "Internal server error!"
			// 	c.JSON(http.StatusInternalServerError, resp)
			// 	return
			// }
		}
	}
	c.JSON(http.StatusBadRequest, resp)
}

// DeactivateAccount to deactivate an account usign the username and password
func (adminhr *AdminHandler) DeactivateAccount(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")
	resp := &struct {
		Msg string `json:"msg"`
	}{}
	if email == "" || password == "" {
		resp.Msg = os.Getenv("INVALID_INPUT")
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	ctx := c.Request.Context()
	if !form.MatchesPattern(email, form.EmailRX) {
		resp.Msg = "Invalid email address!"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	ctx = context.WithValue(ctx, "email", email)
	newAdmin, err := adminhr.AdminSer.AdminByEmail(ctx)
	if err != nil || newAdmin == nil {
		resp.Msg = "Invalid Username or Password!"
		c.JSON(401, resp)
		return
	} else {
		if newAdmin == nil {
			goto InvalidEmailsOrPassword
		}
		// comparing the hashed password and the password
		matches := hash.ComparePassword(newAdmin.Password, password)
		if !matches {
			goto InvalidEmailsOrPassword
		}
		ctx = context.WithValue(ctx, "email", email)
		if success, status := adminhr.AdminSer.DeleteAccountByEmail(ctx); success && status == state.DT_STATUS_OK {
			resp.Msg = "succesfuly deleted!"
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Msg = "Internal server error"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
InvalidEmailsOrPassword:
	{
		resp.Msg = state.InvalidEmailOrPassword
		c.JSON(401, resp)
		return
	}
}

func (adminhr *AdminHandler) UpdateAdmin(c *gin.Context) {
	input := &struct {
		Fullname string `json:"fullname,omitempty"`
		Email    string `json:"email,omitempty"`
	}{}
	res := &struct {
		Succ bool   `json:"success"`
		Msg  string `json:"msg"`
	}{}
	if err := c.BindJSON(input); err == nil {
		fail := false
		if input.Email != "" && !form.MatchesPattern(input.Email, form.EmailRX) {
			res.Msg = "Invalid email address!"
			fail = true
		}
		if (input.Fullname != "") && (len(input.Fullname) <= 5 || len(strings.Split(strings.Trim(input.Fullname, " "), " ")) <= 1) {
			res.Msg = " Invalid Fullname \n Your full name should include yours and your father's name!\n > 5 character total!"
			fail = true
		}
		if fail {
			res.Succ = false
			c.JSON(http.StatusBadRequest, res)
			return
		}
		ctx := c.Request.Context()
		// Get the admin By his ID and tell the newly Created Email to confirm that this is his email.
		session := c.Request.Context().Value("session").(*model.Session)
		if session == nil {
			res.Msg = "not authorized"
			c.JSON(http.StatusUnauthorized, res)
			return
		}
		ctx = context.WithValue(ctx, "user_id", session.ID)
		admin, era := adminhr.AdminSer.AdminByID(ctx)
		if admin == nil || era != nil {
			res.Msg = "internal server Error "
			c.JSON(http.StatusInternalServerError, res)
			return
		}
		changed := false
		if admin.Fullname != input.Fullname {
			admin.Fullname = input.Fullname
			changed = true
		}
		if input.Email != "" && admin.Email != input.Email {
			if success := mail.SendEmailChangeSMTP([]string{input.Email}, session.Password, admin.Fullname, c.Request.Host); success {
				admin.Email = input.Email
				changed = true
			} else {
				res.Msg = "internal server error!"
				c.JSON(http.StatusInternalServerError, res)
				return
			}
		}
		if changed {
			ctx = context.WithValue(ctx, "admin", admin)
			if admin, er := adminhr.AdminSer.UpdateAdmin(ctx); admin == nil || er != nil {
				res.Msg = "update was not succesful please try again!"
				c.JSON(http.StatusNotModified, res)
				return
			}
			res.Succ = true
			res.Msg = "succesfully updated"
			c.JSON(http.StatusOK, res)
			return
		} else {
			// No update was made.
			res.Msg = "no update was made"
			c.JSON(http.StatusNotModified, res)
			return
		}
	}
	res.Succ = false
	res.Msg = "Bad Request"
	c.JSON(http.StatusBadRequest, res)
}

// ChangeProfilePicture
func (adminhr *AdminHandler) ChangeProfilePicture(c *gin.Context) {
	var header *multipart.FileHeader
	var erro error
	var oldImage string
	erro = c.Request.ParseMultipartForm(99999999999)
	if erro != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	image, header, erro := c.Request.FormFile("image")
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	defer image.Close()
	if helper.IsImage(header.Filename) {
		newName := "images/profile/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
		var newImage *os.File
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			newImage, erro = os.Create(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		if erro != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		defer newImage.Close()
		oldImage = adminhr.AdminSer.GetImageUrl(c.Request.Context())
		_, er := io.Copy(newImage, image)
		if er != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		ncon := context.WithValue(c.Request.Context(), "user_id", uint64(c.Request.Context().Value("session").(*model.Session).ID))
		ncon = context.WithValue(ncon, "image_url", newName)
		success := adminhr.AdminSer.ChangeImageUrl(ncon)
		if success {
			if oldImage != "" {
				if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + oldImage)
				} else {
					er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + oldImage)
				}
			}
			c.JSON(http.StatusOK, &model.ShortSuccess{Msg: newName})
			return
		}
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + newName)
		} else {
			er = os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + newName)
		}
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.Writer.WriteHeader(http.StatusUnsupportedMediaType)
	}
}

// DeleteProfilePicture ...
func (adminhr *AdminHandler) DeleteProfilePicture(c *gin.Context) {
	imageUrl := adminhr.AdminSer.GetImageUrl(c.Request.Context())
	success := adminhr.AdminSer.DeleteProfilePicture(c.Request.Context())
	if success {
		if strings.HasSuffix(os.Getenv("ASSETS_DIRECTORY"), "/") {
			os.Remove(os.Getenv("ASSETS_DIRECTORY") + imageUrl)
		} else {
			os.Remove(os.Getenv("ASSETS_DIRECTORY") + "/" + imageUrl)
		}
		c.Writer.Write(helper.MarshalThis(&model.ShortSuccess{Msg: "Succesfully Deleted"}))
		return
	} else {
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (adminhr *AdminHandler) DeleteAdmin(c *gin.Context) {
	eres := &model.ErMsg{}
	res := &model.ShortSuccess{}
	email := c.Query("email")
	id, er := strconv.Atoi(c.Query("id"))

	ctx := c.Request.Context()

	if (email == "") && (er != nil || id == 0) {
		eres.Status = http.StatusBadRequest
		eres.Error = "valid \"id\" and \"email\"  is not provided"
		c.JSON(http.StatusBadRequest, eres)
		return
	}
	var admin *model.Admin
	var err error
	if id > 0 && er == nil {
		ctx = context.WithValue(ctx, "user_id", uint64(id))
		admin, err = adminhr.AdminSer.AdminByID(ctx)
	} else {
		ctx = context.WithValue(ctx, "email", email)
		admin, err = adminhr.AdminSer.AdminByID(ctx)
	}

	if admin == nil || err != nil {
		eres.Error = " admin with this id doesn't exist"
		eres.Status = http.StatusNotFound
		c.JSON(http.StatusNotFound, eres)
		return
	}
	var success bool
	var status int
	if id > 0 && er == nil {
		println(" Deleting Account by ID ")
		ctx = context.WithValue(ctx, "admin_id", uint64(id))
		success, status = adminhr.AdminSer.DeleteAccountByID(ctx)
	} else {
		ctx = context.WithValue(ctx, "email", email)
		success, status = adminhr.AdminSer.DeleteAccountByEmail(ctx)
	}
	if !success {
		if status == state.DT_STATUS_OK {
			eres.Error = " can't delete the admin instance please try again "
		} else {
			eres.Error = " internal problem, please try again!"
		}
		eres.Status = http.StatusInternalServerError
		c.JSON(http.StatusInternalServerError, eres)
		return
	}
	res.Msg = "admin deleted succesfully"
	c.JSON(http.StatusOK, res)
}

func (adminhr *AdminHandler) GetAdminsOfSystem(c *gin.Context) {
	ctx := c.Request.Context()
	admins, er := adminhr.AdminSer.GetAdmins(ctx)
	if er != nil || len(admins) == 0 {
		c.JSON(http.StatusNotFound, admins)
		return
	}
	c.JSON(http.StatusOK, admins)
}
