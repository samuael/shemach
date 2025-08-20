package user

import (
	"context"
	"log"

	"github.com/samuael/shemach/backend/pkg/constants/types"
)

type IUserService interface {
	// Inserts or replaces an image and return the image id with the old image path and status information, and error value if any
	// status infomrations:
	// -1: user has profile image but unable to update the profile image
	// -2: Error when inserting the image information.
	// -3 error when updating the user table information and the newly added image is removed because of that.
	// 0: image information is replaced and the old image url is retured
	// 1: new profile image information is registerd for the user.
	InsertReplaceUserProfileImage(ctx context.Context, path string, userid uint64) (status int8, oldPath string, imgID uint64, err error)

	GetUserByID(ctx context.Context, id uint64) (*types.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*types.User, error)

	// SaveForgotPasswordDetail saved a forgot password instance to the table
	//  1 if trial exceeded try again later
	//  2 update was not succesful
	//  3 insertion was not succesful
	//  {created_at} if number found in pending code confirmation
	SaveForgotPasswordDetail(ctx context.Context, phone, code string, phoneUsed, emailUsed, useTelegram bool, maxTrials uint8) (statusCode uint64, shortCode string, err error)
	DeleteExpiredForgotPasswordSession(ctx context.Context, timestamp uint64) (int64, error)

	// ConfirmForgotPasswordShortcode checks if the phone and the shortcode exists in the forgot_password_shortcode table
	//  -1 : account information not found
	//  -2 : trial already exceeded
	//  -3 : shortcode doesn;t match and trial exceeded
	//  -4 : shortcode does not match
	//  0 : success
	ConfirmForgotPasswordShortcode(ctx context.Context, phone, shortCode string) (int64, error)

	// ConfirmPhoneShortCode matches the short code and returns success status codes as follows
	// 0 for success
	// -1 for incorrect
	// -2 for not found
	// -3 for trial exceeded
	// -4 for internal operation error
	// -5 for verified phone creation error
	// otherwise the row id of verified phone number
	ConfirmPhoneShortCode(ctx context.Context, phone, code string) (int64, error)
	ChangeUserPassword(ctx context.Context, userID uint64, newPassword string) error
	// TempoRegisterNewUser registers a new user temporarily
	// checks the availability of the number in the system and returns status message if found or timestamp of creation if not found or at the stage of being verified.
	// -1 if number found in already registered users
	// -2 if insertion is not successful
	// -{created_at} if number found in pending code confirmation
	// {created_at} otherwise
	TempoRegisterNewUser(ctx context.Context, userData *types.UserAuthData) error

	CheckVerifiedInfoAndRegisterUser(ctx context.Context, session *types.TempoRegistrationSession, userData *types.RegistrationData) (*types.User, int64, error)

	DeleteExpiredTempoUsers(ctx context.Context, timestamp int64) (int, error)
	DeleteExpiredVerifiedPhones(ctx context.Context, timestamp int64) (int, error)
}

type UserService struct {
	Repo IUserRepo
}

func NewUserService(repo IUserRepo) IUserService {
	return &UserService{
		Repo: repo,
	}
}

func (service *UserService) GetUserByID(ctx context.Context, id uint64) (*types.User, error) {
	return service.Repo.GetUserByID(ctx, id)
}

// CheckVerifiedInfoAndRegisterUser checks if the verified phone number is found in the verified phone number informations and if found,
// it will register a new user with the specified detailed information
// -1 for not found -2 for intrenal database error -3 for user creation error
func (service *UserService) CheckVerifiedInfoAndRegisterUser(ctx context.Context, session *types.TempoRegistrationSession, userData *types.RegistrationData) (*types.User, int64, error) {
	statusCode, err := service.Repo.CheckVerifiedInfoAndRegisterUser(ctx, session, userData)
	if err != nil {
		return nil, -3, err
	}
	var user *types.User
	if statusCode > 0 {
		user, err = service.Repo.GetUserByID(ctx, uint64(statusCode))
		if err != nil {
			return nil, statusCode, err
		}
	}
	return user, statusCode, nil
}

// ConfirmPhoneShortCode returns an error code, timestamp if verified successfully, and error message
func (service *UserService) ConfirmPhoneShortCode(ctx context.Context, phone, code string) (int64, error) {
	statusCode, err := service.Repo.ConfirmPhoneShortCode(ctx, phone, code)
	if err != nil {
		log.Println(err.Error())
		return -6, ErrInternalDatabaseError
	}
	return statusCode, nil
}

func (service *UserService) DeleteExpiredVerifiedPhones(ctx context.Context, timestamp int64) (int, error) {
	return service.Repo.DeleteExpiredVerifiedPhones(ctx, timestamp)
}

func (service *UserService) DeleteExpiredTempoUsers(ctx context.Context, timestamp int64) (int, error) {
	return service.Repo.DeleteExpiredTempoUsers(ctx, timestamp)
}

// GetTempoUserInfo retrieves tempo user information
func (service *UserService) GetTempoUserInfo(ctx context.Context, phone string) (*types.UserAuthData, error) {
	return service.Repo.GetTempoUserInfo(ctx, phone)
}

// TempoRegisterNewUser registers a new user temporarily
// checks the availability of the number in the system and returns status message if found or timestamp of creation if not found or at the stage of being verified.
// -1 if number found in already registered users
// -2 if insertion is not successful
// -{created_at} if number found in pending code confirmation
// {created_at} otherwise
func (service *UserService) TempoRegisterNewUser(ctx context.Context, userData *types.UserAuthData) error {
	return service.Repo.TempoRegisterNewUser(ctx, userData)
}

// GetUserByPhone retrieves user's detail by phone
func (service *UserService) GetUserByPhone(ctx context.Context, phone string) (*types.User, error) {
	return service.Repo.GetUserByPhone(ctx, phone)
}

func (service *UserService) InsertReplaceUserProfileImage(ctx context.Context, path string, userid uint64) (status int8, oldPath string, imgID uint64, err error) {
	return service.Repo.InsertReplaceUserProfileImage(ctx, path, userid)
}

func (service *UserService) SaveForgotPasswordDetail(ctx context.Context, phone, code string, phoneUsed, emailUsed, useTelegram bool, maxTrials uint8) (statusCode uint64, shortCode string, err error) {
	return service.Repo.SaveForgotPasswordDetail(ctx, phone, code, phoneUsed, emailUsed, useTelegram, maxTrials)
}

func (service *UserService) DeleteExpiredForgotPasswordSession(ctx context.Context, timestamp uint64) (int64, error) {
	return service.Repo.DeleteExpiredForgotPasswordSession(ctx, timestamp)
}

func (service *UserService) ConfirmForgotPasswordShortcode(ctx context.Context, phone, shortCode string) (int64, error) {
	return service.Repo.ConfirmForgotPasswordShortcode(ctx, phone, shortCode)
}

func (service *UserService) ChangeUserPassword(ctx context.Context, userID uint64, newPassword string) error {
	return service.Repo.ChangeUserPassword(ctx, userID, newPassword)
}
