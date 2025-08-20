package rest

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/julienschmidt/httprouter"
	"github.com/samuael/shemach/backend/cmd/main/service"
	"github.com/samuael/shemach/backend/pkg/constants/types"
	"github.com/samuael/shemach/backend/pkg/http/rest/auth"
	"github.com/samuael/shemach/backend/pkg/user"
	"github.com/samuael/shemach/backend/platforms/form"
	"github.com/samuael/shemach/backend/platforms/helper"

	goJson "github.com/goccy/go-json"
)

type IUserHandler interface {
	RegisterUser(http.ResponseWriter, *http.Request, httprouter.Params)
	RegisterPhoneVerify(http.ResponseWriter, *http.Request, httprouter.Params)
	ConfirmCode(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UploadProfilePicture(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetUserLoginMethods(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GenerateForgotPasswordShortCode(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ForgotPasswordShortcodeConfirmation(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type UserHandler struct {
	Service       user.IUserService
	Authenticator auth.Authenticator
	Templates     *template.Template
	OTPService    *service.OtpService
}

func NewUserHandler(
	templates *template.Template,
	service user.IUserService,
	authenticator auth.Authenticator,
	otpService *service.OtpService,
) IUserHandler {
	return &UserHandler{
		Service:       service,
		Authenticator: authenticator,
		Templates:     templates,
		OTPService:    otpService,
	}
}

func (uHandler *UserHandler) UploadProfilePicture(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseMultipartForm(999999999)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input: " + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	image, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "valid image 'file' required", Status: types.ScBadRequest})
		w.Write(messageBody)
		return
	}
	defer image.Close()
	if !helper.IsImage(header.Filename) {
		fileSplit := strings.Split(header.Filename, ".")
		extension := header.Filename
		if len(fileSplit) > 1 && len(fileSplit[0]) > 0 {
			extension = fileSplit[0]
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "unsupported media type " + extension, Status: types.ScUnsupportedMediaType})
		w.Write(messageBody)
		return
	}
	newName := "images/users/" + helper.GenerateRandomString(5, helper.CHARACTERS) + "." + helper.GetExtension(header.Filename)
	var newImage *os.File
	newImage, err = os.Create(os.Getenv("ASSETS_DIRECTORY") + newName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	defer newImage.Close()

	_, err = io.Copy(newImage, image)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	ctx := r.Context()
	session, err := SessionFromContext(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helper.MarshalThis(&types.StatusMsg{Status: types.ScUnauthorized, Error: "unauthorized"}))
		return
	}

	status, oldPath, imageID, err := uHandler.Service.InsertReplaceUserProfileImage(ctx, newName, session.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error saving the image; please try again" + err.Error(), Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	switch status {
	case -1, -2, -3:
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error saving the image; please try again status" + strconv.Itoa(int(status))})
		w.Write(messageBody)
		return
	case 0:
		// remove the old image using the old path.
		// TODO: remove image by path.
		println("Old path: ", oldPath)
	default:
		// 0, and 1
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(helper.MarshalThis(types.StatusMsg{Success: true, Status: types.ScOK, Data: map[string]uint64{"imageId": imageID}}))
}

// ConfirmCode confirms a phone and short code for registration.
func (uhandler *UserHandler) ConfirmCode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	var codeConfirmation *types.ConfirmationCode
	err := goJson.NewDecoder(r.Body).Decode(&codeConfirmation)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input. correct and retry", Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	sCode, err := uhandler.Service.ConfirmPhoneShortCode(ctx, codeConfirmation.Phone, codeConfirmation.Code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "internal server error. please try again", Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	var statusCode int64
	var statusMessage string
	switch sCode {
	case -1:
		w.WriteHeader(http.StatusUnauthorized)
		statusCode = types.ScIncorrect
		statusMessage = "incorrect secret code"
	case -2:
		w.WriteHeader(http.StatusNotFound)
		statusCode = types.ScNotFound
		statusMessage = "unknown or expired phone number to verify"
	case -3:
		w.WriteHeader(http.StatusLocked)
		statusCode = types.ScTrialExceeded
		statusMessage = "trial exceeded"
	case -4, -5, -6:
		w.WriteHeader(http.StatusInternalServerError)
		statusCode = types.ScInternalError
		statusMessage = "internal server error"
	default:
		w.WriteHeader(http.StatusOK)
		statusCode = types.ScOK
		statusMessage = ""
	}
	var tokenString string
	if statusCode == types.ScOK {
		var success bool
		tokenString, success = uhandler.Authenticator.SaveRegistrationSession(w, &types.TempoRegistrationSession{
			ID:    sCode,
			Phone: codeConfirmation.Phone,
		})
		if !success {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(helper.MarshalThis(&types.StatusMsg{Error: "internal server error. please try again", Status: types.ScInternalError}))
			return
		}
	}
	getMessage := func(errMsg bool) string {
		switch {
		case errMsg && statusCode != types.ScOK:
			return statusMessage
		case !errMsg && statusCode == types.ScOK:
			return statusMessage
		}
		return ""
	}
	w.Write(helper.MarshalThis(&types.StatusMsg{Success: statusCode == types.ScOK, Error: getMessage(true), Message: getMessage(false), Status: statusCode, AuthToken: tokenString}))
}

// Login authenticates user.
func (uHandler *UserHandler) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	input := &struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{}
	jsonDecoder := goJson.NewDecoder(r.Body)
	err := jsonDecoder.Decode(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "bad request payload", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}
	if !form.MatchesPattern(input.Phone, form.PhoneRX) {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid phone number provided", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}
	input.Phone = strings.TrimPrefix(input.Phone, "+")
	input.Password = strings.Trim(input.Password, " ")
	if len(input.Password) < 4 {
		w.WriteHeader(http.StatusUnauthorized)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "incorrect phone number or password", Status: types.ScUnauthorized})
		w.Write(byteData)
		return
	}
	ctx := r.Context()
	user, err := uHandler.Service.GetUserByPhone(ctx, input.Phone)
	if err != nil {
		println("err: ", err.Error())
		w.WriteHeader(http.StatusNotFound)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid phone number or password", Status: http.StatusNotFound})
		w.Write(byteData)
		return
	}
	if !(helper.CompareHash(user.Password, input.Password)) {
		w.WriteHeader(http.StatusUnauthorized)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "incorrect phone number or password", Status: types.ScUnauthorized})
		w.Write(byteData)
		return
	}
	var byteData []byte
	authToken, err := uHandler.Authenticator.SaveSession(w, &types.Session{
		ID:   user.ID,
		Role: user.Role,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	byteData = helper.MarshalThis(&types.StatusMsg{Message: "Welcome", Status: types.ScOK, Success: true, Data: user, AuthToken: authToken})
	w.WriteHeader(http.StatusOK)
	w.Write(byteData)
}

func (uHandler *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	userID, err := strconv.ParseUint(params.ByName("userID"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid user ID provided", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}
	user, err := uHandler.Service.GetUserByID(ctx, userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "user not found", Status: types.ScNotFound})
		w.Write(byteData)
		return
	}
	w.Write(helper.MarshalThis(types.StatusMsg{Success: true, Status: types.ScOK, Data: user}))
}

func (uHandler *UserHandler) GetUserLoginMethods(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	phone := strings.Trim(r.URL.Query().Get("p"), " +")
	if !helper.ValidatePhone("+" + phone) {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid phone number provided", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}
	user, err := uHandler.Service.GetUserByPhone(ctx, phone)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		w.WriteHeader(http.StatusNotFound)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "user with the provided phone number not found", Status: types.ScNotFound}))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "internal server error; please try again" + err.Error(), Status: types.ScInternalError}))
		return
	}
	var email string
	if user.Email != "" && helper.MatchesPattern(user.Email, helper.EmailRX) && len(strings.Split(user.Email, "@")) == 2 {
		splits := strings.Split(user.Email, "@")
		if len(splits[0]) > 2 {
			email = splits[0][:2] + "****@" + splits[1]
		} else {
			email = "****@" + splits[1]
		}
	}
	var telegram string
	if len(user.Telegram) > 2 {
		telegram = "@" + user.Telegram[:2] + "****"
	}
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(&types.StatusMsg{Error: "user found", Status: types.ScOK, Data: &struct {
		Phone    string `json:"phone"`
		Email    string `json:"email,omitempty"`
		Telegram string `json:"telegram,omitempty"`
	}{
		Phone:    user.Phone,
		Email:    email,
		Telegram: telegram,
	}}))
}

func (uHandler *UserHandler) GenerateForgotPasswordShortCode(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	phone := strings.Trim(r.URL.Query().Get("acc"), " +")
	if !helper.ValidatePhone("+" + phone) {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid account phone number provided", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}
	method := strings.ToLower(strings.Trim(r.URL.Query().Get("m"), " "))
	possibleTags := []string{"e", "email", "phone", "p", "t", "telegram"}
	if !slices.Contains(possibleTags, method) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "invalid account login method provided; possible values are " + strings.Join(possibleTags, ","), Status: types.ScBadRequest}))
		return
	}
	shortCode := helper.GenerateRandomString(5, helper.NUMBERS)
	statusCode, shortCode, err := uHandler.Service.SaveForgotPasswordDetail(ctx, phone, shortCode, method == "p" || method == "phone", method == "e" || method == "email", method == "t" || method == "telegram", uint8(10))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "internal server error: " + err.Error(), Status: types.ScInternalError}))
		return
	}

	// uHandler.OTPService.DeleteChannel
	// TODO: Send shortcode to the process that handles short code sending
	println(shortCode)

	switch statusCode {
	case 1:
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "too many requests; please try again later", Status: types.ScTooManyRequest}))
		return
	case 2, 3:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "internal server error; please try again later", Status: types.ScInternalError}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(&types.StatusMsg{Message: "shortcode sent", Status: types.ScOK, Data: map[string]uint64{"expires_at": statusCode + uint64(5*60)}}))
}

func (uHandler *UserHandler) ForgotPasswordShortcodeConfirmation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	input := &struct {
		ShortCode string `json:"shortCode"`
		Acc       string `json:"acc"`
	}{}
	err := goJson.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		messageBody := helper.MarshalThis(&types.StatusMsg{Error: "bad input. correct and retry", Status: types.ScInternalError})
		w.Write(messageBody)
		return
	}
	if !helper.ValidatePhone("+" + input.Acc) {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "invalid account phone number provided", Status: types.ScBadRequest})
		w.Write(byteData)
		return
	}

	if len(input.ShortCode) != 5 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "invalid short code value", Status: types.ScInternalError}))
		return
	}
	status, err := uHandler.Service.ConfirmForgotPasswordShortcode(ctx, input.Acc, input.ShortCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "internal server error", Status: types.ScInternalError}))
		return
	}
	switch status {
	case -1:
		w.WriteHeader(http.StatusNotFound)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "account not found for forgot password confirmation", Status: types.ScNotFound}))
		return
	case -2, -3:
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "trial exceeded", Status: types.ScTooManyRequest}))
		return
	case -4:
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "incorrect shortcode", Status: types.ScIncorrect}))
		return
	}

	user, err := uHandler.Service.GetUserByPhone(ctx, input.Acc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "invalid phone number or password", Status: types.ScInternalError}))
		return
	}
	authToken, err := uHandler.Authenticator.SaveSession(w, &types.Session{
		ID:   user.ID,
		Role: user.Role,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(&types.StatusMsg{Message: "Confirmed", Status: types.ScOK, Success: true, Data: user, AuthToken: authToken}))
}

func (uHandler *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	input := &struct {
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}{}
	err := goJson.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "bad input. correct and retry", Status: types.ScInternalError}))
		return
	}
	if input.NewPassword != input.ConfirmPassword {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.MarshalThis(&types.StatusMsg{Error: "new password and confirm password does not match", Status: types.ScBadRequest}))
		return
	}
	input.NewPassword, err = helper.HashPassword(input.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Status: types.ScBadRequest, Error: "invalid password provided"})
		w.Write(byteData)
		return
	}
	session, err := SessionFromContext(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helper.MarshalThis(&types.StatusMsg{Status: types.ScUnauthorized, Error: "unauthorized"}))
		return
	}
	err = uHandler.Service.ChangeUserPassword(ctx, session.ID, input.NewPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.MarshalThis(&types.StatusMsg{Status: types.ScInternalError, Error: "internal server error; please try again"}))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(helper.MarshalThis(&types.StatusMsg{Success: true, Status: types.ScOK, Message: "password changed"}))
}

// RegisterPhoneVerify temporarily registers a user until complete registration
func (uhandler *UserHandler) RegisterPhoneVerify(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	var userData *types.UserAuthData
	err := goJson.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !helper.ValidatePhone(userData.Phone) {
		w.WriteHeader(http.StatusBadRequest)
		data := helper.MarshalThis(types.StatusMsg{Status: http.StatusBadRequest, Error: "Error:Invalid phone number"})
		w.Write(data)
		return
	}
	userData.Code = "1234"
	userData.Trials = 5
	err = uhandler.Service.TempoRegisterNewUser(ctx, userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	switch userData.CreatedAt {
	case -1: // if number found in already registered users
		w.WriteHeader(http.StatusConflict)
		byteData := helper.MarshalThis(types.StatusMsg{Status: types.ScConflict, Error: "Error: account with the provided phone number already found"})
		w.Write(byteData)
		return
	case -2: // if insertion is not successful
		w.WriteHeader(http.StatusInternalServerError)
		return
	default:
		var sCode int64
		if userData.CreatedAt > 0 {
			sCode = types.ScOK
		} else {
			sCode = types.ScFound
		}
		byteData, err := goJson.Marshal(&types.StatusMsg{Success: true, Status: sCode, Message: "code sent",
			Data: &types.OTPExpirationInfo{
				StartTimestamp:        userData.CreatedAt,
				ExpMinutes:            int64(uhandler.OTPService.PendingConfirmationCodeDuration),
				ResendDurationSeconds: int64(uhandler.OTPService.PendingConfirmationCodeDuration / 3),
			},
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(byteData)
	}
}

// RegisterUser handles user registration
func (uHandler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx, _ := context.WithDeadline(r.Context(), time.Now().Add(time.Second))
	w.Header().Set("Content-Type", "application/json")
	var userData *types.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	session, err := uHandler.Authenticator.GetTempoRegistrationSession(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userData.Password, err = helper.HashPassword(userData.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		byteData := helper.MarshalThis(&types.StatusMsg{Status: types.ScBadRequest, Error: "invalid password provided"})
		w.Write(byteData)
		return
	}
	user, statusCode, err := uHandler.Service.CheckVerifiedInfoAndRegisterUser(ctx, session, userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	switch statusCode {
	case -3, -2: // creation error and internal database error
		println("statusCode: ", statusCode)
		w.WriteHeader(http.StatusInternalServerError)
		return
	case -1: // not found
		w.WriteHeader(http.StatusNotFound)
		byteData := helper.MarshalThis(&types.StatusMsg{Error: "registration timeout; please try again", Status: types.ScNotFound})
		w.Write(byteData)
	default:
		var byteData []byte
		var authToken string
		authToken, err = uHandler.Authenticator.SaveSession(w, &types.Session{
			ID:   uint64(statusCode),
			Role: 0,
		})
		if err != nil {
			byteData = helper.MarshalThis(&types.StatusMsg{Message: "Registered succesfully", Status: types.ScCreated, Success: true, AuthToken: authToken})
		} else {
			byteData = helper.MarshalThis(&types.StatusMsg{Message: "Welcome", Status: types.ScOK, Success: true, Data: user, AuthToken: authToken})
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(byteData)
	}
}
